package management

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// Management is an Auth0 management client used to interact with the Auth0
// Management API v2.
//
type Management struct {
	// Client manages Auth0 Client (also known as Application) resources.
	Client *ClientManager

	// ClientGrant manages Auth0 ClientGrant resources.
	ClientGrant *ClientGrantManager

	// ResourceServer manages Auth0 Resource Server (also known as API)
	// resources.
	ResourceServer *ResourceServerManager

	// Connection manages Auth0 Connection resources.
	Connection *ConnectionManager

	// CustomDomain manages Auth0 Custom Domains.
	CustomDomain *CustomDomainManager

	// Grant manages Auth0 Grants.
	Grant *GrantManager

	// Log reads Auth0 Logs.
	Log *LogManager

	// RuleManager manages Auth0 Rules.
	Rule *RuleManager

	// RuleManager manages Auth0 Rule Configurations.
	RuleConfig *RuleConfigManager

	// Email manages Auth0 Email Providers.
	Email *EmailManager

	// EmailTemplate manages Auth0 Email Templates.
	EmailTemplate *EmailTemplateManager

	// User manages Auth0 User resources.
	User *UserManager

	// Tenant manages your Auth0 Tenant.
	Tenant *TenantManager

	// Ticket creates verify email or change password tickets.
	Ticket *TicketManager

	// Stat is used to retrieve usage statistics.
	Stat *StatManager

	domain   string
	basePath string
	timeout  time.Duration
	debug    bool

	http *http.Client
}

// New creates a new Auth0 Management client by authenticating using the
// supplied client id and secret.
func New(domain, clientID, clientSecret string, options ...apiOption) (*Management, error) {

	m := &Management{
		domain:   domain,
		basePath: "api/v2",
		timeout:  1 * time.Minute,
		debug:    false,
	}

	for _, option := range options {
		option(m)
	}

	m.http = newClient(domain, clientID, clientSecret, m.debug)

	m.Client = NewClientManager(m)
	m.ClientGrant = NewClientGrantManager(m)
	m.Connection = NewConnectionManager(m)
	m.CustomDomain = NewCustomDomainManager(m)
	m.Grant = NewGrantManager(m)
	m.Log = NewLogManager(m)
	m.ResourceServer = NewResourceServerManager(m)
	m.Rule = NewRuleManager(m)
	m.RuleConfig = NewRuleConfigManager(m)
	m.EmailTemplate = NewEmailTemplateManager(m)
	m.Email = NewEmailManager(m)
	m.User = NewUserManager(m)
	m.Tenant = NewTenantManager(m)
	m.Ticket = NewTicketManager(m)
	m.Stat = NewStatManager(m)

	return m, nil
}

func (m *Management) uri(path ...string) string {
	return (&url.URL{
		Scheme: "https",
		Host:   m.domain,
		Path:   m.basePath + "/" + strings.Join(path, "/"),
	}).String()
}

func (m *Management) q(options []reqOption) string {
	if len(options) == 0 {
		return ""
	}
	v := make(url.Values)
	for _, option := range options {
		option(v)
	}
	return "?" + v.Encode()
}

func (m *Management) request(method, uri string, v interface{}) error {

	var payload bytes.Buffer
	if v != nil {
		json.NewEncoder(&payload).Encode(v)
	}
	req, _ := http.NewRequest(method, uri, &payload)
	req.Header.Add("Content-Type", "application/json")

	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()

	if m.http == nil {
		m.http = http.DefaultClient
	}

	res, err := m.http.Do(req.WithContext(ctx))
	if err != nil {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			return err
		}
	}

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return newError(res.Body)
	}

	if res.StatusCode != http.StatusNoContent {
		defer res.Body.Close()
		return json.NewDecoder(res.Body).Decode(v)
	}

	return nil
}

func (m *Management) get(uri string, v interface{}) error {
	return m.request("GET", uri, v)
}

func (m *Management) post(uri string, v interface{}) error {
	return m.request("POST", uri, v)
}

func (m *Management) put(uri string, v interface{}) error {
	return m.request("PUT", uri, v)
}

func (m *Management) patch(uri string, v interface{}) error {
	return m.request("PATCH", uri, v)
}

func (m *Management) delete(uri string) error {
	return m.request("DELETE", uri, nil)
}

type apiOption func(*Management)

// WithTimeout configures the management client with a request timeout.
func WithTimeout(t time.Duration) apiOption {
	return func(m *Management) {
		m.timeout = t
	}
}

// WithDebug configures the management client to dump http requests and
// responses to stdout.
func WithDebug(d bool) apiOption {
	return func(m *Management) {
		m.debug = d
	}
}

type Error interface {
	Status() int
	error
}

type managementError struct {
	StatusCode int    `json:"statusCode"`
	Err        string `json:"error"`
	Message    string `json:"message"`
}

func newError(r io.Reader) error {
	m := &managementError{}
	err := json.NewDecoder(r).Decode(m)
	if err != nil {
		return err
	}
	return m
}

func (m *managementError) Error() string {
	return fmt.Sprintf("%d %s: %s", m.StatusCode, m.Err, m.Message)
}

func (m *managementError) Status() int {
	return m.StatusCode
}

// reqOption configures a call (typically to retrieve a resource) to Auth0 with
// query parameters.
type reqOption func(url.Values)

// WithFields configures a call to include the desired fields.
func WithFields(fields ...string) reqOption {
	return func(v url.Values) {
		v.Set("fields", strings.Join(fields, ","))
		v.Set("include_fields", "true")
	}
}

// WithoutFields configures a call to exclude the desired fields.
func WithoutFields(fields ...string) reqOption {
	return func(v url.Values) {
		v.Set("fields", strings.Join(fields, ","))
		v.Set("include_fields", "false")
	}
}

// Page configures a call to receive a specific page, if the results where
// concatenated.
func Page(page int) reqOption {
	return func(v url.Values) {
		v.Set("page", strconv.FormatInt(int64(page), 10))
	}
}

// PerPage configures a call to limit the amount of items in the result.
func PerPage(items int) reqOption {
	return func(v url.Values) {
		v.Set("per_page", strconv.FormatInt(int64(items), 10))
	}
}

// IncludeTotals configures a call to include totals.
func IncludeTotals(include bool) reqOption {
	return func(v url.Values) {
		v.Set("include_totals", strconv.FormatBool(include))
	}
}

// Parameter is a generic configuration to add arbitrary query parameters to
// calls made to Auth0.
func Parameter(key, value string) reqOption {
	return func(v url.Values) {
		v.Set(key, value)
	}
}

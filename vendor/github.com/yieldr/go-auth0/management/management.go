package management

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"

	"golang.org/x/oauth2"
)

// Config is the payload used to receive an Auth0 management token. This token
// is a JWT, it contains specific granted permissions (known as scopes), and it
// is signed with a application API key and secret for the entire tenant.
//
// 	{
// 	  "audience": "https://YOUR_AUTH0_DOMAIN/api/v2/",
// 	  "client_id": "YOUR_CLIENT_ID",
// 	  "client_secret": "YOUR_CLIENT_SECRET",
// 	  "grant_type": "client_credentials"
// 	}
//
// See: https://auth0.com/docs/api/management/v2/tokens#1-get-a-token
//
type Config struct {
	Audience     string `json:"audience"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
}

// Token is the response body from the request to receive an Auth0 management
// token.
//
// 	{
// 	  "access_token": "eyJ...Ggg",
// 	  "expires_in": 86400,
// 	  "scope": "read:clients create:clients read:client_keys",
// 	  "token_type": "Bearer"
// 	}
//
// See: https://auth0.com/docs/api/management/v2/tokens#2-use-the-token
//
type Token struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

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

	// RuleManager manages Auth0 Rules.
	Rule *RuleManager

	// RuleManager manages Auth0 Rule Configurations.
	RuleConfig *RuleConfigManager

	// Email manages Auth0 Email Providers.
	Email *EmailManager

	// EmailTemplate manages Auth0 Email Templates.
	EmailTemplate *EmailTemplateManager

	domain   string
	basePath string
	timeout  time.Duration
	debug    bool

	http *http.Client
}

// New creates a new Auth0 Management client by authenticating using the
// supplied client id and secret.
func New(domain, clientID, clientSecret string) (*Management, error) {

	m := &Management{
		domain:   domain,
		basePath: "api/v2",
		timeout:  1 * time.Minute,
		// debug:    true,
	}

	config := Config{
		Audience:     "https://" + domain + "/api/v2/",
		ClientID:     clientID,
		ClientSecret: clientSecret,
		GrantType:    "client_credentials",
	}

	var payload bytes.Buffer
	err := json.NewEncoder(&payload).Encode(config)
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest("POST", "https://"+domain+"/oauth/token", &payload)
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if m.debug {
		m.dump(req, res)
	}

	if res.StatusCode != http.StatusOK {
		return nil, newError(res.Body)
	}

	var token Token
	if err = json.NewDecoder(res.Body).Decode(&token); err != nil {
		return nil, err
	}

	ts := oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: token.AccessToken,
		TokenType:   token.TokenType,
		Expiry:      time.Now().Add(time.Duration(token.ExpiresIn) * time.Second),
	})

	m.http = RetryClient(oauth2.NewClient(context.Background(), ts))

	m.Client = NewClientManager(m)
	m.ClientGrant = NewClientGrantManager(m)
	m.Connection = NewConnectionManager(m)
	m.CustomDomain = NewCustomDomainManager(m)
	m.ResourceServer = NewResourceServerManager(m)
	m.Rule = NewRuleManager(m)
	m.RuleConfig = NewRuleConfigManager(m)
	m.EmailTemplate = NewEmailTemplateManager(m)
	m.Email = NewEmailManager(m)

	return m, nil
}

func (m *Management) getURI(parts ...string) string {
	return fmt.Sprintf("https://%s/%s/%s",
		m.domain,
		m.basePath,
		strings.Join(parts, "/"))
}

func (m *Management) get(uri string, v interface{}) error {

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()

	res, err := m.http.Do(req.WithContext(ctx))
	if err != nil {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			return err
		}
	}
	defer res.Body.Close()

	if m.debug {
		m.dump(req, res)
	}

	if res.StatusCode != http.StatusOK {
		return newError(res.Body)
	}

	return json.NewDecoder(res.Body).Decode(v)
}

func (m *Management) post(uri string, v interface{}) error {

	var payload bytes.Buffer
	json.NewEncoder(&payload).Encode(v)

	req, _ := http.NewRequest("POST", uri, &payload)
	req.Header.Add("Content-Type", "application/json")

	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()

	res, err := m.http.Do(req.WithContext(ctx))
	if err != nil {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			return err
		}
	}
	defer res.Body.Close()

	if m.debug {
		m.dump(req, res)
	}

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return newError(res.Body)
	}

	return json.NewDecoder(res.Body).Decode(v)
}

func (m *Management) put(uri string, v interface{}) error {

	var payload bytes.Buffer
	json.NewEncoder(&payload).Encode(v)

	req, _ := http.NewRequest("PUT", uri, &payload)
	req.Header.Add("Content-Type", "application/json")

	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()

	res, err := m.http.Do(req.WithContext(ctx))
	if err != nil {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			return err
		}
	}
	defer res.Body.Close()

	if m.debug {
		m.dump(req, res)
	}

	if res.StatusCode != http.StatusOK {
		return newError(res.Body)
	}

	return json.NewDecoder(res.Body).Decode(v)
}

func (m *Management) patch(uri string, v interface{}) error {

	var payload bytes.Buffer
	json.NewEncoder(&payload).Encode(v)

	req, _ := http.NewRequest("PATCH", uri, &payload)
	req.Header.Add("Content-Type", "application/json")

	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()

	res, err := m.http.Do(req.WithContext(ctx))
	if err != nil {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			return err
		}
	}
	defer res.Body.Close()

	if m.debug {
		m.dump(req, res)
	}

	if res.StatusCode != http.StatusOK {
		return newError(res.Body)
	}

	return json.NewDecoder(res.Body).Decode(v)
}

func (m *Management) delete(uri string) error {

	req, _ := http.NewRequest("DELETE", uri, nil)

	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()

	res, err := m.http.Do(req.WithContext(ctx))
	if err != nil {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			return err
		}
	}
	defer res.Body.Close()

	if m.debug {
		m.dump(req, res)
	}

	if res.StatusCode != http.StatusNoContent {
		return newError(res.Body)
	}

	return nil
}

func (m *Management) dump(req *http.Request, res *http.Response) {
	b1, _ := httputil.DumpRequest(req, true)
	b2, _ := httputil.DumpResponse(res, true)
	fmt.Printf("%s\n%s\b\n", b1, b2)
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

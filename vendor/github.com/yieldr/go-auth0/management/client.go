package management

import "encoding/json"

type Client struct {

	// The name of the client
	Name *string `json:"name,omitempty"`

	// Free text description of the purpose of the Client. (Max character length
	// is 140)
	Description *string `json:"description,omitempty"`

	// The id of the client
	ClientID *string `json:"client_id,omitempty"`

	// The client secret, it must not be public
	ClientSecret *string `json:"client_secret,omitempty"`

	// The type of application this client represents
	AppType *string `json:"app_type,omitempty"`

	// The URL of the client logo (recommended size: 150x150)
	LogoURI *string `json:"logo_uri,omitempty"`

	// Whether this client a first party client or not
	IsFirstParty *bool `json:"is_first_party,omitempty"`

	// Set header `auth0-forwarded-for` as trusted to be used as source
	// of end user ip for brute-force-protection on token endpoint.
	IsTokenEndpointIPHeaderTrusted *bool `json:"is_token_endpoint_ip_header_trusted,omitempty"`

	// Whether this client will conform to strict OIDC specifications
	OIDCConformant *bool `json:"oidc_conformant,omitempty"`

	// The URLs that Auth0 can use to as a callback for the client
	Callbacks      []interface{} `json:"callbacks,omitempty"`
	AllowedOrigins []interface{} `json:"allowed_origins,omitempty"`

	// A set of URLs that represents valid web origins for use with web message
	// response mode
	WebOrigins        []interface{}           `json:"web_origins,omitempty"`
	ClientAliases     []interface{}           `json:"client_aliases,omitempty"`
	AllowedClients    []interface{}           `json:"allowed_clients,omitempty"`
	AllowedLogoutURLs []interface{}           `json:"allowed_logout_urls,omitempty"`
	JWTConfiguration  *ClientJWTConfiguration `json:"jwt_configuration,omitempty"`

	// Client signing keys
	SigningKeys   []map[string]string `json:"signing_keys,omitempty"`
	EncryptionKey map[string]string   `json:"encryption_key,omitempty"`
	SSO           *bool               `json:"sso,omitempty"`

	// True to disable Single Sign On, false otherwise (default: false)
	SSODisabled *bool `json:"sso_disabled,omitempty"`

	// True if this client can be used to make cross-origin authentication
	// requests, false otherwise (default: false)
	CrossOriginAuth *bool `json:"cross_origin_auth,omitempty"`

	// List of acceptable Grant Types for this Client
	GrantTypes []interface{} `json:"grant_types,omitempty"`

	// URL for the location in your site where the cross origin verification
	// takes place for the cross-origin auth flow when performing Auth in your
	// own domain instead of Auth0 hosted login page
	CrossOriginLocation *string `json:"cross_origin_loc,omitempty"`

	// True if the custom login page is to be used, false otherwise. Defaults to
	// true
	CustomLoginPageOn      *bool                  `json:"custom_login_page_on,omitempty"`
	CustomLoginPage        *string                `json:"custom_login_page,omitempty"`
	CustomLoginPagePreview *string                `json:"custom_login_page_preview,omitempty"`
	FormTemplate           *string                `json:"form_template,omitempty"`
	Addons                 map[string]interface{} `json:"addons,omitempty"`

	// Defines the requested authentication method for the token endpoint.
	// Possible values are:
	// 	'none' (public client without a client secret),
	// 	'client_secret_post' (client uses HTTP POST parameters) or
	// 	'client_secret_basic' (client uses HTTP Basic)
	TokenEndpointAuthMethod *string                `json:"token_endpoint_auth_method,omitempty"`
	ClientMetadata          map[string]string      `json:"client_metadata,omitempty"`
	Mobile                  map[string]interface{} `json:"mobile,omitempty"`
}

func (c *Client) String() string {
	b, _ := json.MarshalIndent(c, "", "  ")
	return string(b)
}

type ClientJWTConfiguration struct {
	// The amount of seconds the JWT will be valid (affects exp claim)
	LifetimeInSeconds *int `json:"lifetime_in_seconds,omitempty"`

	// True if the client secret is base64 encoded, false otherwise. Defaults to
	// true
	SecretEncoded *bool `json:"secret_encoded,omitempty"`

	Scopes interface{} `json:"scopes,omitempty"`

	// Algorithm used to sign JWTs. Can be "HS256" or "RS256"
	Algorithm *string `json:"alg,omitempty"`
}

type ClientManager struct {
	m *Management
}

func NewClientManager(m *Management) *ClientManager {
	return &ClientManager{m}
}

func (cm *ClientManager) Create(c *Client) (err error) {
	return cm.m.post(cm.m.uri("clients"), c)
}

func (cm *ClientManager) Read(id string, opts ...reqOption) (*Client, error) {
	c := new(Client)
	err := cm.m.get(cm.m.uri("clients", id)+cm.m.q(opts), c)
	return c, err
}

func (cm *ClientManager) List(opts ...reqOption) ([]*Client, error) {
	var c []*Client
	err := cm.m.get(cm.m.uri("clients")+cm.m.q(opts), &c)
	return c, err
}

func (cm *ClientManager) Update(id string, c *Client) (err error) {
	return cm.m.patch(cm.m.uri("clients", id), c)
}

func (cm *ClientManager) RotateSecret(id string) (*Client, error) {
	c := new(Client)
	err := cm.m.post(cm.m.uri("clients", id, "rotate-secret"), c)
	return c, err
}

func (cm *ClientManager) Delete(id string) (err error) {
	return cm.m.delete(cm.m.uri("clients", id))
}

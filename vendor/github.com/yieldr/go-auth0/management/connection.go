package management

type Connection struct {

	// A generated string identifying the connection.
	ID string `json:"id,omitempty"`

	// The name of the connection. Must start and end with an alphanumeric
	// character and can only contain alphanumeric characters and '-'. Max
	// length 128.
	Name string `json:"name,omitempty"`

	// The identity provider identifier for the connection. Can be any of the
	// following:
	//
	// "ad", "adfs", "amazon", "dropbox", "bitbucket", "aol", "auth0-adldap",
	// "auth0-oidc", "auth0", "baidu", "bitly", "box", "custom", "daccount",
	// "dwolla", "email", "evernote-sandbox", "evernote", "exact", "facebook",
	// "fitbit", "flickr", "github", "google-apps", "google-oauth2", "guardian",
	//  "instagram", "ip", "linkedin", "miicard", "oauth1", "oauth2",
	// "office365", "paypal", "paypal-sandbox", "pingfederate",
	// "planningcenter", "renren", "salesforce-community", "salesforce-sandbox",
	//  "salesforce", "samlp", "sharepoint", "shopify", "sms", "soundcloud",
	// "thecity-sandbox", "thecity", "thirtysevensignals", "twitter", "untappd",
	//  "vkontakte", "waad", "weibo", "windowslive", "wordpress", "yahoo",
	// "yammer" or "yandex".
	Strategy string `json:"strategy,omitempty"`

	// Options for validation.
	Options *ConnectionOptions `json:"options,omitempty"`

	// The identifiers of the clients for which the connection is to be
	// enabled. If the array is empty or the property is not specified, no
	// clients are enabled.
	EnabledClients []interface{} `json:"enabled_clients,omitempty"`

	// Defines the realms for which the connection will be used (ie: email
	// domains). If the array is empty or the property is not specified, the
	// connection name will be added as realm.
	Realms []interface{} `json:"realms,omitempty"`

	Metadata interface{} `json:"metadata,omitempty"`
}

// ConnectionOptions general options
type ConnectionOptions struct {

	// Options for validation.
	Validation map[string]interface{} `json:"validation,omitempty"`

	// Password strength level, can be one of:
	// "none", "low", "fair", "good", "excellent" or null.
	PasswordPolicy string `json:"passwordPolicy,omitempty"`

	// Options for password history policy.
	PasswordHistory map[string]interface{} `json:"password_history,omitempty"`

	// Options for password expiration policy.
	PasswordNoPersonalInfo map[string]interface{} `json:"password_no_personal_info,omitempty"`

	// Options for password dictionary policy.
	PasswordDictionary map[string]interface{} `json:"password_dictionary,omitempty"`

	APIEnableUsers   bool `json:"api_enable_users,omitempty"`
	BasicProfile     bool `json:"basic_profile,omitempty"`
	ExtAdmin         bool `json:"ext_admin,omitempty"`
	ExtIsSuspended   bool `json:"ext_is_suspended,omitempty"`
	ExtAgreedTerms   bool `json:"ext_agreed_terms,omitempty"`
	ExtGroups        bool `json:"ext_groups,omitempty"`
	ExtAssignedPlans bool `json:"ext_assigned_plans,omitempty"`
	ExtProfile       bool `json:"ext_profile,omitempty"`

	// Options for adding parameters in the request to the upstream IdP.
	UpstreamParams interface{} `json:"upstream_params,omitempty"`

	ClientID           string        `json:"client_id,omitempty"`
	ClientSecret       string        `json:"client_secret,omitempty"`
	TenantDomain       string        `json:"tenant_domain,omitempty"`
	DomainAliases      []interface{} `json:"domain_aliases,omitempty"`
	UseWsfed           bool          `json:"use_wsfed,omitempty"`
	WaadProtocol       string        `json:"waad_protocol,omitempty"`
	WaadCommonEndpoint bool          `json:"waad_common_endpoint,omitempty"`
	AppID              string        `json:"app_id,omitempty"`
	AppDomain          string        `json:"app_domain,omitempty"`
}

type ConnectionManager struct {
	m *Management
}

func NewConnectionManager(m *Management) *ConnectionManager {
	return &ConnectionManager{m}
}

func (cm *ConnectionManager) Create(c *Connection) error {
	return cm.m.post(cm.m.getURI("connections"), c)
}

func (cm *ConnectionManager) Read(id string) (*Connection, error) {
	c := new(Connection)
	err := cm.m.get(cm.m.getURI("connections", id), c)
	return c, err
}

func (cm *ConnectionManager) Update(id string, c *Connection) (err error) {
	return cm.m.patch(cm.m.getURI("connections", id), c)
}

func (cm *ConnectionManager) Delete(id string) (err error) {
	return cm.m.delete(cm.m.getURI("connections", id))
}

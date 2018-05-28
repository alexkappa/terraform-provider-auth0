package management

type Email struct {

	// The name of the email provider. Can be one of "mandrill", "sendgrid",
	// "sparkpost", "ses" or "smtp".
	Name string `json:"name,omitempty"`

	// True if the email provider is enabled, false otherwise (defaults to true)
	Enabled bool `json:"enabled,omitempty"`

	// The default FROM address
	DefaultFromAddress string `json:"default_from_address,omitempty"`

	Credentials *EmailCredentials      `json:"credentials,omitempty"`
	Settings    map[string]interface{} `json:"settings,omitempty"`
}

type EmailCredentials struct {
	// API User
	APIUser string `json:"api_user,omitempty"`
	// API Key
	APIKey string `json:"api_key,omitempty"`
	// AWS Access Key ID
	AccessKeyID string `json:"accessKeyId,omitempty"`
	// AWS Secret Access Key
	SecretAccessKey string `json:"secretAccessKey,omitempty"`
	// AWS default region
	Region string `json:"region,omitempty"`
	// SMTP host
	SMTPHost string `json:"smtp_host,omitempty"`
	// SMTP port
	SMTPPort int `json:"smtp_port,omitempty"`
	// SMTP user
	SMTPUser string `json:"smtp_user,omitempty"`
	// SMTP password
	SMTPPass string `json:"smtp_pass,omitempty"`
}

type EmailManager struct {
	m *Management
}

func NewEmailManager(m *Management) *EmailManager {
	return &EmailManager{m}
}

func (em *EmailManager) Create(e *Email) error {
	return em.m.post(em.m.getURI("emails", "provider"), e)
}

func (em *EmailManager) Read() (*Email, error) {
	e := new(Email)
	err := em.m.get(em.m.getURI("emails", "provider"), e)
	return e, err
}

func (em *EmailManager) Update(e *Email) (err error) {
	return em.m.patch(em.m.getURI("emails", "provider"), e)
}

func (em *EmailManager) Delete() (err error) {
	return em.m.delete(em.m.getURI("emails", "provider"))
}

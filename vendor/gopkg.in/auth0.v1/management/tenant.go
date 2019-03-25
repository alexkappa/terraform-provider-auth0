package management

import "encoding/json"

type Tenant struct {
	// Change password page settings
	ChangePassword *TenantChangePassword `json:"change_password,omitempty"`

	// Guardian MFA page settings
	GuardianMFAPage *TenantGuardianMFAPage `json:"guardian_mfa_page,omitempty"`

	// Default audience for API Authorization
	DefaultAudience *string `json:"default_audience,omitempty"`

	// Name of the connection that will be used for password grants at the token
	// endpoint. Only the following connection types are supported: LDAP, AD,
	// Database Connections, Passwordless, Windows Azure Active Directory, ADFS.
	DefaultDirectory *string `json:"default_directory,omitempty"`

	ErrorPage *TenantErrorPage `json:"error_page,omitempty"`

	Flags *TenantFlags `json:"flags,omitempty"`

	// The friendly name of the tenant
	FriendlyName *string `json:"friendly_name,omitempty"`

	// The URL of the tenant logo (recommended size: 150x150)
	PictureURL *string `json:"picture_url,omitempty"`

	// User support email
	SupportEmail *string `json:"support_email,omitempty"`

	// User support URL
	SupportURL *string `json:"support_url,omitempty"`

	// A set of URLs that are valid to redirect to after logout from Auth0.
	AllowedLogoutURLs []interface{} `json:"allowed_logout_urls,omitempty"`

	// Login session lifetime, how long the session will stay valid (unit:
	// hours).
	SessionLifetime *int `json:"session_lifetime,omitempty"`

	// The selected sandbox version to be used for the extensibility environment
	SandboxVersion *string `json:"sandbox_version,omitempty"`

	// A set of available sandbox versions for the extensibility environment
	SandboxVersionAvailable []interface{} `json:"sandbox_versions_available,omitempty"`
}

func (t *Tenant) String() string {
	b, _ := json.MarshalIndent(t, "", "  ")
	return string(b)
}

type TenantChangePassword struct {
	// True to use the custom change password html, false otherwise.
	Enabled *bool `json:"enabled,omitempty"`
	// Replace default change password page with a custom HTML (Liquid syntax is
	// supported).
	HTML *string `json:"html,omitempty"`
}

type TenantGuardianMFAPage struct {
	// True to use the custom html for Guardian page, false otherwise.
	Enabled *bool `json:"enabled,omitempty"`
	// Replace default Guardian page with a custom HTML (Liquid syntax is
	// supported).
	HTML *string `json:"html,omitempty"`
}

type TenantErrorPage struct {
	// Replace default error page with a custom HTML (Liquid syntax is
	// supported).
	HTML *string `json:"html,omitempty"`
	// True to show link to log as part of the default error page, false
	// otherwise (default: true).
	ShowLogLink *bool `json:"show_log_link,omitempty"`
	// Redirect to specified url instead of show the default error page
	URL *string `json:"url,omitempty"`
}

type TenantFlags struct {
	// Enables the first version of the Change Password flow. We've deprecated
	// this option and recommending a safer flow. This flag is only for
	// backwards compatibility.
	ChangePasswordFlowV1 *bool `json:"change_pwd_flow_v1,omitempty"`

	// This flag determines whether all current connections shall be enabled
	// when a new client is created. Default value is true.
	EnableClientConnections *bool `json:"enable_client_connections,omitempty"`

	// This flag enables the API section in the Auth0 Management Dashboard.
	EnableAPIsSection *bool `json:"enable_apis_section,omitempty"`

	// If set to true all Impersonation functionality is disabled for the
	// Tenant. This is a read-only attribute.
	DisableImpersonation *bool `json:"disable_impersonation,omitempty"`

	// This flag enables advanced API Authorization scenarios.
	EnablePipeline2 *bool `json:"enable_pipeline2,omitempty"`

	// This flag enables dynamic client registration.
	EnableDynamicClientRegistration *bool `json:"enable_dynamic_client_registration,omitempty"`

	// If enabled, All your email links and urls will use your configured custom
	// domain. If no custom domain is found the email operation will fail.
	EnableCustomDomainInEmails *bool `json:"enable_custom_domain_in_emails,omitempty"`
}

type TenantManager struct {
	m *Management
}

func NewTenantManager(m *Management) *TenantManager {
	return &TenantManager{m}
}

func (tm *TenantManager) Read(opts ...reqOption) (*Tenant, error) {
	t := new(Tenant)
	err := tm.m.get(tm.m.uri("tenants/settings")+tm.m.q(opts), t)
	return t, err
}

func (tm *TenantManager) Update(t *Tenant) (err error) {
	return tm.m.patch(tm.m.uri("tenants/settings"), t)
}

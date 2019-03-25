package management

import "encoding/json"

type ResourceServer struct {

	// A generated string identifying the resource server.
	ID *string `json:"id,omitempty"`

	// The name of the resource server. Must contain at least one character.
	// Does not allow '<' or '>'
	Name *string `json:"name,omitempty"`

	// The identifier of the resource server.
	Identifier *string `json:"identifier,omitempty"`

	// Scopes supported by the resource server.
	Scopes []*ResourceServerScope `json:"scopes,omitempty"`

	// The algorithm used to sign tokens ["HS256" or "RS256"].
	SigningAlgorithm *string `json:"signing_alg,omitempty"`

	// The secret used to sign tokens when using symmetric algorithms.
	SigningSecret *string `json:"signing_secret,omitempty"`

	// Allows issuance of refresh tokens for this entity.
	AllowOfflineAccess *bool `json:"allow_offline_access,omitempty"`

	// The amount of time in seconds that the token will be valid after being
	// issued.
	TokenLifetime *int `json:"token_lifetime,omitempty"`

	// The amount of time in seconds that the token will be valid after being
	// issued from browser based flows. Value cannot be larger than
	// token_lifetime.
	TokenLifetimeForWeb *int `json:"token_lifetime_for_web,omitempty"`

	// Flag this entity as capable of skipping consent
	SkipConsentForVerifiableFirstPartyClients *bool `json:"skip_consent_for_verifiable_first_party_clients,omitempty"`

	// A URI from which to retrieve JWKs for this resource server used for
	// verifying the JWT sent to Auth0 for token introspection.
	VerificationLocation *string `json:"verificationLocation,omitempty"`

	Options map[string]interface{} `json:"options,omitempty"`
}

func (r *ResourceServer) String() string {
	b, _ := json.MarshalIndent(r, "", "  ")
	return string(b)
}

type ResourceServerScope struct {
	// The scope name. Use the format <action>:<resource> for example
	// 'delete:client_grants'.
	Value *string `json:"value,omitempty"`

	// Description of the scope
	Description *string `json:"description,omitempty"`
}

type ResourceServerManager struct {
	m *Management
}

func NewResourceServerManager(m *Management) *ResourceServerManager {
	return &ResourceServerManager{m}
}

func (r *ResourceServerManager) Create(rs *ResourceServer) (err error) {
	return r.m.post(r.m.uri("resource-servers"), rs)
}

func (r *ResourceServerManager) Read(id string, opts ...reqOption) (*ResourceServer, error) {
	rs := new(ResourceServer)
	err := r.m.get(r.m.uri("resource-servers", id)+r.m.q(opts), rs)
	return rs, err
}

func (r *ResourceServerManager) Update(id string, rs *ResourceServer) (err error) {
	return r.m.patch(r.m.uri("resource-servers", id), rs)
}

func (r *ResourceServerManager) Delete(id string) (err error) {
	return r.m.delete(r.m.uri("resource-servers", id))
}

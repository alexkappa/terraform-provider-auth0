package management

import "encoding/json"

type Grant struct {

	// The id of the grant.
	ID *string `json:"id,omitempty"`

	// The id of the client.
	ClientID *string `json:"clientID,omitempty"`

	// The id of the user.
	UserID *string `json:"user_id`

	// The grant's audience.
	Audience *string `json:"audience,omitempty"`

	Scope []interface{} `json:"scope,omitempty"`
}

func (g *Grant) String() string {
	b, _ := json.MarshalIndent(g, "", "  ")
	return string(b)
}

type GrantManager struct {
	m *Management
}

func NewGrantManager(m *Management) *GrantManager {
	return &GrantManager{m}
}

func (gm *GrantManager) List(opts ...reqOption) ([]*Grant, error) {
	var g []*Grant
	err := gm.m.get(gm.m.uri("grants")+gm.m.q(opts), &g)
	return g, err

}

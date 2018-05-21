package management

type ClientGrant struct {

	// A generated string identifying the client grant.
	ID string `json:"id,omitempty"`

	// The identifier of the client.
	ClientID string `json:"client_id,omitempty"`

	// The audience.
	Audience string `json:"audience,omitempty"`

	Scope []interface{} `json:"scope,omitempty"`
}

type ClientGrantManager struct {
	m *Management
}

func NewClientGrantManager(m *Management) *ClientGrantManager {
	return &ClientGrantManager{m}
}

func (r *ClientGrantManager) Create(g *ClientGrant) (err error) {
	return r.m.post(r.m.getURI("client-grants"), g)
}

func (r *ClientGrantManager) Read(id string) (*ClientGrant, error) {
	var gs []*ClientGrant
	err := r.m.get(r.m.getURI("client-grants"), &gs)
	if err != nil {
		return nil, err
	}
	for _, g := range gs {
		if g.ID == id {
			return g, nil
		}
	}
	return nil, &managementError{
		StatusCode: 404,
		Err:        "Not Found",
		Message:    "Client grant not found",
	}
}

func (r *ClientGrantManager) Update(id string, g *ClientGrant) (err error) {
	return r.m.patch(r.m.getURI("client-grants", id), g)
}

func (r *ClientGrantManager) Delete(id string) (err error) {
	return r.m.delete(r.m.getURI("client-grants", id))
}

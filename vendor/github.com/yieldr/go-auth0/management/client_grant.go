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

func (cg *ClientGrantManager) Create(g *ClientGrant) (err error) {
	return cg.m.post(cg.m.uri("client-grants"), g)
}

func (cg *ClientGrantManager) Read(id string) (*ClientGrant, error) {
	var gs []*ClientGrant
	err := cg.m.get(cg.m.uri("client-grants"), &gs)
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

func (cg *ClientGrantManager) Update(id string, g *ClientGrant) (err error) {
	return cg.m.patch(cg.m.uri("client-grants", id), g)
}

func (cg *ClientGrantManager) Delete(id string) (err error) {
	return cg.m.delete(cg.m.uri("client-grants", id))
}

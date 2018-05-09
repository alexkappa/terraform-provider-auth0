package management

import (
	"fmt"
	"testing"
	"time"
)

func TestClientGrant(t *testing.T) {

	var err error

	// We need a client and resource server to connect using a client grant. So
	// first we must create them.

	c := &Client{
		Name: fmt.Sprintf("Terraform Test - Client Grant (%s)",
			time.Now().Format(time.StampMilli)),
	}
	err = m.Client.Create(c)
	if err != nil {
		t.Fatal(err)
	}
	defer m.Client.Delete(c.ClientID)

	s := &ResourceServer{
		Name: fmt.Sprintf("Terraform Test - Client Grant (%s)",
			time.Now().Format(time.StampMilli)),
		Identifier: "https://api.example.com/client-grant",
		Scopes: []*ResourceServerScope{
			{
				Value:       "create:resource",
				Description: "Create Resource",
			},
			{
				Value:       "update:resource",
				Description: "Update Resource",
			},
		},
	}
	err = m.ResourceServer.Create(s)
	if err != nil {
		t.Fatal(err)
	}
	defer m.ResourceServer.Delete(s.ID)

	g := &ClientGrant{
		ClientID: c.ClientID,
		Audience: s.Identifier,
		Scope:    []interface{}{"create:resource"},
	}

	t.Run("Create", func(t *testing.T) {
		err = m.ClientGrant.Create(g)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", g)
	})

	t.Run("Read", func(t *testing.T) {
		g, err = m.ClientGrant.Read(g.ID)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", g)
	})

	t.Run("Update", func(t *testing.T) {
		id := g.ID
		g.ID = ""
		g.Audience = "" // read-only
		g.ClientID = "" // read-only
		g.Scope = append(g.Scope, "update:resource")

		err = m.ClientGrant.Update(id, g)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", g)
	})

	t.Run("Delete", func(t *testing.T) {
		err = m.ClientGrant.Delete(g.ID)
		if err != nil {
			t.Fatal(err)
		}
	})
}

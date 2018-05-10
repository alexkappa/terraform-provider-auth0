package management

import (
	"fmt"
	"testing"
	"time"
)

func TestResourceServer(t *testing.T) {

	s := &ResourceServer{
		Name:             fmt.Sprintf("Test Resource Server (%s)", time.Now().Format(time.StampMilli)),
		Identifier:       "https://api.example.com/",
		SigningAlgorithm: "HS256",
		Scopes: []*ResourceServerScope{
			{
				Value:       "create:resource",
				Description: "Create Resource",
			},
		},
	}

	var err error

	t.Run("Create", func(t *testing.T) {
		err = m.ResourceServer.Create(s)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", s)
	})

	t.Run("Read", func(t *testing.T) {
		s, err = m.ResourceServer.Read(s.ID)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", s)
	})

	t.Run("Update", func(t *testing.T) {
		id := s.ID
		s.ID = ""         // read-only
		s.Identifier = "" // read-only
		s.AllowOfflineAccess = true
		s.SigningAlgorithm = "RS256"
		s.SkipConsentForVerifiableFirstPartyClients = true
		s.Scopes = append(s.Scopes, &ResourceServerScope{
			Value:       "update:resource",
			Description: "Update Resource",
		})

		err = m.ResourceServer.Update(id, s)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", s)
	})

	t.Run("Delete", func(t *testing.T) {
		err = m.ResourceServer.Delete(s.ID)
		if err != nil {
			t.Error(err)
		}
	})
}

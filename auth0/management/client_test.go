package management

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestClient(t *testing.T) {

	c := &Client{
		Name: fmt.Sprintf("Terraform Test (%s)", time.Now().Format(time.StampMilli)),
		Description: "This is just a test client. It has been created during " +
			"tests to the Auth0 Terraform Provider",
	}

	var err error

	t.Run("Create", func(t *testing.T) {
		err = m.Client.Create(c)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", c)
	})

	t.Run("Read", func(t *testing.T) {
		c, err = m.Client.Read(c.ClientID)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", c)
	})

	t.Run("Update", func(t *testing.T) {
		id := c.ClientID
		c.ClientID = "" // read-only
		c.Description = strings.Replace(c.Description, "just", "more than", 1)
		err = m.Client.Update(id, c)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", c)
	})

	t.Run("Delete", func(t *testing.T) {
		err = m.Client.Delete(c.ClientID)
		if err != nil {
			t.Error(err)
		}
	})
}

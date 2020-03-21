package auth0

import (
	"errors"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"gopkg.in/auth0.v4/management"
)

func newGlobalClient() *schema.Resource {
	client := newClient()
	client.Create = createGlobalClient
	client.Delete = deleteGlobalClient

	name := client.Schema["name"]
	name.Required = false
	name.Computed = true

	return client
}

func createGlobalClient(d *schema.ResourceData, m interface{}) error {
	if err := readGlobalClientId(d, m); err != nil {
		return err
	}
	return updateClient(d, m)
}

func readGlobalClientId(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	clients, err := api.Client.List(management.Parameter("is_global", "true"), management.WithFields("client_id"))
	if err != nil {
		return err
	}
	if len(clients.Clients) == 0 {
		return errors.New("no auth0 global client found")
	}
	d.SetId(clients.Clients[0].GetClientID())
	return nil
}

func deleteGlobalClient(d *schema.ResourceData, m interface{}) error {
	return nil
}

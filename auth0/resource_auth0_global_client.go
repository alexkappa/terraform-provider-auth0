package auth0

import (
	"errors"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"gopkg.in/auth0.v5/management"
)

func newGlobalClient() *schema.Resource {
	client := newClient()
	client.Create = createGlobalClient
	client.Delete = deleteGlobalClient

	exclude := []string{"client_secret_rotation_trigger"}

	// Mark all values computed and optional. This because the global client has
	// already been created for all tenants.
	for key := range client.Schema {

		// Exclude certain fields from being marked as computed.
		if in(key, exclude) {
			continue
		}

		client.Schema[key].Required = false
		client.Schema[key].Optional = true
		client.Schema[key].Computed = true
	}

	return client
}

func in(needle string, haystack []string) bool {
	for i := 0; i < len(haystack); i++ {
		if needle == haystack[i] {
			return true
		}
	}
	return false
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

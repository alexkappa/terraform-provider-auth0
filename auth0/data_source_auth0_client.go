package auth0

import (
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"gopkg.in/auth0.v5"
	"gopkg.in/auth0.v5/management"
)

func newDataClient() *schema.Resource {
	clientSchema := newComputedClientSchema()
	addOptionalFieldsToSchema(clientSchema, "name", "client_id")

	return &schema.Resource{
		Read:   readDataClient,
		Schema: clientSchema,
	}
}

func newComputedClientSchema() map[string]*schema.Schema {
	clientSchema := datasourceSchemaFromResourceSchema(newClient().Schema)
	delete(clientSchema, "client_secret_rotation_trigger")
	return clientSchema
}

func readDataClient(d *schema.ResourceData, m interface{}) error {
	clientId := auth0.StringValue(String(d, "client_id"))
	if clientId == "" {
		name := auth0.StringValue(String(d, "name"))
		if name != "" {
			api := m.(*management.Management)
			clients, err := api.Client.List(management.WithFields("client_id", "name"))
			if err != nil {
				return err
			}
			for _, client := range clients.Clients {
				if auth0.StringValue(client.Name) == name {
					clientId = auth0.StringValue(client.ClientID)
					break
				}
			}
			if clientId == "" {
				return fmt.Errorf("no client found with 'name' = '%s'", name)
			}
		} else {
			return errors.New("no 'client_id' or 'name' was specified")
		}
	}
	d.SetId(clientId)
	return readClient(d, m)
}

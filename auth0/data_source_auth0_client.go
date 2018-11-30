package auth0

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/yieldr/go-auth0/management"
)

const (
	clientDataSchemaName              = "name"
	clientDataSchemaDescription       = "description"
	clientDataSchemaClientID          = "client_id"
	clientDataSchemaClientSecret      = "client_secret"
	clientDataSchemaSigningKeys       = "signing_keys"
)

func newClientDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readClientDataSource,
		Schema: map[string]*schema.Schema{
			clientDataSchemaClientID: {
				Type:     schema.TypeString,
				Optional: true,
			},
			clientDataSchemaName: {
				Type:     schema.TypeString,
				Computed: true,
			},
			clientDataSchemaDescription: {
				Type:     schema.TypeString,
				Computed: true,
			},
			clientDataSchemaClientSecret: {
				Type:     schema.TypeString,
				Computed: true,
				Sensitive: true,
			},
			clientDataSchemaSigningKeys: {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeMap,
				},
				Computed: true,
			},
		},
	}
}

func readClientDataSource(r *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)

	clientID, ok := r.Get(clientDataSchemaClientID).(string)

	if !ok {
		return fmt.Errorf("field '%s' was either not set or not a string", clientDataSchemaClientID)
	}

	c, err := api.Client.Read(clientID)

	if err != nil {
		return err
	}

	r.SetId(clientID)
	r.Set(clientDataSchemaName, c.Name)
	r.Set(clientDataSchemaDescription, c.Description)
	r.Set(clientDataSchemaClientSecret, c.ClientSecret)
	r.Set(clientDataSchemaSigningKeys, c.SigningKeys)

	return nil
}



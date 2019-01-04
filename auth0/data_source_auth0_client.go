package auth0

import (
	"errors"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/yieldr/go-auth0/management"
)

func newClientDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readClientDataSource,
		Schema: map[string]*schema.Schema{
			"client_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"client_secret": {
				Type:     schema.TypeString,
				Computed: true,
				Sensitive: true,
			},
			"signing_keys": {
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

	clientID := String(r, "client_id")

	if clientID == nil {
		return errors.New("client_id was not set")
	}

	c, err := api.Client.Read(*clientID)

	if err != nil {
		return err
	}

	r.SetId(*clientID)
	r.Set("name", c.Name)
	r.Set("description", c.Description)
	r.Set("client_secret", c.ClientSecret)
	r.Set("signing_keys", c.SigningKeys)

	return nil
}



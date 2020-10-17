package auth0

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"gopkg.in/auth0.v4/management"
)

func dataSourceAuth0Client() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAuth0ClientRead,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"client_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"client_secret": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
		},
	}
}

func dataSourceAuth0ClientRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)

	log.Printf("[INFO] Reading Auth0 Clients")
	var page int
	for {
		l, err := api.Client.List(management.WithFields("name", "description", "client_id", "client_secret"), management.Page(page))
		if err != nil {
			return err
		}
		log.Printf("[DEBUG] Reading Auth0 Clients for page %d", page)
		log.Printf("[DEBUG] Searching Auth0 Clients for Name %q", d.Get("name"))
		for _, c := range l.Clients {
			if strings.Contains(c.GetName(), fmt.Sprintf("%v", d.Get("name"))) {
				d.SetId(*c.ClientID)
				d.Set("client_id", c.ClientID)
				d.Set("client_secret", c.ClientSecret)
				d.Set("name", c.Name)
				d.Set("description", c.Description)
				log.Printf("[DEBUG] Found Auth0 Client with Name %v. SetID to %q", c.Name, d.Id())
				return nil
			}
		}
		if err != nil {
			return err
		}
		if !l.HasNext() {
			break
		}
		page++
	}
	return errors.New("No client found matching " + fmt.Sprintf("%v", d.Get("name")))
}

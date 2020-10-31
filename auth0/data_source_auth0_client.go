package auth0

import (
	"errors"
	"fmt"
	"log"
	"net/http"
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
				Computed: true,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"client_id": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
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

	name := d.Get("name").(string)

	clientID := d.Get("client_id").(string)

	if name == "" && clientID == "" {
		return errors.New(`The argument "name" or "client_id" should be configured`)
	}
	var tmpKeys map[string]string = make(map[string]string)

	if clientID != "" {
		log.Printf("[INFO] Reading Auth0 Client")
		c, err := api.Client.Read(clientID)
		if err != nil {
			if mErr, ok := err.(management.Error); ok {
				if mErr.Status() == http.StatusNotFound {
					d.SetId("")
					return nil
				}
			}
			return err
		}
		tmpKeys["client_id"] = c.GetClientID()
		tmpKeys["client_secret"] = c.GetClientSecret()
		tmpKeys["name"] = c.GetName()
		tmpKeys["description"] = c.GetDescription()
	} else {
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
				if strings.Compare(c.GetName(), fmt.Sprintf("%v", d.Get("name"))) == 0 {
					if len(tmpKeys) > 0 {
						log.Printf("[DEBUG] Found Multiple Auth0 Clients with name %v", c.GetName())
						return fmt.Errorf("Found Multiple Auth0 Clients with name %v", c.GetName())
					}
					tmpKeys["client_id"] = c.GetClientID()
					tmpKeys["client_secret"] = c.GetClientSecret()
					tmpKeys["name"] = c.GetName()
					tmpKeys["description"] = c.GetDescription()
					log.Printf("[DEBUG] Found Auth0 Client with Name %v", c.GetName())
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
	}
	if len(tmpKeys) > 0 {
		d.SetId(tmpKeys["client_id"])
		d.Set("client_id", tmpKeys["client_id"])
		d.Set("client_secret", tmpKeys["client_secret"])
		d.Set("name", tmpKeys["name"])
		d.Set("description", tmpKeys["description"])
		return nil
	}
	return errors.New("No client found matching " + fmt.Sprintf("%v", d.Get("name")))
}

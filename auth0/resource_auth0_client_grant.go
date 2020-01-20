package auth0

import (
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"gopkg.in/auth0.v3"
	"gopkg.in/auth0.v3/management"
)

func newClientGrant() *schema.Resource {
	return &schema.Resource{

		Create: createClientGrant,
		Read:   readClientGrant,
		Update: updateClientGrant,
		Delete: deleteClientGrant,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"client_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"audience": {
				Type:     schema.TypeString,
				Required: true,
			},
			"scope": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func createClientGrant(d *schema.ResourceData, m interface{}) error {
	g := buildClientGrant(d)
	api := m.(*management.Management)
	if err := api.ClientGrant.Create(g); err != nil {
		return err
	}
	d.SetId(auth0.StringValue(g.ID))
	return readClientGrant(d, m)
}

func readClientGrant(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	g, err := api.ClientGrant.Read(d.Id())
	if err != nil {
		if mErr, ok := err.(management.Error); ok {
			if mErr.Status() == http.StatusNotFound {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	d.SetId(auth0.StringValue(g.ID))
	d.Set("client_id", g.ClientID)
	d.Set("audience", g.Audience)
	d.Set("scope", g.Scope)
	return nil
}

func updateClientGrant(d *schema.ResourceData, m interface{}) error {
	g := buildClientGrant(d)
	g.Audience = nil
	g.ClientID = nil
	api := m.(*management.Management)
	err := api.ClientGrant.Update(d.Id(), g)
	if err != nil {
		return err
	}
	return readClientGrant(d, m)
}

func deleteClientGrant(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	err := api.ClientGrant.Delete(d.Id())
	if err != nil {
		if mErr, ok := err.(management.Error); ok {
			if mErr.Status() == http.StatusNotFound {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	return err
}

func buildClientGrant(d *schema.ResourceData) *management.ClientGrant {
	g := &management.ClientGrant{
		ClientID: String(d, "client_id"),
		Audience: String(d, "audience"),
	}
	if scope, ok := d.GetOk("scope"); ok {
		g.Scope = scope.([]interface{})
	} else {
		g.Scope = []interface{}{}
	}
	return g
}

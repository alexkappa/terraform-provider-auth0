package auth0

import (
	"github.com/hashicorp/terraform/helper/schema"
	"gopkg.in/auth0.v1"
	"gopkg.in/auth0.v1/management"
)

func newRole() *schema.Resource {
	return &schema.Resource{

		Create: createRole,
		Read:   readRole,
		Update: updateRole,
		Delete: deleteRole,
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
				Optional: true,
			},
		},
	}
}

func createRole(d *schema.ResourceData, m interface{}) error {
	c := buildRole(d)
	api := m.(*management.Management)
	if err := api.Role.Create(c); err != nil {
		return err
	}
	d.SetId(auth0.StringValue(c.ID))
	return readRole(d, m)
}

func readRole(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	c, err := api.Role.Read(d.Id())
	if err != nil {
		return err
	}
	d.Set("name", c.Name)
	d.Set("description", c.Description)
	return nil
}

func updateRole(d *schema.ResourceData, m interface{}) error {
	c := buildRole(d)
	api := m.(*management.Management)
	err := api.Role.Update(d.Id(), c)
	if err != nil {
		return err
	}
	return readRole(d, m)
}

func deleteRole(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	return api.Role.Delete(d.Id())
}

func buildRole(d *schema.ResourceData) *management.Role {
	return &management.Role{
		Name:        String(d, "name"),
		Description: String(d, "description"),
	}
}

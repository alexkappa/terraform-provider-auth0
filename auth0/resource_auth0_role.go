package auth0

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"net/http"

	"gopkg.in/auth0.v2"
	"gopkg.in/auth0.v2/management"
)

func newRole() *schema.Resource {
	return &schema.Resource{

		Create: createRole,
		Update: updateRole,
		Read:   readRole,
		Delete: deleteRole,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"role_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_ids": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				ForceNew: true,
				Removed:  `This field has been removed. Use "auth0_user.roles" instead`,
			},
			"permissions": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"resource_server_identifier": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
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

	// Enable partial state mode. Sub-resources can potentially cause partial
	// state. Therefore we must explicitly tell Terraform what is safe to
	// persist and what is not.
	//
	// See: https://www.terraform.io/docs/extend/writing-custom-providers.html
	d.Partial(true)
	if err := assignRolePermissions(d, m); err != nil {
		return err
	}
	// We succeeded, disable partial mode. This causes Terraform to save
	// all fields again.
	d.Partial(false)

	return readRole(d, m)
}

func readRole(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	c, err := api.Role.Read(d.Id())
	if err != nil {
		if mErr, ok := err.(management.Error); ok {
			if mErr.Status() == http.StatusNotFound {
				d.SetId("")
				return nil
			}
		}
		return err
	}

	d.SetId(auth0.StringValue(c.ID))
	d.Set("role_id", c.ID)
	d.Set("name", c.Name)
	d.Set("description", c.Description)

	permissions, err := api.Role.Permissions(d.Id())
	if err != nil {
		return err
	}
	d.Set("permissions", func() (m []map[string]interface{}) {
		for _, permission := range permissions {
			m = append(m, map[string]interface{}{
				"name":                       permission.Name,
				"resource_server_identifier": permission.ResourceServerIdentifier,
			})
		}
		return m
	}())

	return nil
}

func updateRole(d *schema.ResourceData, m interface{}) error {
	c := buildRole(d)
	api := m.(*management.Management)
	err := api.Role.Update(d.Id(), c)
	if err != nil {
		return err
	}
	d.Partial(true)
	if err := assignRolePermissions(d, m); err != nil {
		return err
	}
	d.Partial(false)
	return readRole(d, m)
}

func deleteRole(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	err := api.Role.Delete(d.Id())
	if err != nil {
		if mErr, ok := err.(management.Error); ok {
			if mErr.Status() == http.StatusNotFound {
				d.SetId("")
				return nil
			}
		}
	}
	return err
}

func buildRole(d *schema.ResourceData) *management.Role {
	return &management.Role{
		ID:          String(d, "role_id"),
		Name:        String(d, "name"),
		Description: String(d, "description"),
	}
}

func assignRolePermissions(d *schema.ResourceData, m interface{}) error {

	add, rm := Diff(d, "permissions")

	var addPermissions []*management.Permission
	for _, addPermission := range add {
		permission := addPermission.(map[string]interface{})
		addPermissions = append(addPermissions, &management.Permission{
			Name:                     auth0.String(permission["name"].(string)),
			ResourceServerIdentifier: auth0.String(permission["resource_server_identifier"].(string)),
		})
	}

	var rmPermissions []*management.Permission
	for _, rmPermission := range rm {
		permission := rmPermission.(map[string]interface{})
		rmPermissions = append(rmPermissions, &management.Permission{
			Name:                     auth0.String(permission["name"].(string)),
			ResourceServerIdentifier: auth0.String(permission["resource_server_identifier"].(string)),
		})
	}

	api := m.(*management.Management)

	if len(rmPermissions) > 0 {
		err := api.Role.RemovePermissions(d.Id(), rmPermissions...)
		if err != nil {
			return err
		}
	}

	if len(addPermissions) > 0 {
		err := api.Role.AssociatePermissions(d.Id(), addPermissions...)
		if err != nil {
			return err
		}
	}

	d.SetPartial("permissions")
	return nil
}

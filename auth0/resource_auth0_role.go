package auth0

import (
	"github.com/hashicorp/terraform/helper/schema"
	"gopkg.in/auth0.v1"
	"gopkg.in/auth0.v1/management"
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
			},
			"permissions": {
				Type: schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"resource_server_identifier": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
					},
				},
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

	// Enable partial state mode. Sub-resources can potentially cause partial
	// state. Therefore we must explicitly tell Terraform what is safe to
	// persist and what is not.
	//
	// See: https://www.terraform.io/docs/extend/writing-custom-providers.html
	d.Partial(true)

	if d.HasChange("user_ids") {
		users := buildUsers(d)
		if len(users) > 0 {
			err := api.Role.AssignUsers(*c.ID, users...)
			if err != nil {
				return err
			}
		}
		d.SetPartial("user_ids")
	}

	if d.HasChange("permissions") {
		permissions := buildPermissions(d)
		if len(permissions) > 0 {
			err := api.Role.AssignPermissions(*c.ID, permissions...)
			if err != nil {
				return err
			}
		}
		d.SetPartial("permissions")
	}

	// We succeeded, disable partial mode. This causes Terraform to save
	// all fields again.
	d.Partial(false)
	d.SetId(auth0.StringValue(c.ID))

	return readRole(d, m)
}

func readRole(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	c, err := api.Role.Read(d.Id())
	if err != nil {
		return err
	}

	d.SetId(auth0.StringValue(c.ID))
	d.Set("role_id", c.ID)
	d.Set("name", c.Name)
	d.Set("description", c.Description)

	users, err := api.Role.Users(d.Id())
	if err != nil {
		return err
	}

	userIDs := []string{}
	for _, user := range users {
		userIDs = append(userIDs, *user.ID)
	}
	d.Set("user_ids", userIDs)

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
	return readRole(d, m)
}

func deleteRole(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	return api.Role.Delete(d.Id())
}

func buildRole(d *schema.ResourceData) *management.Role {
	return &management.Role{
		ID:          String(d, "role_id"),
		Name:        String(d, "name"),
		Description: String(d, "description"),
	}
}

func buildUsers(d *schema.ResourceData) []*management.User {
	var users []*management.User
	for _, val := range Slice(d, "user_ids") {
		userID, _ := val.(string)
		users = append(users, &management.User{
			ID: &userID,
		})
	}
	return users
}

func buildPermissions(d *schema.ResourceData) []*management.Permission {
	var permissions []*management.Permission
	for _, val := range Slice(d, "permissions") {
		permission := val.(map[string]interface{})
		permissions = append(permissions, &management.Permission{
			Name:                     String(MapData(permission), "name"),
			ResourceServerIdentifier: String(MapData(permission), "resource_server_identifier"),
		})
	}
	return permissions
}

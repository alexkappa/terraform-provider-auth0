package auth0

import (
	"github.com/hashicorp/terraform/helper/schema"
	"gopkg.in/auth0.v1/management"
)

func newUserRole() *schema.Resource {
	return &schema.Resource{

		Create: assignRoles,
		Read:   getRoles,
		Delete: unassignRoles,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"user_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"roles": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
				ForceNew: true,
			},
		},
	}
}

// Associate an array of roles with a user.
// A list of roles ids to associated with the user.
func assignRoles(d *schema.ResourceData, m interface{}) error {

	u := &management.User{
		ID: String(d, "user_id"),
	}

	roles := buildRoles(d)
	api := m.(*management.Management)
	if err := api.User.AssignRoles(*u.ID, roles...); err != nil {
		return err
	}
	return getRoles(d, m)
}

// List the the roles associated with a user.
func getRoles(d *schema.ResourceData, m interface{}) error {

	u := &management.User{
		ID: String(d, "user_id"),
	}

	api := m.(*management.Management)
	roles, err := api.User.GetRoles(*u.ID)
	if err != nil {
		return err
	}
	d.Set("roles", roles) // here I'm not sure should I change the schema to management.Role?
	return nil
}

// Removes an array of roles from a user.
// A list of roles ids to unassociate from the user.
func unassignRoles(d *schema.ResourceData, m interface{}) error {
	roles := buildRoles(d)
	api := m.(*management.Management)

	u := &management.User{
		ID: String(d, "user_id"),
	}

	return api.User.UnassignRoles(*u.ID, roles...)
}

func buildRoles(d *schema.ResourceData) []*management.Role {

	var result []*management.Role

	for _, val := range Slice(d, "roles") {
		roleID, _ := val.(string)
		result = append(result, &management.Role{
			ID: &roleID,
		})
	}

	return result
}

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
		},
	}
}

func createRole(d *schema.ResourceData, m interface{}) error {
	c := buildRole(d)
	api := m.(*management.Management)
	if err := api.Role.Create(c); err != nil {
		return err
	}

	users := buildUsers(d)
	if err := api.Role.AssignUsers(*c.ID, users...); err != nil {
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

	users, err := api.Role.Users(d.Id())
	if err != nil {
		return err
	}

	d.Set("role_id", c.ID)
	d.Set("name", c.Name)
	d.Set("description", c.Description)

	user_ids := []string{}
	for _, user := range users {
		user_ids = append(user_ids, *user.ID)
	}
	d.Set("user_ids", user_ids)
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

	var result []*management.User

	for _, val := range Slice(d, "user_ids") {
		userID, _ := val.(string)
		result = append(result, &management.User{
			ID: &userID,
		})
	}

	return result
}

package auth0

import (
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"gopkg.in/auth0.v5"
	"gopkg.in/auth0.v5/management"
)

func newUserRoles() *schema.Resource {
	return &schema.Resource{
		Create: createUserRoles,
		Read:   readUserRoles,
		Update: updateUserRoles,
		Delete: deleteUserRoles,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"user_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return old == "auth0|"+new
				},
				StateFunc: func(s interface{}) string {
					str := s.(string)
					switch {
					// AAD user IDs are case sensitive and can't be ran through ToLower
					case strings.HasPrefix(str, "waad|"):
						return str
					default:
						return strings.ToLower(str)
					}
				},
			},
			"roles": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func readUserRoles(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)

	d.Set("user_id", d.Id())

	l, err := api.User.Roles(d.Id())
	if err != nil {
		return err
	}
	d.Set("roles", func() (v []interface{}) {
		for _, role := range l.Roles {
			v = append(v, auth0.StringValue(role.ID))
		}
		return
	}())

	return nil
}

func createUserRoles(d *schema.ResourceData, m interface{}) error {
	d.SetId(d.Get("user_id").(string))

	d.Partial(true)
	err := assignUserRoles(d, m)
	if err != nil {
		return err
	}
	d.Partial(false)

	return readUserRoles(d, m)
}

func updateUserRoles(d *schema.ResourceData, m interface{}) error {
	d.Partial(true)
	err := assignUserRoles(d, m)
	if err != nil {
		return err
	}
	d.Partial(false)

	return readUserRoles(d, m)
}

func deleteUserRoles(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	l, err := api.User.Roles(d.Id())
	if err != nil {
		return err
	}

	if len(l.Roles) > 0 {
		err := api.User.RemoveRoles(d.Id(), l.Roles)
		if err != nil {
			// Ignore 404 errors as the role may have been deleted prior to
			// unassigning them from the user.
			if mErr, ok := err.(management.Error); ok {
				if mErr.Status() != http.StatusNotFound {
					return err
				}
			} else {
				return err
			}
		}
	}

	d.SetId("")

	return nil
}

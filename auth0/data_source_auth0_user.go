package auth0

import (
	"errors"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"gopkg.in/auth0.v4/management"
)

func datasourceUser() *schema.Resource {
	return &schema.Resource{
		Read: datasourceUserRead,

		Schema: map[string]*schema.Schema{
			"user_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"email": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"email_verified": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"username": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"phone_number": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"phone_verified": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"created_at": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"updated_at": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"identities": {
				Type: schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"connection": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"user_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"provider": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"is_social": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},

				Optional: true,
			},

			"app_metadata": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"user_metadata": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"picture": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"nickname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"blocked": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"family_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"given_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func datasourceUserRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)

	if d.Get("user_id") == nil && d.Get("email") == nil {
		return errors.New("user_id or email should be configured")
	}

	var u *management.User

	var err error

	if d.Get("email").(string) != "" {

		users, listByEmailError := api.User.ListByEmail(d.Get("email").(string))

		if listByEmailError != nil {
			return listByEmailError
		}

		if len(users) == 0 {
			return errors.New("User not found")
		}

		u = users[0]

	} else {
		u, err = api.User.Read(d.Get("user_id").(string))
	}

	if err != nil {
		if mErr, ok := err.(management.Error); ok {
			if mErr.Status() == http.StatusNotFound {
				d.SetId("")
				return nil
			}
		}
		return err
	}

	d.SetId(*u.ID)

	d.Set("user_id", *u.ID)
	d.Set("email", u.Email)
	d.Set("email_verified", u.EmailVerified)

	d.Set("username", u.Username)
	d.Set("phone_number", u.PhoneNumber)
	d.Set("phone_verified", u.PhoneVerified)
	d.Set("created_at", u.CreatedAt.String())
	d.Set("updated_at", u.UpdatedAt.String())

	identities := make([]map[string]interface{}, len(u.Identities))

	for index, i := range u.Identities {
		newIdentity := make(map[string]interface{})
		newIdentity["connection"] = i.Connection
		newIdentity["user_id"] = i.UserID
		newIdentity["provider"] = i.Provider
		newIdentity["is_social"] = i.IsSocial
		identities[index] = newIdentity
	}

	d.Set("identities", identities)
	d.Set("app_metadata", u.AppMetadata)
	d.Set("user_metadata", u.UserMetadata)
	d.Set("picture", u.Picture)
	d.Set("name", u.Name)
	d.Set("nickname", u.Nickname)
	d.Set("blocked", u.Blocked)
	d.Set("family_name", u.FamilyName)
	d.Set("give_name", u.GivenName)

	return nil
}

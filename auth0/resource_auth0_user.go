package auth0

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/structure"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"gopkg.in/auth0.v5"
	"gopkg.in/auth0.v5/management"
)

func newUser() *schema.Resource {
	return &schema.Resource{
		Create: createUser,
		Read:   readUser,
		Update: updateUser,
		Delete: deleteUser,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"user_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return old == "auth0|"+new
				},
				StateFunc: func(s interface{}) string {
					return strings.ToLower(s.(string))
				},
			},
			"connection_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"username": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"family_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"given_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"nickname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"password": {
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"email": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"email_verified": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"verify_email": {
				Type:     schema.TypeBool,
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
			"user_metadata": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateFunc:     validation.ValidateJsonString,
				DiffSuppressFunc: structure.SuppressJsonDiff,
			},
			"app_metadata": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateFunc:     validation.ValidateJsonString,
				DiffSuppressFunc: structure.SuppressJsonDiff,
			},
			"blocked": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"picture": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"roles": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func readUser(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	u, err := api.User.Read(d.Id())
	if err != nil {
		if mErr, ok := err.(management.Error); ok {
			if mErr.Status() == http.StatusNotFound {
				d.SetId("")
				return nil
			}
		}
		return err
	}

	d.Set("user_id", u.ID)
	d.Set("username", u.Username)
	d.Set("name", u.Name)
	d.Set("family_name", u.FamilyName)
	d.Set("given_name", u.GivenName)
	d.Set("nickname", u.Nickname)
	d.Set("email", u.Email)
	d.Set("email_verified", u.EmailVerified)
	d.Set("verify_email", u.VerifyEmail)
	d.Set("phone_number", u.PhoneNumber)
	d.Set("phone_verified", u.PhoneVerified)
	d.Set("blocked", u.Blocked)
	d.Set("picture", u.Picture)

	userMeta, err := structure.FlattenJsonToString(u.UserMetadata)
	if err != nil {
		return err
	}
	d.Set("user_metadata", userMeta)

	appMeta, err := structure.FlattenJsonToString(u.AppMetadata)
	if err != nil {
		return err
	}
	d.Set("app_metadata", appMeta)

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

func createUser(d *schema.ResourceData, m interface{}) error {
	u, err := buildUser(d)
	if err != nil {
		return err
	}
	api := m.(*management.Management)
	if err := api.User.Create(u); err != nil {
		return err
	}
	d.SetId(*u.ID)

	d.Partial(true)
	err = assignUserRoles(d, m)
	if err != nil {
		return err
	}
	d.Partial(false)

	return readUser(d, m)
}

func updateUser(d *schema.ResourceData, m interface{}) error {
	u, err := buildUser(d)
	if err != nil {
		return err
	}
	if err = validateUser(u); err != nil {
		return err
	}
	api := m.(*management.Management)
	if userHasChange(u) {
		if err := api.User.Update(d.Id(), u); err != nil {
			return err
		}
	}
	d.Partial(true)
	err = assignUserRoles(d, m)
	if err != nil {
		return fmt.Errorf("failed assigning user roles. %s", err)
	}
	d.Partial(false)
	return readUser(d, m)
}

func deleteUser(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	err := api.User.Delete(d.Id())
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

func buildUser(d *schema.ResourceData) (u *management.User, err error) {

	u = new(management.User)
	u.ID = String(d, "user_id", IsNewResource())
	u.Connection = String(d, "connection_name")

	u.Name = String(d, "name")
	u.FamilyName = String(d, "family_name")
	u.GivenName = String(d, "given_name")
	u.Nickname = String(d, "nickname")

	u.Username = String(d, "username", IsNewResource(), HasChange())

	u.Email = String(d, "email", IsNewResource(), HasChange())
	u.EmailVerified = Bool(d, "email_verified", IsNewResource(), HasChange())
	u.VerifyEmail = Bool(d, "verify_email", IsNewResource(), HasChange())

	u.PhoneNumber = String(d, "phone_number", IsNewResource(), HasChange())
	u.PhoneVerified = Bool(d, "phone_verified", IsNewResource(), HasChange())

	u.Password = String(d, "password", IsNewResource(), HasChange())

	u.Blocked = Bool(d, "blocked")
	u.Picture = String(d, "picture")

	u.UserMetadata, err = JSON(d, "user_metadata")
	if err != nil {
		return nil, err
	}

	u.AppMetadata, err = JSON(d, "app_metadata")
	if err != nil {
		return nil, err
	}

	return u, nil
}

func validateUser(u *management.User) error {
	var validation error
	for _, fn := range []validateUserFunc{
		validateNoUsernameAndPasswordSimultaneously(),
		validateNoUsernameAndEmailVerifiedSimultaneously(),
		validateNoPasswordAndEmailVerifiedSimultaneously(),
	} {
		if err := fn(u); err != nil {
			validation = multierror.Append(validation, err)
		}
	}
	return validation
}

type validateUserFunc func(*management.User) error

func validateNoUsernameAndPasswordSimultaneously() validateUserFunc {
	return func(u *management.User) (err error) {
		if u.Username != nil && u.Password != nil {
			err = fmt.Errorf("Cannot update username and password simultaneously")
		}
		return
	}
}

func validateNoUsernameAndEmailVerifiedSimultaneously() validateUserFunc {
	return func(u *management.User) (err error) {
		if u.Username != nil && u.EmailVerified != nil {
			err = fmt.Errorf("Cannot update username and email_verified simultaneously")
		}
		return
	}
}

func validateNoPasswordAndEmailVerifiedSimultaneously() validateUserFunc {
	return func(u *management.User) (err error) {
		if u.Password != nil && u.EmailVerified != nil {
			err = fmt.Errorf("Cannot update password and email_verified simultaneously")
		}
		return
	}
}

func assignUserRoles(d *schema.ResourceData, m interface{}) error {

	add, rm := Diff(d, "roles")

	var addRoles []*management.Role
	for _, addRole := range add {
		addRoles = append(addRoles, &management.Role{
			ID: auth0.String(addRole.(string)),
		})
	}

	var rmRoles []*management.Role
	for _, rmRole := range rm {
		rmRoles = append(rmRoles, &management.Role{
			ID: auth0.String(rmRole.(string)),
		})
	}

	api := m.(*management.Management)

	if len(rmRoles) > 0 {
		err := api.User.RemoveRoles(d.Id(), rmRoles)
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

	if len(addRoles) > 0 {
		err := api.User.AssignRoles(d.Id(), addRoles)
		if err != nil {
			return err
		}
	}

	d.SetPartial("roles")
	return nil
}

func userHasChange(u *management.User) bool {
	// hacky but we need to tell if an empty json is sent to the api.
	return u.String() != "{}"
}

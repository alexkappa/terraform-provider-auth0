package auth0

import (
	"errors"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/structure"
	"gopkg.in/auth0.v5/management"
)

func newDataUser() *schema.Resource {
	userSchema := newComputedDataSchema()
	addOptionalFieldsToSchema(userSchema, "user_id", "email", "connection_name")

	return &schema.Resource{
		Read:   readDataUser,
		Schema: userSchema,
	}
}

func newComputedDataSchema() map[string]*schema.Schema {
	userSchema := datasourceSchemaFromResourceSchema(newUser().Schema)
	return userSchema
}

func readDataUser(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)

	email := d.Get("email").(string)

	userID := d.Get("user_id").(string)

	connection := d.Get("connection_name").(string)

	if email == "" && userID == "" {
		return errors.New(`The argument "user_id" or "email" should be configured`)
	}

	var user *management.User

	var err error

	if email != "" {

		log.Printf("[DEBUG] Searching User by email %s", email)
		usersList, listByEmailError := api.User.ListByEmail(email)

		if listByEmailError != nil {
			return listByEmailError
		}

		if len(usersList) == 0 {
			return errors.New("No user found matching email " + email)
		}

		log.Printf("[DEBUG Found %d users with this email %s", len(usersList), email)

		if len(usersList) > 1 && connection != "" {

			log.Printf("[DEBUG] Selecting user by connection %s", connection)
			for _, u := range usersList {
				var found bool
				for _, i := range u.Identities {
					if found = i.GetConnection() == connection; found {
						break
					}
				}

				if found {
					user = u
					break
				}
			}

		} else {
			user = usersList[0]
		}

	} else {
		log.Printf("[DEBUG] Getting User by id %s", userID)
		user, err = api.User.Read(userID)
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

	d.SetId(*user.ID)

	d.Set("user_id", *user.ID)
	d.Set("email", user.Email)

	if connection == "" {
		d.Set("connection_name", user.Identities[0].Connection)
	}
	d.Set("email_verified", user.EmailVerified)

	d.Set("username", user.Username)
	d.Set("phone_number", user.PhoneNumber)
	d.Set("phone_verified", user.PhoneVerified)
	d.Set("created_at", user.CreatedAt.String())
	d.Set("updated_at", user.UpdatedAt.String())

	identities := make([]map[string]interface{}, len(user.Identities))

	for index, i := range user.Identities {
		newIdentity := make(map[string]interface{})
		newIdentity["connection"] = i.Connection
		newIdentity["user_id"] = i.UserID
		newIdentity["provider"] = i.Provider
		newIdentity["is_social"] = i.IsSocial
		identities[index] = newIdentity
	}

	d.Set("identities", identities)
	userMeta, err := structure.FlattenJsonToString(user.UserMetadata)
	if err != nil {
		return err
	}
	d.Set("user_metadata", userMeta)

	appMeta, err := structure.FlattenJsonToString(user.AppMetadata)
	if err != nil {
		return err
	}
	d.Set("app_metadata", appMeta)
	d.Set("picture", user.Picture)
	d.Set("name", user.Name)
	d.Set("nickname", user.Nickname)
	d.Set("blocked", user.Blocked)
	d.Set("family_name", user.FamilyName)
	d.Set("given_name", user.GivenName)

	return nil
}

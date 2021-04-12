package auth0

import (
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"gopkg.in/auth0.v5"
	"gopkg.in/auth0.v5/management"
)

func newPrompt() *schema.Resource {

	return &schema.Resource{

		Create: createPrompt,
		Read:   readPrompt,
		Update: updatePrompt,
		Delete: deletePrompt,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"universal_login_experience": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"new", "classic",
				}, false),
			},
			"identifier_first": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func createPrompt(d *schema.ResourceData, m interface{}) error {
	d.SetId(resource.UniqueId())
	return updatePrompt(d, m)
}

func readPrompt(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	p, err := api.Prompt.Read()
	if err != nil {
		if mErr, ok := err.(management.Error); ok {
			if mErr.Status() == http.StatusNotFound {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	d.Set("universal_login_experience", p.UniversalLoginExperience)
	d.Set("identifier_first", p.IdentifierFirst)
	return nil
}

func updatePrompt(d *schema.ResourceData, m interface{}) error {
	p := buildPrompt(d)
	api := m.(*management.Management)
	err := api.Prompt.Update(p)
	if err != nil {
		return err
	}
	return readPrompt(d, m)
}

func deletePrompt(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}

func buildPrompt(d *schema.ResourceData) *management.Prompt {
	return &management.Prompt{
		UniversalLoginExperience: auth0.StringValue(String(d, "universal_login_experience")),
		IdentifierFirst:          Bool(d, "identifier_first"),
	}
}

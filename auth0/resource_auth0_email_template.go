package auth0

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"github.com/yieldr/go-auth0/management"
)

func newEmailTemplate() *schema.Resource {
	return &schema.Resource{

		Create: createEmailTemplate,
		Read:   readEmailTemplate,
		Update: updateEmailTemplate,
		Delete: deleteEmailTemplate,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"template": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"verify_email",
					"reset_email",
					"welcome_email",
					"blocked_account",
					"stolen_credentials",
					"enrollment_email",
					"change_password",
					"password_reset",
					"mfa_oob_code",
				}, true),
			},
			"body": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"from": {
				Type:     schema.TypeString,
				Required: true,
			},
			"result_url": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subject": {
				Type:     schema.TypeString,
				Required: true,
			},
			"syntax": {
				Type:     schema.TypeString,
				Required: true,
			},
			"url_lifetime_in_seconds": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Required: true,
			},
		},
	}
}

func createEmailTemplate(d *schema.ResourceData, m interface{}) error {
	e := buildEmailTemplate(d)
	api := m.(*management.Management)
	if err := api.EmailTemplate.Update(e.Template, e); err != nil {
		return err
	}
	d.SetId(e.Template)
	return nil
}

func readEmailTemplate(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	e, err := api.EmailTemplate.Read(d.Id())
	if err != nil {
		return err
	}
	d.SetId(e.Template)
	d.Set("template", e.Template)
	d.Set("body", e.Body)
	d.Set("from", e.From)
	d.Set("result_url", e.ResultURL)
	d.Set("subject", e.Subject)
	d.Set("syntax", e.Syntax)
	d.Set("url_lifetime_in_seconds", e.URLLifetimeInSecoonds)
	d.Set("enabled", e.Enabled)
	return nil
}

func updateEmailTemplate(d *schema.ResourceData, m interface{}) error {
	e := buildEmailTemplate(d)
	api := m.(*management.Management)
	return api.EmailTemplate.Replace(e.Template, e)
}

func deleteEmailTemplate(d *schema.ResourceData, m interface{}) error {
	d.Set("enabled", false)
	return updateEmailTemplate(d, m)
}

func buildEmailTemplate(d *schema.ResourceData) *management.EmailTemplate {
	return &management.EmailTemplate{
		Template:              d.Get("template").(string),
		Body:                  d.Get("body").(string),
		From:                  d.Get("from").(string),
		ResultURL:             d.Get("result_url").(string),
		Subject:               d.Get("subject").(string),
		Syntax:                d.Get("syntax").(string),
		URLLifetimeInSecoonds: d.Get("url_lifetime_in_seconds").(int),
		Enabled:               d.Get("enabled").(bool),
	}
}

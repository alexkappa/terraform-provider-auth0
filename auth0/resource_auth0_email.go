package auth0

import (
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"gopkg.in/auth0.v4"
	"gopkg.in/auth0.v4/management"
)

func newEmail() *schema.Resource {
	return &schema.Resource{

		Create: createEmail,
		Read:   readEmail,
		Update: updateEmail,
		Delete: deleteEmail,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"default_from_address": {
				Type:     schema.TypeString,
				Required: true,
			},
			"credentials": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"api_user": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"api_key": {
							Type:      schema.TypeString,
							Optional:  true,
							Sensitive: true,
							ForceNew:  true,
						},
						"access_key_id": {
							Type:      schema.TypeString,
							Optional:  true,
							Sensitive: true,
							ForceNew:  true,
						},
						"secret_access_key": {
							Type:      schema.TypeString,
							Optional:  true,
							Sensitive: true,
							ForceNew:  true,
						},
						"region": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"domain": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"smtp_host": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"smtp_port": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"smtp_user": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"smtp_pass": {
							Type:      schema.TypeString,
							Optional:  true,
							Sensitive: true,
							ForceNew:  true,
						},
					},
				},
			},
		},
	}
}

func createEmail(d *schema.ResourceData, m interface{}) error {
	e := buildEmail(d)
	api := m.(*management.Management)
	if err := api.Email.Create(e); err != nil {
		return err
	}
	d.SetId(auth0.StringValue(e.Name))
	return readEmail(d, m)
}

func readEmail(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	e, err := api.Email.Read()
	if err != nil {
		if mErr, ok := err.(management.Error); ok {
			if mErr.Status() == http.StatusNotFound {
				d.SetId("")
				return nil
			}
		}
		return err
	}

	d.SetId(auth0.StringValue(e.Name))
	d.Set("name", e.Name)
	d.Set("enabled", e.Enabled)
	d.Set("default_from_address", e.DefaultFromAddress)

	if credentials := e.Credentials; credentials != nil {
		credentialsMap := make(map[string]interface{})
		credentialsMap["api_user"] = credentials.APIUser
		credentialsMap["api_key"] = d.Get("credentials.0.api_key")
		credentialsMap["access_key_id"] = d.Get("credentials.0.access_key_id")
		credentialsMap["secret_access_key"] = d.Get("credentials.0.secret_access_key")
		credentialsMap["region"] = credentials.Region
		credentialsMap["domain"] = credentials.Domain
		credentialsMap["smtp_host"] = credentials.SMTPHost
		credentialsMap["smtp_port"] = credentials.SMTPPort
		credentialsMap["smtp_user"] = credentials.SMTPUser
		credentialsMap["smtp_pass"] = d.Get("credentials.0.smtp_pass")
		d.Set("credentials", []map[string]interface{}{credentialsMap})
	}

	return nil
}

func updateEmail(d *schema.ResourceData, m interface{}) error {
	e := buildEmail(d)
	api := m.(*management.Management)
	err := api.Email.Update(e)
	if err != nil {
		return err
	}
	return readEmail(d, m)
}

func deleteEmail(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	err := api.Email.Delete()
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

func buildEmail(d *schema.ResourceData) *management.Email {
	e := &management.Email{
		Name:               String(d, "name"),
		Enabled:            Bool(d, "enabled"),
		DefaultFromAddress: String(d, "default_from_address"),
	}

	List(d, "credentials").Elem(func(d ResourceData) {
		// e.Credentials = buildEmailCredentials(v.(map[string]interface{}))
		e.Credentials = &management.EmailCredentials{
			APIUser:         String(d, "api_user"),
			APIKey:          String(d, "api_key"),
			AccessKeyID:     String(d, "access_key_id"),
			SecretAccessKey: String(d, "secret_access_key"),
			Region:          String(d, "region"),
			Domain:          String(d, "domain"),
			SMTPHost:        String(d, "smtp_host"),
			SMTPPort:        Int(d, "smtp_port"),
			SMTPUser:        String(d, "smtp_user"),
			SMTPPass:        String(d, "smtp_pass"),
		}
	})

	return e
}

func buildEmailCredentials(m map[string]interface{}) *management.EmailCredentials {
	return &management.EmailCredentials{
		APIUser:         String(MapData(m), "api_user"),
		APIKey:          String(MapData(m), "api_key"),
		AccessKeyID:     String(MapData(m), "access_key_id"),
		SecretAccessKey: String(MapData(m), "secret_access_key"),
		Region:          String(MapData(m), "region"),
		Domain:          String(MapData(m), "domain"),
		SMTPHost:        String(MapData(m), "smtp_host"),
		SMTPPort:        Int(MapData(m), "smtp_port"),
		SMTPUser:        String(MapData(m), "smtp_user"),
		SMTPPass:        String(MapData(m), "smtp_pass"),
	}
}

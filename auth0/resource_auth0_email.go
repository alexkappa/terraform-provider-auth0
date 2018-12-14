package auth0

import (
	"github.com/hashicorp/terraform/helper/schema"
	auth0 "github.com/yieldr/go-auth0"
	"github.com/yieldr/go-auth0/management"
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
							Type:     schema.TypeString,
							Optional: true,
						},
						"access_key_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"secret_access_key": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"region": {
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
							Type:     schema.TypeString,
							Optional: true,
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
	e, err := api.Email.Read(management.WithFields("name", "enabled", "default_from_address", "credentials"))
	if err != nil {
		return err
	}
	d.SetId(auth0.StringValue(e.Name))
	d.Set("name", e.Name)
	d.Set("enabled", e.Enabled)
	d.Set("default_from_address", e.DefaultFromAddress)

	if credentials := e.Credentials; credentials != nil {
		credentialsMap := make(map[string]interface{})
		credentialsMap["api_user"] = credentials.APIUser
		credentialsMap["api_key"] = credentials.APIKey
		credentialsMap["access_key_id"] = credentials.AccessKeyID
		credentialsMap["secret_access_key"] = credentials.SecretAccessKey
		credentialsMap["region"] = credentials.Region
		credentialsMap["smtp_host"] = credentials.SMTPHost
		credentialsMap["smtp_port"] = credentials.SMTPPort
		credentialsMap["smtp_user"] = credentials.SMTPUser
		credentialsMap["smtp_pass"] = credentials.SMTPPass
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
	return api.Email.Delete()
}

func buildEmail(d *schema.ResourceData) *management.Email {
	e := &management.Email{
		Name:               String(d, "name"),
		Enabled:            Bool(d, "enabled"),
		DefaultFromAddress: String(d, "default_from_address"),
	}

	List(d, "credentials").First(func(v interface{}) {
		e.Credentials = buildEmailCredentials(v.(map[string]interface{}))
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
		SMTPHost:        String(MapData(m), "smtp_host"),
		SMTPPort:        Int(MapData(m), "smtp_port"),
		SMTPUser:        String(MapData(m), "smtp_user"),
		SMTPPass:        String(MapData(m), "smtp_pass"),
	}
}

package auth0

import (
	"net/http"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"gopkg.in/auth0.v3"
	"gopkg.in/auth0.v3/management"
)

var hookNameRegexp = regexp.MustCompile("^[^\\s-][\\w -]+[^\\s-]$")

func newHook() *schema.Resource {
	return &schema.Resource{

		Create: createHook,
		Read:   readHook,
		Update: updateHook,
		Delete: deleteHook,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringMatch(
					hookNameRegexp,
					"Can only contain alphanumeric characters, spaces and '-'. "+
						"Can neither start nor end with '-' or spaces."),
			},
			"script": {
				Type:     schema.TypeString,
				Required: true,
			},
			"trigger_id": {
				Type: schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{
					"credentials-exchange", "pre-user-registration",
					"post-user-registration", "post-change-password",
				}, false),
				ForceNew: true,
				Optional: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func createHook(d *schema.ResourceData, m interface{}) error {
	c := buildHook(d)
	api := m.(*management.Management)
	if err := api.Hook.Create(c); err != nil {
		return err
	}
	d.SetId(auth0.StringValue(c.ID))
	return readHook(d, m)
}

func readHook(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	c, err := api.Hook.Read(d.Id())
	if err != nil {
		if mErr, ok := err.(management.Error); ok {
			if mErr.Status() == http.StatusNotFound {
				d.SetId("")
				return nil
			}
		}
		return err
	}

	d.Set("name", c.Name)
	d.Set("script", c.Script)
	d.Set("trigger_id", c.TriggerID)
	d.Set("enabled", c.Enabled)
	return nil
}

func updateHook(d *schema.ResourceData, m interface{}) error {
	c := buildHook(d)
	api := m.(*management.Management)
	err := api.Hook.Update(d.Id(), c)
	if err != nil {
		return err
	}
	return readHook(d, m)
}

func deleteHook(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	err := api.Hook.Delete(d.Id())
	if err != nil {
		if mErr, ok := err.(management.Error); ok {
			if mErr.Status() == http.StatusNotFound {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	return err
}

func buildHook(d *schema.ResourceData) *management.Hook {
	return &management.Hook{
		Name:      String(d, "name"),
		Script:    String(d, "script"),
		TriggerID: String(d, "trigger_id"),
		Enabled:   Bool(d, "enabled"),
	}
}

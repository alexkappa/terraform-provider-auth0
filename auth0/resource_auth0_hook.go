package auth0

import (
	"net/http"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"gopkg.in/auth0.v5"
	"gopkg.in/auth0.v5/management"
)

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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateHookNameFunc(),
				Description:  "Name of this hook",
			},
			"dependencies": {
				Type:        schema.TypeMap,
				Elem:        schema.TypeString,
				Required:    false,
				Optional:    true,
				Description: "Dependencies of this hook used by webtask server",
			},
			"script": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Code to be executed when this hook runs",
			},
			"trigger_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"credentials-exchange",
					"pre-user-registration",
					"post-user-registration",
					"post-change-password",
					"send-phone-message",
				}, false),
				Description: "Execution stage of this rule. Can be " +
					"credentials-exchange, pre-user-registration, " +
					"post-user-registration, post-change-password" +
					", or send-phone-message",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether the hook is enabled, or disabled",
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
	d.Set("dependencies", c.Dependencies)
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
	deps := Map(d, "dependencies")
	return &management.Hook{
		Name:         String(d, "name"),
		Script:       String(d, "script"),
		TriggerID:    String(d, "trigger_id", IsNewResource()),
		Enabled:      Bool(d, "enabled"),
		Dependencies: &deps,
	}
}

func validateHookNameFunc() schema.SchemaValidateFunc {
	return validation.StringMatch(
		regexp.MustCompile("^[^\\s-][\\w -]+[^\\s-]$"),
		"Can only contain alphanumeric characters, spaces and '-'. Can neither start nor end with '-' or spaces.")
}

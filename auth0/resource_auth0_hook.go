package auth0

import (
	"net/http"
	"reflect"
	"regexp"
	"sort"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"gopkg.in/auth0.v4"
	"gopkg.in/auth0.v4/management"
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
				Description: "Name of this hook",
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
				}, false),
				Description: "Execution stage of this rule. Can be " +
					"credentials-exchange, pre-user-registration, " +
					"post-user-registration, post-change-password" +
					", or send-phone-message",
			},
			"secrets": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "The secrets associated with the hook",
				Elem:        schema.TypeString,
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
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
	if err := upsertHookSecrets(d, m); err != nil {
		return err
	}
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
	if err = upsertHookSecrets(d, m); err != nil {
		return err
	}
	return readHook(d, m)
}

func upsertHookSecrets(d *schema.ResourceData, m interface{}) error {
	if d.IsNewResource() || d.HasChange("secrets") {
		secrets := MapIfExists(d, "secrets")
		api := m.(*management.Management)
		hookSecrets := management.HookSecrets{}
		for key, value := range secrets {
			if strVal, ok := value.(string); ok {
				hookSecrets[key] = strVal
			}
		}
		if !d.IsNewResource() {
			if secretsBefore, err := api.Hook.Secrets(d.Id()); err == nil && secretsBefore != nil {
				keysBefore := secretsBefore.Keys()
				if len(secrets) > 0 && len(keysBefore) > 0 {
					sort.Strings(keysBefore)
					i := 0
					keysNow := make([]string, len(secrets))
					for k := range secrets {
						keysNow[i] = k
						i++
					}
					sort.Strings(keysNow)
					if reflect.DeepEqual(keysBefore, keysNow) {
						// can only update secrets if the keys are unchanged (i.e. you can't add a key)
						return api.Hook.UpdateSecrets(d.Id(), &hookSecrets)
					}
				}
				if len(keysBefore) > 0 {
					// otherwise remove all the secrets before we try to create
					if err := api.Hook.RemoveSecrets(d.Id(), keysBefore...); err != nil {
						return err
					}
				}
			}
		}
		if len(secrets) > 0 {
			return api.Hook.CreateSecrets(d.Id(), &hookSecrets)
		}
	}
	return nil
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
		Name:      StringIfExists(d, "name"),
		Script:    StringIfExists(d, "script"),
		Enabled:   BoolIfExists(d, "enabled"),
		TriggerID: String(d, "trigger_id"),
	}
}

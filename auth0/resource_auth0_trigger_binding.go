package auth0

import (
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"gopkg.in/auth0.v5"
	"gopkg.in/auth0.v5/management"
)

func newTriggerBinding() *schema.Resource {
	return &schema.Resource{

		Create: createTriggerBinding,
		Read:   readTriggerBinding,
		Update: updateTriggerBinding,
		Delete: deleteTriggerBinding,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"trigger": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"post-login",
					"credentials-exchange",
					"pre-user-registration",
					"post-user-registration",
					"post-change-password",
					"send-phone-message",
					"iga-approval",
					"iga-certification",
					"iga-fulfillment-assignment",
					"iga-fulfillment-execution",
				}, false),
				Description: "The id of the trigger to bind with",
			},
			"actions": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Trigger ID",
						},
						"display_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The name of an action",
						},
					},
				},
				Description: "The actions bound to this trigger",
			},
		},
	}
}

func createTriggerBinding(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	id := d.Get("trigger").(string)
	b := expandTriggerBindings(d)
	err := api.Action.UpdateBindings(id, b)
	if err != nil {
		return err
	}
	d.SetId(id)
	return readTriggerBinding(d, m)
}

func readTriggerBinding(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	b, err := api.Action.Bindings(d.Id())
	if err != nil {
		if mErr, ok := err.(management.Error); ok {
			if mErr.Status() == http.StatusNotFound {
				d.SetId("")
				return nil
			}
		}
		return err
	}

	d.Set("actions", flattenTriggerBindingActions(b.Bindings))

	return nil
}

func updateTriggerBinding(d *schema.ResourceData, m interface{}) error {
	b := expandTriggerBindings(d)
	api := m.(*management.Management)
	err := api.Action.UpdateBindings(d.Id(), b)
	if err != nil {
		return err
	}
	return readTriggerBinding(d, m)
}

func deleteTriggerBinding(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	if err := api.Action.UpdateBindings(d.Id(), []*management.ActionBinding{}); err != nil {
		if mErr, ok := err.(management.Error); ok {
			if mErr.Status() == http.StatusNotFound {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	return nil
}

func expandTriggerBindings(d *schema.ResourceData) (b []*management.ActionBinding) {
	List(d, "actions").Elem(func(d ResourceData) {
		b = append(b, &management.ActionBinding{
			Ref: &management.ActionBindingReference{
				Type:  auth0.String("action_id"),
				Value: String(d, "id"),
			},
			DisplayName: String(d, "display_name"),
		})
	})
	return
}

func flattenTriggerBindingActions(bindings []*management.ActionBinding) (r []interface{}) {
	for _, b := range bindings {
		r = append(r, map[string]interface{}{
			"id":           b.Action.GetID(),
			"display_name": b.GetDisplayName(),
		})
	}
	return
}

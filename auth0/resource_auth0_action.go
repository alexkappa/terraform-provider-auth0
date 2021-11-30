package auth0

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/alexkappa/terraform-provider-auth0/auth0/internal/hash"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"gopkg.in/auth0.v5/management"
)

func newAction() *schema.Resource {
	return &schema.Resource{

		Create: createAction,
		Read:   readAction,
		Update: updateAction,
		Delete: deleteAction,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of an action",
			},
			"supported_triggers": {
				Type:     schema.TypeList,
				Required: true,
				MinItems: 1,
				MaxItems: 1, // NOTE: Changes must be made together with expandAction()
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Trigger ID",
						},
						"version": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Trigger version",
						},
					},
				},
				Description: "List of triggers that this action supports. At " +
					"this time, an action can only target a single trigger at" +
					" a time",
			},
			"code": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The source code of the action.",
			},
			"dependencies": {
				Type:     schema.TypeSet,
				Optional: true,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Dependency name. For example lodash",
						},
						"version": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Dependency version. For example `latest` or `4.17.21`",
						},
					},
				},
				Set:         hash.StringKey("name"),
				Description: "List of third party npm modules, and their versions, that this action depends on",
			},
			"runtime": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ValidateFunc: validation.StringInSlice([]string{
					"node12",
					"node16",
				}, false),
				Description: "The Node runtime. For example `node16`, defaults to `node12`",
			},
			"secrets": {
				Type:     schema.TypeList,
				Optional: true,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Secret name",
						},
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							Sensitive:   true,
							Description: "Secret value",
						},
					},
				},
				Description: "List of secrets that are included in an action or a version of an action",
			},
			"deploy": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				Description: "Deploying an action will create a new immutable" +
					" version of the action. If the action is currently bound" +
					" to a trigger, then the system will begin executing the " +
					"newly deployed version of the action immediately",
			},
			"version_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Version ID of the action. This value is available if `deploy` is set to true",
			},
		},
	}
}

func createAction(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	a := expandAction(d)
	err := api.Action.Create(a)
	if err != nil {
		return err
	}
	d.SetId(a.GetID())

	d.Partial(true)
	err = deployAction(d, m)
	if err != nil {
		return err
	}
	d.Partial(false)

	return readAction(d, m)
}

func readAction(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	a, err := api.Action.Read(d.Id())
	if err != nil {
		if mErr, ok := err.(management.Error); ok {
			if mErr.Status() == http.StatusNotFound {
				d.SetId("")
				return nil
			}
		}
		return err
	}

	d.Set("name", a.Name)
	d.Set("supported_triggers", flattenActionTriggers(a.SupportedTriggers))
	d.Set("code", a.Code)
	d.Set("dependencies", flattenActionDependencies(a.Dependencies))
	d.Set("runtime", a.Runtime)

	if a.DeployedVersion != nil {
		d.Set("version_id", a.DeployedVersion.GetID())
	}

	return nil
}

func updateAction(d *schema.ResourceData, m interface{}) error {
	a := expandAction(d)
	api := m.(*management.Management)
	err := api.Action.Update(d.Id(), a)
	if err != nil {
		return err
	}
	d.Partial(true)
	err = deployAction(d, m)
	if err != nil {
		return err
	}
	d.Partial(false)
	return readAction(d, m)
}

func deployAction(d *schema.ResourceData, m interface{}) error {

	if d.Get("deploy").(bool) == true {

		api := m.(*management.Management)

		err := resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {

			a, err := api.Action.Read(d.Id())
			if err != nil {
				return resource.NonRetryableError(err)
			}

			if strings.ToLower(a.GetStatus()) != "built" {
				return resource.RetryableError(
					fmt.Errorf(`Expected action status %q to equal "built"`, a.GetStatus()),
				)
			}

			return nil
		})
		if err != nil {
			return fmt.Errorf("Action never reached built state. %w", err)
		}

		v, err := api.Action.Deploy(d.Id())
		if err != nil {
			return err
		}

		d.Set("version_id", v.GetID())
	}
	return nil
}

func deleteAction(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	if err := api.Action.Delete(d.Id()); err != nil {
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

func expandAction(d *schema.ResourceData) *management.Action {

	a := &management.Action{
		Name:    String(d, "name"),
		Code:    String(d, "code"),
		Runtime: String(d, "runtime"),
	}

	List(d, "supported_triggers").Elem(func(d ResourceData) {
		a.SupportedTriggers = []*management.ActionTrigger{
			{
				ID:      String(d, "id"),
				Version: String(d, "version"),
			},
		}
	})

	Set(d, "dependencies").Elem(func(d ResourceData) {
		a.Dependencies = append(a.Dependencies, &management.ActionDependency{
			Name:    String(d, "name"),
			Version: String(d, "version"),
		})
	})

	Set(d, "secrets").Elem(func(d ResourceData) {
		a.Secrets = append(a.Secrets, &management.ActionSecret{
			Name:  String(d, "name"),
			Value: String(d, "value"),
		})
	})

	return a
}

func flattenActionTriggers(triggers []*management.ActionTrigger) (ret []interface{}) {
	for _, trigger := range triggers {
		ret = append(ret, map[string]interface{}{
			"id":      trigger.ID,
			"version": trigger.Version,
		})
	}
	return
}

func flattenActionDependencies(dependencies []*management.ActionDependency) (ret []interface{}) {
	for _, dependency := range dependencies {
		ret = append(ret, map[string]interface{}{
			"name":    dependency.Name,
			"version": dependency.Version,
		})
	}
	return
}

func flattenActionSecrets(secrets []*management.ActionSecret) (ret []interface{}) {
	for _, secret := range secrets {
		ret = append(ret, map[string]interface{}{
			"name":  secret.Name,
			"value": secret.Value,
		})
	}
	return
}

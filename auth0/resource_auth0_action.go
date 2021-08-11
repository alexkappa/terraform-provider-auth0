package auth0

import (
	"net/http"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"gopkg.in/auth0.v5"
	"gopkg.in/auth0.v5/management"
)

var actionNameRegexp = regexp.MustCompile("^[^\\s-][\\w -]+[^\\s-]$")

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
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringMatch(
					actionNameRegexp,
					"Can only contain alphanumeric characters, spaces and '-'. "+
						"Can neither start nor end with '-' or spaces."),
			},
			"supported_triggers": {
				Type:     schema.TypeList,
				Optional: true,
			},
			"code": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dependencies": {
				Type: schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"version": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
				Optional: true,
				Computed: true,
			},
			"runtime": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"secrets": {
				Type: schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
				Computed: true,
				Optional: true,
			},
		},
	}
}

func createAction(d *schema.ResourceData, m interface{}) error {
	c := buildAction(d)
	api := m.(*management.Management)
	if err := api.Action.Create(c); err != nil {
		return err
	}
	d.SetId(auth0.StringValue(c.ID))
	return readAction(d, m)
}

func readAction(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	c, err := api.Action.Read(d.Id())
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
	d.Set("supported_triggers", c.SupportedTriggers)
	d.Set("code", c.Code)
	d.Set("dependencies", c.Dependencies)
	d.Set("secrets", c.Secrets)

	return nil
}

func updateAction(d *schema.ResourceData, m interface{}) error {
	c := buildAction(d)
	api := m.(*management.Management)
	err := api.Action.Update(d.Id(), c)
	if err != nil {
		return err
	}
	return readAction(d, m)
}

func deleteAction(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	err := api.Action.Delete(d.Id())
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

func buildAction(d *schema.ResourceData) *management.Action {
	return &management.Action{
		Name:              String(d, "name"),
		SupportedTriggers: expandActionTrigger(d),
		Code:              String(d, "code"),
		Dependencies:      expandActionDependency(d),
		Secrets:           expandActionSecret(d),
	}
}

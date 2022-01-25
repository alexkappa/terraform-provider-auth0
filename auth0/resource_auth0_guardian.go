package auth0

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"gopkg.in/auth0.v5/management"
)

func newGuardian() *schema.Resource {
	return &schema.Resource{

		Create: createGuardian,
		Read:   readGuardian,
		Update: updateGuardian,
		Delete: deleteGuardian,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"policy": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"all-applications",
					"confidence-score",
					"never",
				}, false),
			},
			"phone": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				MinItems: 0,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"provider": {
							Type:     schema.TypeString,
							Required: true,
							ValidateFunc: validation.StringInSlice([]string{
								"auth0",
								"twilio",
								"phone-message-hook",
							}, false),
						},
						"message_types": {
							Type:     schema.TypeList,
							Required: true,

							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"options": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enrollment_message": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"verification_message": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"from": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"messaging_service_sid": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"auth_token": {
										Type:      schema.TypeString,
										Sensitive: true,
										Optional:  true,
									},
									"sid": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			"email": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func createGuardian(d *schema.ResourceData, m interface{}) error {
	d.SetId(resource.UniqueId())
	return updateGuardian(d, m)
}

func deleteGuardian(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	if err := api.Guardian.MultiFactor.Phone.Enable(false); err != nil {
		return err
	}
	if err := api.Guardian.MultiFactor.Email.Enable(false); err != nil {
		return err
	}
	d.SetId("")
	return nil
}
func updateGuardian(d *schema.ResourceData, m interface{}) (err error) {
	api := m.(*management.Management)

	if d.HasChange("policy") {
		p := d.Get("policy").(string)
		if p == "never" {
			//Passing empty array to set it to the "never" policy.
			err = api.Guardian.MultiFactor.UpdatePolicy(&management.MultiFactorPolicies{})
		} else {
			err = api.Guardian.MultiFactor.UpdatePolicy(&management.MultiFactorPolicies{p})
		}
	}
	if err := updatePhoneFactor(d, api); err != nil {
		return err
	}
	if err := updateEmailFactor(d, api); err != nil {
		return err
	}
	return readGuardian(d, m)
}

func updatePhoneFactor(d *schema.ResourceData, api *management.Management) error {
	ok, err := factorShouldBeUpdated(d, "phone")
	if err != nil {
		return err
	}
	if ok {
		if err := api.Guardian.MultiFactor.Phone.Enable(true); err != nil {
			return err
		}
		return configurePhone(d, api)
	}
	return api.Guardian.MultiFactor.Phone.Enable(false)
}

func updateEmailFactor(d *schema.ResourceData, api *management.Management) error {
	if changed := d.HasChange("email"); changed {
		enabled := d.Get("email").(bool)
		return api.Guardian.MultiFactor.Email.Enable(enabled)
	}
	return nil
}

func configurePhone(d *schema.ResourceData, api *management.Management) (err error) {
	md := make(MapData)
	List(d, "phone").Elem(func(d ResourceData) {
		md.Set("provider", String(d, "provider", HasChange()))
		md.Set("message_types", Slice(d, "message_types", HasChange()))
		md.Set("options", List(d, "options"))
		switch *String(d, "provider") {
		case "twilio":
			err = updateTwilioOptions(md["options"].(Iterator), api)
		case "auth0":
			err = updateAuth0Options(md["options"].(Iterator), api)
		}
	})

	if s, ok := md.GetOk("provider"); ok {
		if err := api.Guardian.MultiFactor.Phone.UpdateProvider(&management.MultiFactorProvider{Provider: s.(*string)}); err != nil {
			return err
		}
	}

	mtypes := typeAssertToStringArray(Slice(md, "message_types"))
	if mtypes != nil {
		if err := api.Guardian.MultiFactor.Phone.UpdateMessageTypes(&management.PhoneMessageTypes{MessageTypes: mtypes}); err != nil {
			return err
		}
	}

	return err
}

func updateAuth0Options(opts Iterator, api *management.Management) (err error) {
	opts.Elem(func(d ResourceData) {
		err = api.Guardian.MultiFactor.SMS.UpdateTemplate(&management.MultiFactorSMSTemplate{
			EnrollmentMessage:   String(d, "enrollment_message"),
			VerificationMessage: String(d, "verification_message"),
		})
	})
	if err != nil {
		return err
	}
	return nil
}

func updateTwilioOptions(opts Iterator, api *management.Management) error {
	md := make(map[string]*string)
	opts.Elem(func(d ResourceData) {
		md["sid"] = String(d, "sid")
		md["auth_token"] = String(d, "auth_token")
		md["from"] = String(d, "from")
		md["messaging_service_sid"] = String(d, "messaging_service_sid")
		md["enrollment_message"] = String(d, "enrollment_message")
		md["verification_message"] = String(d, "verification_message")
	})

	err := api.Guardian.MultiFactor.SMS.UpdateTwilio(&management.MultiFactorProviderTwilio{
		From:                md["from"],
		MessagingServiceSid: md["messaging_service_sid"],
		AuthToken:           md["auth_token"],
		SID:                 md["sid"],
	})
	if err != nil {
		return err
	}
	err = api.Guardian.MultiFactor.SMS.UpdateTemplate(&management.MultiFactorSMSTemplate{
		EnrollmentMessage:   md["enrollment_message"],
		VerificationMessage: md["verification_message"],
	})
	if err != nil {
		return err
	}
	return nil
}

func readGuardian(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	mt, err := api.Guardian.MultiFactor.Phone.MessageTypes()
	if err != nil {
		return err
	}
	phoneData := make(map[string]interface{})
	phoneData["message_types"] = mt.MessageTypes
	prv, err := api.Guardian.MultiFactor.Phone.Provider()
	if err != nil {
		return err
	}
	phoneData["provider"] = prv.Provider

	p, err := api.Guardian.MultiFactor.Policy()
	if err != nil {
		return err
	}
	if len(*p) == 0 {
		d.Set("policy", "never")
	} else {
		err = d.Set("policy", (*p)[0])
	}
	var md map[string]interface{}
	switch *prv.Provider {
	case "twilio":
		md, err = flattenTwilioOptions(api)
	case "auth0":
		md, err = flattenAuth0Options(api)
	}
	if err != nil {
		return err
	}

	ok, err := factorShouldBeUpdated(d, "phone")
	if err != nil {
		return err
	}
	if ok {
		phoneData["options"] = []interface{}{md}
		err = d.Set("phone", []interface{}{phoneData})
	} else {
		d.Set("phone", nil)
	}
	if err != nil {
		return err
	}

	factors, err := api.Guardian.MultiFactor.List()
	if err != nil {
		return err
	}
	for _, v := range factors {
		if v.Name != nil && *v.Name == "email" {
			d.Set("email", v.Enabled)
		}
	}
	return nil
}

func hasBlockPresentInNewState(d *schema.ResourceData, factor string) bool {
	if ok := d.HasChange(factor); ok {
		_, n := d.GetChange(factor)
		newState := n.([]interface{})
		return len(newState) > 0
	}
	return false
}

func flattenAuth0Options(api *management.Management) (map[string]interface{}, error) {
	md := make(map[string]interface{})
	t, err := api.Guardian.MultiFactor.SMS.Template()
	if err != nil {
		return nil, err
	}
	md["enrollment_message"] = t.EnrollmentMessage
	md["verification_message"] = t.VerificationMessage
	return md, nil
}

func flattenTwilioOptions(api *management.Management) (map[string]interface{}, error) {
	md := make(map[string]interface{})
	t, err := api.Guardian.MultiFactor.SMS.Template()
	if err != nil {
		return nil, err
	}
	md["enrollment_message"] = t.EnrollmentMessage
	md["verification_message"] = t.VerificationMessage
	tw, err := api.Guardian.MultiFactor.SMS.Twilio()
	if err != nil {
		return nil, err
	}
	md["auth_token"] = tw.AuthToken
	md["from"] = tw.From
	md["messaging_service_sid"] = tw.MessagingServiceSid
	md["sid"] = tw.SID
	return md, nil
}

func typeAssertToStringArray(from []interface{}) *[]string {
	length := len(from)
	if length < 1 {
		return nil
	}
	stringArray := make([]string, length)
	for i, v := range from {
		stringArray[i] = v.(string)
	}
	return &stringArray
}

// Determines if the factor should be updated. This depends on if it is in the state, if it is about to be added to the state.
func factorShouldBeUpdated(d *schema.ResourceData, factor string) (bool, error) {
	_, ok := d.GetOk(factor)
	return ok || hasBlockPresentInNewState(d, factor), nil
}

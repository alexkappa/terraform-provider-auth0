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
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:     schema.TypeBool,
							Required: true,
						},
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
										Type:     schema.TypeString,
										Optional: true,
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
		},
	}
}
func createGuardian(d *schema.ResourceData, m interface{}) error {
	d.SetId(resource.UniqueId())
	return updateGuardian(d, m)
}

func deleteGuardian(d *schema.ResourceData, m interface{}) error {
	return updateGuardian(d, m)
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
	//TODO: Extend for other MFA types
	if _, ok := d.GetOk("phone"); ok {
		err = configurePhone(d, api)
	}
	if err != nil {
		return err
	}
	return readGuardian(d, m)
}

func configurePhone(d *schema.ResourceData, api *management.Management) (err error) {
	md := make(MapData)
	List(d, "phone").Elem(func(d ResourceData) {
		md.Set("provider", String(d, "provider", HasChange()))
		md.Set("message_types", List(d, "message_types", HasChange()).List())
		md.Set("enabled", Bool(d, "enabled", HasChange()))
		md.Set("options", List(d, "options"))
		switch *String(d, "provider") {
		case "twilio":
			err = updateTwilioOptions(md["options"].(Iterator), api)
		case "auth0":
			err = updateAuth0Options(md["options"].(Iterator), api)
		}
	})

	if p, ok := md.GetOk("provider"); ok {
		if err := api.Guardian.MultiFactor.Phone.UpdateProvider(&management.MultiFactorProvider{Provider: p.(*string)}); err != nil {
			return err
		}
	}

	if mt, ok := md.GetOk("message_types"); ok {
		if mtypes, ok := mt.([]string); ok {
			if err := api.Guardian.MultiFactor.Phone.UpdateMessageTypes(&management.PhoneMessageTypes{MessageTypes: &mtypes}); err != nil {
				return err
			}
		}
	}

	if enabled, ok := md.GetOk("enabled"); ok {
		if err := api.Guardian.MultiFactor.Phone.Enable(*enabled.(*bool)); err != nil {
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
	if mt, err := api.Guardian.MultiFactor.Phone.MessageTypes(); err != nil {
		d.Set("message_types", mt)
	} else {
		return err
	}

	if prv, err := api.Guardian.MultiFactor.Phone.Provider(); err != nil {
		d.Set("provider", prv)
	} else {
		return err
	}
	if p, err := api.Guardian.MultiFactor.Policy(); err != nil {
		d.Set("policy", p)
	} else {
		return err
	}

	api.Guardian.MultiFactor.SMS.Template()
	api.Guardian.MultiFactor.SMS.Twilio()
	md := make(map[string]interface{})
	if t, err := api.Guardian.MultiFactor.SMS.Template(); err != nil {
		md["enrollment_message"] = t.EnrollmentMessage
		md["verification_message"] = t.VerificationMessage
	} else {
		return err
	}
	if tw, err := api.Guardian.MultiFactor.SMS.Twilio(); err != nil {
		md["auth_token"] = tw.AuthToken
		md["from"] = tw.From
		md["messaging_service_sid"] = tw.MessagingServiceSid
		md["sid"] = tw.SID
	} else {
		return err
	}
	d.Set("options", md)
	return nil
}

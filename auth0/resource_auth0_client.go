package auth0

import (
	"github.com/90poe/go-auth0/management"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func newClient() *schema.Resource {
	return &schema.Resource{

		Create: createClient,
		Read:   readClient,
		Update: updateClient,
		Delete: deleteClient,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"client_secret": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
			"app_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"logo_uri": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_first_party": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"oidc_conformant": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"callbacks": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"allowed_logout_urls": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"allowed_origins": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"web_origins": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"jwt_configuration": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"lifetime_in_seconds": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"secret_encoded": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"scopes": {
							Type:     schema.TypeMap,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"alg": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"encryption_key": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeMap},
				Optional: true,
			},
			"sso": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"sso_disabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"cross_origin_auth": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"cross_origin_loc": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"custom_login_page_on": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"custom_login_page": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"custom_login_page_preview": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"form_template": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"addons": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},
			"token_endpoint_auth_method": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"none",
					"client_secret_post",
					"client_secret_basic",
				}, false),
			},
			"client_metadata": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     schema.TypeString,
			},
			"mobile": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"android": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"app_package_name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"sha256_cert_fingerprints": {
										Type:     schema.TypeList,
										Optional: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
									},
								},
							},
						},
						"ios": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"team_id": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"app_bundle_identifier": {
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

func createClient(d *schema.ResourceData, m interface{}) error {
	c := buildClient(d)
	api := m.(*management.Management)
	if err := api.Client.Create(c); err != nil {
		return err
	}
	d.SetId(c.ClientID)
	d.Set("client_id", c.ClientID)
	d.Set("client_secret", c.ClientSecret)
	return nil
}

func readClient(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	c, err := api.Client.Read(d.Id())
	if err != nil {
		return err
	}
	d.SetId(c.ClientID)
	d.Set("client_id", c.ClientID)
	d.Set("client_secret", c.ClientSecret)
	return nil
}

func updateClient(d *schema.ResourceData, m interface{}) error {
	c := buildClient(d)
	api := m.(*management.Management)
	err := api.Client.Update(d.Id(), c)
	if err != nil {
		return err
	}
	d.Set("client_id", c.ClientID)
	d.Set("client_secret", c.ClientSecret)
	return nil
}

func deleteClient(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	return api.Client.Delete(d.Id())
}

func buildClient(d *schema.ResourceData) *management.Client {

	c := &management.Client{
		Name:                    d.Get("name").(string),
		Description:             d.Get("description").(string),
		AppType:                 d.Get("app_type").(string),
		LogoURI:                 d.Get("logo_uri").(string),
		IsFirstParty:            d.Get("is_first_party").(bool),
		OIDCConformant:          d.Get("oidc_conformant").(bool),
		Callbacks:               d.Get("callbacks").([]interface{}),
		AllowedLogoutURLs:       d.Get("allowed_logout_urls").([]interface{}),
		AllowedOrigins:          d.Get("allowed_origins").([]interface{}),
		WebOrigins:              d.Get("web_origins").([]interface{}),
		SSO:                     d.Get("sso").(bool),
		SSODisabled:             d.Get("sso_disabled").(bool),
		CrossOriginAuth:         d.Get("cross_origin_auth").(bool),
		CrossOriginLocation:     d.Get("cross_origin_loc").(string),
		CustomLoginPageOn:       d.Get("custom_login_page_on").(bool),
		CustomLoginPage:         d.Get("custom_login_page").(string),
		CustomLoginPagePreview:  d.Get("custom_login_page_preview").(string),
		FormTemplate:            d.Get("form_template").(string),
		TokenEndpointAuthMethod: d.Get("token_endpoint_auth_method").(string),
	}

	if v, ok := d.GetOk("jwt_configuration"); ok {
		vL := v.([]interface{})
		for _, v := range vL {
			jwtC := v.(map[string]interface{})

			c.JWTConfiguration = &management.ClientJWTConfiguration{
				LifetimeInSeconds: jwtC["lifetime_in_seconds"].(int),
				Scopes:            jwtC["scopes"],
				Algorithm:         jwtC["alg"].(string),
			}
		}
	}

	if v, ok := d.GetOk("encryption_key"); ok {

		c.EncryptionKey = make(map[string]string)

		for _, item := range v.(*schema.Set).List() {
			for key, val := range item.(map[string]string) {
				c.EncryptionKey[key] = val
			}
		}
	}

	if v, ok := d.GetOk("addons"); ok {

		c.Addons = make(map[string]interface{})

		for _, item := range v.(*schema.Set).List() {
			for key, val := range item.(map[string]interface{}) {
				c.Addons[key] = val
			}
		}
	}

	if v, ok := d.GetOk("client_metadata"); ok {

		c.ClientMetadata = make(map[string]string)

		for key, val := range v.(map[string]interface{}) {
			c.ClientMetadata[key] = val.(string)
		}

	}

	if v, ok := d.GetOk("mobile"); ok {

		c.Mobile = make(map[string]interface{})

		for _, item := range v.([]interface{}) {

			for key, val := range item.(map[string]interface{}) {

				for _, valItem := range val.([]interface{}) {
					c.Mobile[key] = valItem
				}
			}
		}

	}

	return c
}

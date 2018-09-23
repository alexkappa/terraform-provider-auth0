package auth0

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	auth0 "github.com/yieldr/go-auth0"
	"github.com/yieldr/go-auth0/management"
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
				Computed: true,
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
			"grant_types": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
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
				MinItems: 1,
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
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
				Computed: true,
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
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},
			"token_endpoint_auth_method": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
	d.SetId(auth0.StringValue(c.ClientID))
	return readClient(d, m)
}

func readClient(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	c, err := api.Client.Read(d.Id())
	if err != nil {
		return err
	}

	d.Set("client_id", c.ClientID)
	d.Set("client_secret", c.ClientSecret)
	d.Set("name", c.Name)
	d.Set("description", c.Description)
	d.Set("app_type", c.AppType)
	d.Set("logo_uri", c.LogoURI)
	d.Set("is_first_party", c.IsFirstParty)
	d.Set("oidc_conformant", c.OIDCConformant)
	d.Set("callbacks", c.Callbacks)
	d.Set("allowed_logout_urls", c.AllowedLogoutURLs)
	d.Set("allowed_origins", c.AllowedOrigins)
	d.Set("grant_types", c.GrantTypes)
	d.Set("web_origins", c.WebOrigins)
	d.Set("sso", c.SSO)
	d.Set("sso_disabled", c.SSODisabled)
	d.Set("cross_origin_auth", c.CrossOriginAuth)
	d.Set("cross_origin_loc", c.CrossOriginLocation)
	d.Set("custom_login_page_on", c.CustomLoginPageOn)
	d.Set("custom_login_page", c.CustomLoginPage)
	d.Set("custom_login_page_preview", c.CustomLoginPagePreview)
	d.Set("form_template", c.FormTemplate)
	d.Set("token_endpoint_auth_method", c.TokenEndpointAuthMethod)

	if jwtConfiguration := c.JWTConfiguration; jwtConfiguration != nil {
		d.Set("jwt_configuration", map[string]interface{}{
			"lifetime_in_seconds": jwtConfiguration.Algorithm,
			"secret_encoded":      jwtConfiguration.LifetimeInSeconds,
			"scopes":              jwtConfiguration.Scopes,
			"alg":                 jwtConfiguration.SecretEncoded,
		})
	}

	d.Set("encryption_key", c.EncryptionKey)
	d.Set("addons", c.Addons)
	d.Set("client_metadata", c.ClientMetadata)
	d.Set("mobile", c.Mobile)

	return nil
}

func updateClient(d *schema.ResourceData, m interface{}) error {
	c := buildClient(d)
	api := m.(*management.Management)
	err := api.Client.Update(d.Id(), c)
	if err != nil {
		return err
	}
	return readClient(d, m)
}

func deleteClient(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	return api.Client.Delete(d.Id())
}

func buildClient(d *schema.ResourceData) *management.Client {

	c := &management.Client{
		Name:                    String(d, "name"),
		Description:             String(d, "description"),
		AppType:                 String(d, "app_type"),
		LogoURI:                 String(d, "logo_uri"),
		IsFirstParty:            Bool(d, "is_first_party"),
		OIDCConformant:          Bool(d, "oidc_conformant"),
		Callbacks:               Slice(d, "callbacks"),
		AllowedLogoutURLs:       Slice(d, "allowed_logout_urls"),
		AllowedOrigins:          Slice(d, "allowed_origins"),
		GrantTypes:              Slice(d, "grant_types"),
		WebOrigins:              Slice(d, "web_origins"),
		SSO:                     Bool(d, "sso"),
		SSODisabled:             Bool(d, "sso_disabled"),
		CrossOriginAuth:         Bool(d, "cross_origin_auth"),
		CrossOriginLocation:     String(d, "cross_origin_loc"),
		CustomLoginPageOn:       Bool(d, "custom_login_page_on"),
		CustomLoginPage:         String(d, "custom_login_page"),
		CustomLoginPagePreview:  String(d, "custom_login_page_preview"),
		FormTemplate:            String(d, "form_template"),
		TokenEndpointAuthMethod: String(d, "token_endpoint_auth_method"),
	}

	if v, ok := d.GetOk("jwt_configuration"); ok {
		vL := v.([]interface{})
		for _, v := range vL {
			jwtConfiguration := v.(map[string]interface{})

			c.JWTConfiguration = &management.ClientJWTConfiguration{
				LifetimeInSeconds: auth0.Int(jwtConfiguration["lifetime_in_seconds"].(int)),
				Scopes:            jwtConfiguration["scopes"],
				Algorithm:         auth0.String(jwtConfiguration["alg"].(string)),
			}
		}
	}

	if v, ok := d.GetOk("encryption_key"); ok {
		c.EncryptionKey = make(map[string]string)

		for _, item := range v.([]interface{}) {
			for key, val := range item.(map[string]string) {
				c.EncryptionKey[key] = val
			}
		}
	}

	if v, ok := d.GetOk("addons"); ok {

		c.Addons = make(map[string]interface{})

		for _, item := range v.([]interface{}) {
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

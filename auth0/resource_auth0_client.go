package auth0

import (
	"strconv"

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
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"aws": {
							Type:     schema.TypeMap,
							Optional: true,
						},
						"azure_blob": {
							Type:     schema.TypeMap,
							Optional: true,
						},
						"azure_sb": {
							Type:     schema.TypeMap,
							Optional: true,
						},
						"rms": {
							Type:     schema.TypeMap,
							Optional: true,
						},
						"mscrm": {
							Type:     schema.TypeMap,
							Optional: true,
						},
						"slack": {
							Type:     schema.TypeMap,
							Optional: true,
						},
						"sentry": {
							Type:     schema.TypeMap,
							Optional: true,
						},
						"box": {
							Type:     schema.TypeMap,
							Optional: true,
						},
						"cloudbees": {
							Type:     schema.TypeMap,
							Optional: true,
						},
						"concur": {
							Type:     schema.TypeMap,
							Optional: true,
						},
						"dropbox": {
							Type:     schema.TypeMap,
							Optional: true,
						},
						"echosign": {
							Type:     schema.TypeMap,
							Optional: true,
						},
						"egnyte": {
							Type:     schema.TypeMap,
							Optional: true,
						},
						"firebase": {
							Type:     schema.TypeMap,
							Optional: true,
						},
						"newrelic": {
							Type:     schema.TypeMap,
							Optional: true,
						},
						"office365": {
							Type:     schema.TypeMap,
							Optional: true,
						},
						"salesforce": {
							Type:     schema.TypeMap,
							Optional: true,
						},
						"salesforce_api": {
							Type:     schema.TypeMap,
							Optional: true,
						},
						"salesforce_sandbox_api": {
							Type:     schema.TypeMap,
							Optional: true,
						},
						"samlp": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"audience": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"recipient": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"create_upn_claim": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  true,
									},
									"passthrough_claims_with_no_mapping": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  true,
									},
									"map_unknown_claims_as_is": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
									},
									"map_identities": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  true,
									},
									"signature_algorithm": {
										Type:     schema.TypeString,
										Optional: true,
										Default:  "rsa-sha1",
									},
									"digest_algorithm": {
										Type:     schema.TypeString,
										Optional: true,
										Default:  "sha1",
									},
									"destination": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"lifetime_in_seconds": {
										Type:     schema.TypeInt,
										Optional: true,
										Default:  3600,
									},
									"sign_response": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"typed_attributes": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  true,
									},
									"include_attribute_name_format": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  true,
									},
									"name_identifier_format": {
										Type:     schema.TypeString,
										Optional: true,
										Default:  "urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified",
									},
									"authn_context_class_ref": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"binding": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"mappings": {
										Type:     schema.TypeMap,
										Optional: true,
										Elem:     schema.TypeString,
									},
									"logout": {
										Type:     schema.TypeMap,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"callback": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"slo_enabled": {
													Type:     schema.TypeBool,
													Optional: true,
												},
											},
										},
									},
									"name_identifier_probes": {
										Type:     schema.TypeList,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
									},
								},
							},
						},
						"layer": {
							Type:     schema.TypeMap,
							Optional: true,
						},
						"sap_api": {
							Type:     schema.TypeMap,
							Optional: true,
						},
						"sharepoint": {
							Type:     schema.TypeMap,
							Optional: true,
						},
						"springcm": {
							Type:     schema.TypeMap,
							Optional: true,
						},
						"wams": {
							Type:     schema.TypeMap,
							Optional: true,
						},
						"wsfed": {
							Type:     schema.TypeMap,
							Optional: true,
						},
						"zendesk": {
							Type:     schema.TypeMap,
							Optional: true,
						},
						"zoom": {
							Type:     schema.TypeMap,
							Optional: true,
						},
					},
				},
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

	List(d, "jwt_configuration").First(func(v interface{}) {

		m := v.(map[string]interface{})

		c.JWTConfiguration = &management.ClientJWTConfiguration{
			LifetimeInSeconds: Int(MapData(m), "lifetime_in_seconds"),
			Algorithm:         String(MapData(m), "alg"),
			Scopes:            Map(MapData(m), "scopes"),
		}
	})

	List(d, "encryption_key").First(func(v interface{}) {
		c.EncryptionKey = v.(map[string]string)
	})

	List(d, "addons").First(func(v interface{}) {

		c.Addons = make(map[string]interface{})

		for addonKey, addonValue := range v.(map[string]interface{}) {

			switch addonKey {

			case "samlp":
				for _, v := range addonValue.([]interface{}) {

					addon := v.(map[string]interface{})

					if len(addon) > 0 {
						c.Addons[addonKey] = buildClientAddon(addon)
					}
				}

			default:
				addon := addonValue.(map[string]interface{})
				if len(addon) > 0 {
					c.Addons[addonKey] = buildClientAddon(addon)
				}
			}
		}
	})

	List(d, "client_metadata").First(func(v interface{}) {
		c.ClientMetadata = v.(map[string]string)
	})

	List(d, "mobile").First(func(v interface{}) {

		c.Mobile = make(map[string]interface{})

		for mobileKey, mobileValues := range v.(map[string]interface{}) {

			for _, mobile := range mobileValues.([]interface{}) {
				c.Mobile[mobileKey] = mobile
			}
		}
	})

	return c
}

func buildClientAddon(d map[string]interface{}) map[string]interface{} {

	addon := make(map[string]interface{})

	for key, value := range d {

		switch v := value.(type) {

		case string:
			if i, err := strconv.ParseInt(v, 10, 64); err == nil {
				addon[key] = i
			} else if f, err := strconv.ParseFloat(v, 64); err == nil {
				addon[key] = f
			} else if b, err := strconv.ParseBool(v); err == nil {
				addon[key] = b
			} else {
				addon[key] = v
			}

		case map[string]interface{}:
			addon[key] = buildClientAddon(v)

		case []interface{}:
			addon[key] = v

		default:
			addon[key] = v
		}
	}
	return addon
}

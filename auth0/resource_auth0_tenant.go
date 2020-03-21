package auth0

import (
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"gopkg.in/auth0.v4/management"

	v "github.com/terraform-providers/terraform-provider-auth0/auth0/internal/validation"
)

func newTenant() *schema.Resource {
	return &schema.Resource{

		Create: createTenant,
		Read:   readTenant,
		Update: updateTenant,
		Delete: deleteTenant,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"change_password": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"html": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"guardian_mfa_page": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"html": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"default_audience": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"default_directory": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"error_page": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"html": {
							Type:     schema.TypeString,
							Required: true,
						},
						"show_log_link": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"url": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"picture_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"support_email": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"support_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"allowed_logout_urls": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"session_lifetime": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sandbox_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"idle_session_lifetime": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"enabled_locales": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"flags": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"change_pwd_flow_v1": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"enable_client_connections": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"enable_apis_section": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"enable_pipeline2": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"enable_dynamic_client_registration": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"enable_custom_domain_in_emails": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"universal_login": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"enable_legacy_logs_search_v2": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"disable_clickjack_protection_headers": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"enable_public_signup_user_exists_error": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
			"universal_login": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"colors": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"primary": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"page_background": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			"default_redirection_uri": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.All(
					v.IsURLWithNoFragment,
					validation.IsURLWithScheme([]string{"https"}),
				),
			},
		},
	}
}

func createTenant(d *schema.ResourceData, m interface{}) error {
	d.SetId(resource.UniqueId())
	return updateTenant(d, m)
}

func readTenant(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	t, err := api.Tenant.Read()
	if err != nil {
		if mErr, ok := err.(management.Error); ok {
			if mErr.Status() == http.StatusNotFound {
				d.SetId("")
				return nil
			}
		}
		return err
	}

	d.Set("change_password", flattenTenantChangePassword(t.ChangePassword))
	d.Set("guardian_mfa_page", flattenTenantGuardianMFAPage(t.GuardianMFAPage))

	d.Set("default_audience", t.DefaultAudience)
	d.Set("default_directory", t.DefaultDirectory)

	d.Set("friendly_name", t.FriendlyName)
	d.Set("picture_url", t.PictureURL)
	d.Set("support_email", t.SupportEmail)
	d.Set("support_url", t.SupportURL)
	d.Set("allowed_logout_urls", t.AllowedLogoutURLs)
	d.Set("session_lifetime", t.SessionLifetime)
	d.Set("sandbox_version", t.SandboxVersion)
	d.Set("idle_session_lifetime", t.IdleSessionLifetime)

	// TODO: figure out why it is necessary to wrap what is aleady an
	// []interface{} with another []interface{}.
	d.Set("enabled_locales", []interface{}{t.EnabledLocales})

	d.Set("error_page", flattenTenantErrorPage(t.ErrorPage))
	d.Set("flags", flattenTenantFlags(t.Flags))
	d.Set("universal_login", flattenTenantUniversalLogin(t.UniversalLogin))

	return nil
}

func updateTenant(d *schema.ResourceData, m interface{}) error {
	t := buildTenant(d)
	api := m.(*management.Management)
	err := api.Tenant.Update(t)
	if err != nil {
		return err
	}
	return readTenant(d, m)
}

func deleteTenant(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}

func buildTenant(d *schema.ResourceData) *management.Tenant {
	t := &management.Tenant{
		DefaultAudience:     String(d, "default_audience"),
		DefaultDirectory:    String(d, "default_directory"),
		FriendlyName:        String(d, "friendly_name"),
		PictureURL:          String(d, "picture_url"),
		SupportEmail:        String(d, "support_email"),
		SupportURL:          String(d, "support_url"),
		AllowedLogoutURLs:   Slice(d, "allowed_logout_urls"),
		SessionLifetime:     Int(d, "session_lifetime"),
		SandboxVersion:      String(d, "sandbox_version"),
		IdleSessionLifetime: Int(d, "idle_session_lifetime"),
		EnabledLocales:      Set(d, "enabled_locales").List(),
		ChangePassword:      expandTenantChangePassword(d),
		GuardianMFAPage:     expandTenantGuardianMFAPage(d),
		ErrorPage:           expandTenantErrorPage(d),
		Flags:               expandTenantFlags(d),
		UniversalLogin:      expandTenantUniversalLogin(d),
	}

	return t
}

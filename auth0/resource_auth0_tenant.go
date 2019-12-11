package auth0

import (
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"gopkg.in/auth0.v2/management"
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
		return err
	}

	if changePassword := t.ChangePassword; changePassword != nil {
		d.Set("change_password", []map[string]interface{}{
			{
				"enabled": changePassword.Enabled,
				"html":    changePassword.HTML,
			},
		})
	}

	if guardianMFAPage := t.GuardianMFAPage; guardianMFAPage != nil {
		d.Set("guardian_mfa_page", []map[string]interface{}{
			{
				"enabled": guardianMFAPage.Enabled,
				"html":    guardianMFAPage.HTML,
			},
		})
	}

	d.Set("default_audience", t.DefaultAudience)
	d.Set("default_directory", t.DefaultDirectory)

	if errorPage := t.ErrorPage; errorPage != nil {
		d.Set("error_page", []map[string]interface{}{
			{
				"html":          errorPage.HTML,
				"show_log_link": errorPage.ShowLogLink,
				"url":           errorPage.URL,
			},
		})
	}

	d.Set("friendly_name", t.FriendlyName)
	d.Set("picture_url", t.PictureURL)
	d.Set("support_email", t.SupportEmail)
	d.Set("support_url", t.SupportURL)
	d.Set("allowed_logout_urls", t.AllowedLogoutURLs)
	d.Set("session_lifetime", t.SessionLifetime)
	d.Set("sandbox_version", t.SandboxVersion)
	d.Set("idle_session_lifetime", t.IdleSessionLifetime)

	if flags := t.Flags; flags != nil {
		d.Set("flags", []map[string]interface{}{
			{
				"change_pwd_flow_v1":                     flags.ChangePasswordFlowV1,
				"enable_client_connections":              flags.EnableClientConnections,
				"enable_apis_section":                    flags.EnableAPIsSection,
				"enable_pipeline2":                       flags.EnablePipeline2,
				"enable_dynamic_client_registration":     flags.EnableDynamicClientRegistration,
				"enable_custom_domain_in_emails":         flags.EnableCustomDomainInEmails,
				"universal_login":                        flags.UniversalLogin,
				"enable_legacy_logs_search_v2":           flags.EnableLegacyLogsSearchV2,
				"disable_clickjack_protection_headers":   flags.DisableClickjackProtectionHeaders,
				"enable_public_signup_user_exists_error": flags.EnablePublicSignupUserExistsError,
			},
		})
	}

	if universalLogin := t.UniversalLogin; universalLogin != nil {
		if colors := universalLogin.Colors; colors != nil {
			d.Set("universal_login", []map[string]interface{}{
				{
					"colors": []map[string]interface{}{
						{
							"primary":         universalLogin.Colors.Primary,
							"page_background": universalLogin.Colors.PageBackground,
						},
					},
				},
			})
		}
	}

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
	}

	List(d, "change_password").First(func(v interface{}) {
		m := v.(map[string]interface{})

		t.ChangePassword = &management.TenantChangePassword{
			Enabled: Bool(MapData(m), "enabled"),
			HTML:    String(MapData(m), "html"),
		}
	})

	List(d, "guardian_mfa_page").First(func(v interface{}) {
		m := v.(map[string]interface{})

		t.GuardianMFAPage = &management.TenantGuardianMFAPage{
			Enabled: Bool(MapData(m), "enabled"),
			HTML:    String(MapData(m), "html"),
		}
	})

	List(d, "error_page").First(func(v interface{}) {
		m := v.(map[string]interface{})

		t.ErrorPage = &management.TenantErrorPage{
			HTML:        String(MapData(m), "html"),
			ShowLogLink: Bool(MapData(m), "show_log_link"),
			URL:         String(MapData(m), "url"),
		}
	})

	List(d, "flags").First(func(v interface{}) {
		m := v.(map[string]interface{})

		t.Flags = &management.TenantFlags{
			ChangePasswordFlowV1:              Bool(MapData(m), "change_pwd_flow_v1"),
			EnableClientConnections:           Bool(MapData(m), "enable_client_connections"),
			EnableAPIsSection:                 Bool(MapData(m), "enable_apis_section"),
			EnablePipeline2:                   Bool(MapData(m), "enable_pipeline2"),
			EnableDynamicClientRegistration:   Bool(MapData(m), "enable_dynamic_client_registration"),
			EnableCustomDomainInEmails:        Bool(MapData(m), "enable_custom_domain_in_emails"),
			UniversalLogin:                    Bool(MapData(m), "universal_login"),
			EnableLegacyLogsSearchV2:          Bool(MapData(m), "enable_legacy_logs_search_v2"),
			DisableClickjackProtectionHeaders: Bool(MapData(m), "disable_clickjack_protection_headers"),
			EnablePublicSignupUserExistsError: Bool(MapData(m), "enable_public_signup_user_exists_error"),
		}
	})

	List(d, "universal_login").First(func(v interface{}) {
		m := v.(map[string]interface{})

		t.UniversalLogin = &management.TenantUniversalLogin{}

		List(MapData(m), "colors").First(func(v interface{}) {
			m := v.(map[string]interface{})
			t.UniversalLogin.Colors = &management.TenantUniversalLoginColors{
				Primary:        String(MapData(m), "primary"),
				PageBackground: String(MapData(m), "page_background"),
			}
		})
	})

	return t
}

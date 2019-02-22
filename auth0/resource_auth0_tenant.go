package auth0

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/yieldr/go-auth0/management"
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
			"flags": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"change_pwd_flow_v1": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"enable_client_connections": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"enable_apis_section": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"enable_pipeline2": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"enable_dynamic_client_registration": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"enable_custom_domain_in_emails": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
		},
	}
}

func createTenant(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	d.SetId(api.Domain)
	return readTenant(d, m)
}

func readTenant(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	t, err := api.Tenant.Read()
	if err != nil {
		return err
	}

	d.Set("change_password", t.ChangePassword)
	d.Set("guardian_mfa_page", t.GuardianMFAPage)
	d.Set("default_audience", t.DefaultAudience)
	d.Set("default_directory", t.DefaultDirectory)
	d.Set("error_page", t.ErrorPage)
	d.Set("flags", t.Flags)
	d.Set("friendly_name", t.FriendlyName)
	d.Set("picture_url", t.PictureURL)
	d.Set("support_email", t.SupportEmail)
	d.Set("support_url", t.SupportURL)
	d.Set("allowed_logout_urls", t.AllowedLogoutURLs)
	d.Set("session_lifetime", t.SessionLifetime)
	d.Set("sandbox_version", t.SandboxVersion)

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
	return nil
}

func buildTenant(d *schema.ResourceData) *management.Tenant {
	t := &management.Tenant{
		DefaultAudience:   String(d, "default_audience"),
		DefaultDirectory:  String(d, "default_directory"),
		FriendlyName:      String(d, "friendly_name"),
		PictureURL:        String(d, "picture_url"),
		SupportEmail:      String(d, "support_email"),
		SupportURL:        String(d, "support_url"),
		AllowedLogoutURLs: Slice(d, "allowed_logout_urls"),
		SessionLifetime:   Int(d, "session_lifetime"),
		SandboxVersion:    String(d, "sandbox_version"),
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

	List(d, "flags").First(func(v interface{}) {
		m := v.(map[string]interface{})
		t.Flags = &management.TenantFlags{
			ChangePasswordFlowV1:            Bool(MapData(m), "change_pwd_flow_v1"),
			EnableClientConnections:         Bool(MapData(m), "enable_client_connections"),
			EnableAPIsSection:               Bool(MapData(m), "enable_apis_section"),
			EnablePipeline2:                 Bool(MapData(m), "enable_pipeline2"),
			EnableDynamicClientRegistration: Bool(MapData(m), "enable_dynamic_client_registration"),
			EnableCustomDomainInEmails:      Bool(MapData(m), "enable_custom_domain_in_emails"),
		}
	})

	return t
}

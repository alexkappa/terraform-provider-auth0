package auth0

import (
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"gopkg.in/auth0.v1"
	"gopkg.in/auth0.v1/management"
)

func newConnection() *schema.Resource {
	return &schema.Resource{

		Create: createConnection,
		Read:   readConnection,
		Update: updateConnection,
		Delete: deleteConnection,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"is_domain_connection": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"strategy": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"ad", "adfs", "amazon", "dropbox", "bitbucket", "aol",
					"auth0-adldap", "auth0-oidc", "auth0", "baidu", "bitly",
					"box", "custom", "daccount", "dwolla", "email",
					"evernote-sandbox", "evernote", "exact", "facebook",
					"fitbit", "flickr", "github", "google-apps",
					"google-oauth2", "guardian", "instagram", "ip", "linkedin",
					"miicard", "oauth1", "oauth2", "office365", "paypal",
					"paypal-sandbox", "pingfederate", "planningcenter",
					"renren", "salesforce-community", "salesforce-sandbox",
					"salesforce", "samlp", "sharepoint", "shopify", "sms",
					"soundcloud", "thecity-sandbox", "thecity",
					"thirtysevensignals", "twitter", "untappd", "vkontakte",
					"waad", "weibo", "windowslive", "wordpress", "yahoo",
					"yammer", "yandex",
				}, true),
				ForceNew: true,
			},
			"options": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"validation": {
							Type:     schema.TypeMap,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
						"password_policy": {
							Type:     schema.TypeString,
							Optional: true,
							ValidateFunc: validation.StringInSlice([]string{
								"none", "low", "fair", "good", "excellent",
							}, false),
						},
						"password_history": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enable": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"size": {
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"password_no_personal_info": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enable": {
										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},
						"password_dictionary": {
							Type:     schema.TypeMap,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
						"api_enable_users": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"basic_profile": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"ext_admin": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"ext_is_suspended": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"ext_agreed_terms": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"ext_groups": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"ext_nested_groups": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"ext_assigned_plans": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"ext_profile": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"enabled_database_customization": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"brute_force_protection": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"import_mode": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"disable_signup": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"requires_username": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"custom_scripts": {
							Type:     schema.TypeMap,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
						"configuration": {
							Type:      schema.TypeMap,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Sensitive: true,
							Optional:  true,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								return strings.HasPrefix(old, "2.0$") || new == old
							},
						},
						// waad options
						"app_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"app_domain": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"client_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"client_secret": {
							Type:      schema.TypeString,
							Optional:  true,
							Sensitive: true,
						},
						"domain_aliases": {
							Type:     schema.TypeList,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
						"max_groups_to_retrieve": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"tenant_domain": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"use_wsfed": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"waad_protocol": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"waad_common_endpoint": {
							Type:     schema.TypeBool,
							Optional: true,
						},

						// Twilio/sms options
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"twilio_sid": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"twilio_token": {
							Type:        schema.TypeString,
							Optional:    true,
							Sensitive:   true,
							DefaultFunc: schema.EnvDefaultFunc("TWILIO_TOKEN", nil),
						},
						"from": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"syntax": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"template": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"totp": {
							Type:     schema.TypeMap,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"time_step": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"length": {
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"messaging_service_sid": {
							Type:     schema.TypeString,
							Optional: true,
						},

						// adfs options
						"adfs_server": {
							Type:     schema.TypeString,
							Optional: true,
						},

						// salesforce options
						"community_base_url": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"enabled_clients": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"realms": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createConnection(d *schema.ResourceData, m interface{}) error {
	c := buildConnection(d)
	api := m.(*management.Management)
	if err := api.Connection.Create(c); err != nil {
		return err
	}
	d.SetId(auth0.StringValue(c.ID))
	return readConnection(d, m)
}

func readConnection(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	c, err := api.Connection.Read(d.Id())
	if err != nil {
		return err
	}
	d.SetId(auth0.StringValue(c.ID))
	d.Set("name", c.Name)
	d.Set("is_domain_connection", c.IsDomainConnection)
	d.Set("strategy", c.Strategy)
	d.Set("options", []map[string]interface{}{
		{
			"validation":                     c.Options.Validation,
			"password_policy":                auth0.StringValue(c.Options.PasswordPolicy),
			"password_history":               c.Options.PasswordHistory,
			"password_no_personal_info":      c.Options.PasswordNoPersonalInfo,
			"password_dictionary":            c.Options.PasswordDictionary,
			"api_enable_users":               auth0.BoolValue(c.Options.APIEnableUsers),
			"basic_profile":                  auth0.BoolValue(c.Options.BasicProfile),
			"ext_admin":                      auth0.BoolValue(c.Options.ExtAdmin),
			"ext_is_suspended":               auth0.BoolValue(c.Options.ExtIsSuspended),
			"ext_agreed_terms":               auth0.BoolValue(c.Options.ExtAgreedTerms),
			"ext_groups":                     auth0.BoolValue(c.Options.ExtGroups),
			"ext_nested_groups":              auth0.BoolValue(c.Options.ExtNestedGroups),
			"ext_assigned_plans":             auth0.BoolValue(c.Options.ExtAssignedPlans),
			"ext_profile":                    auth0.BoolValue(c.Options.ExtProfile),
			"enabled_database_customization": auth0.BoolValue(c.Options.EnabledDatabaseCustomization),
			"brute_force_protection":         auth0.BoolValue(c.Options.BruteForceProtection),
			"import_mode":                    auth0.BoolValue(c.Options.ImportMode),
			"disable_signup":                 auth0.BoolValue(c.Options.DisableSignup),
			"requires_username":              auth0.BoolValue(c.Options.RequiresUsername),
			"custom_scripts":                 c.Options.CustomScripts,
			"configuration":                  c.Options.Configuration,

			// waad
			"app_id":                 auth0.StringValue(c.Options.AppID),
			"app_domain":             auth0.StringValue(c.Options.AppDomain),
			"client_id":              auth0.StringValue(c.Options.ClientID),
			"client_secret":          auth0.StringValue(c.Options.ClientSecret),
			"domain_aliases":         c.Options.DomainAliases,
			"max_groups_to_retrieve": auth0.StringValue(c.Options.MaxGroupsToRetrieve),
			"tenant_domain":          auth0.StringValue(c.Options.TenantDomain),
			"use_wsfed":              auth0.BoolValue(c.Options.UseWsfed),
			"waad_protocol":          auth0.StringValue(c.Options.WaadProtocol),
			"waad_common_endpoint":   auth0.BoolValue(c.Options.WaadCommonEndpoint),

			// twilio/sms
			"name":                  auth0.StringValue(c.Options.Name),
			"twilio_sid":            auth0.StringValue(c.Options.TwilioSid),
			"twilio_token":          auth0.StringValue(c.Options.TwilioToken),
			"from":                  auth0.StringValue(c.Options.From),
			"syntax":                auth0.StringValue(c.Options.Syntax),
			"template":              auth0.StringValue(c.Options.Template),
			"messaging_service_sid": auth0.StringValue(c.Options.MessagingServiceSid),
			"totp":                  c.Options.Totp,

			// adfs
			"adfs_server": auth0.StringValue(c.Options.AdfsServer),

			// salesforce
			"community_base_url": auth0.StringValue(c.Options.CommunityBaseURL),
		},
	})

	d.Set("enabled_clients", c.EnabledClients)
	d.Set("realms", c.Realms)
	return nil
}

func updateConnection(d *schema.ResourceData, m interface{}) error {
	c := buildConnection(d)
	c.Strategy = nil
	c.Name = nil
	api := m.(*management.Management)
	err := api.Connection.Update(d.Id(), c)
	if err != nil {
		return err
	}
	return readConnection(d, m)
}

func deleteConnection(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	return api.Connection.Delete(d.Id())
}

func buildConnection(d *schema.ResourceData) *management.Connection {

	c := &management.Connection{
		Name:               String(d, "name"),
		IsDomainConnection: Bool(d, "is_domain_connection"),
		Strategy:           String(d, "strategy"),
		EnabledClients:     Set(d, "enabled_clients").Slice(),
		Realms:             Slice(d, "realms"),
	}

	List(d, "options").First(func(v interface{}) {

		m := v.(map[string]interface{})

		c.Options = &management.ConnectionOptions{
			Validation:                   Map(MapData(m), "validation"),
			PasswordPolicy:               String(MapData(m), "password_policy"),
			PasswordDictionary:           Map(MapData(m), "password_dictionary"),
			APIEnableUsers:               Bool(MapData(m), "api_enable_users"),
			BasicProfile:                 Bool(MapData(m), "basic_profile"),
			ExtAdmin:                     Bool(MapData(m), "ext_admin"),
			ExtIsSuspended:               Bool(MapData(m), "ext_is_suspended"),
			ExtAgreedTerms:               Bool(MapData(m), "ext_agreed_terms"),
			ExtGroups:                    Bool(MapData(m), "ext_groups"),
			ExtNestedGroups:              Bool(MapData(m), "ext_nested_groups"),
			ExtAssignedPlans:             Bool(MapData(m), "ext_assigned_plans"),
			ExtProfile:                   Bool(MapData(m), "ext_profile"),
			EnabledDatabaseCustomization: Bool(MapData(m), "enabled_database_customization"),
			BruteForceProtection:         Bool(MapData(m), "brute_force_protection"),
			ImportMode:                   Bool(MapData(m), "import_mode"),
			DisableSignup:                Bool(MapData(m), "disable_signup"),
			RequiresUsername:             Bool(MapData(m), "requires_username"),
			CustomScripts:                Map(MapData(m), "custom_scripts"),
			Configuration:                Map(MapData(m), "configuration"),

			// Waad
			AppID:               String(MapData(m), "app_id"),
			AppDomain:           String(MapData(m), "app_domain"),
			ClientID:            String(MapData(m), "client_id"),
			ClientSecret:        String(MapData(m), "client_secret"),
			DomainAliases:       Slice(MapData(m), "domain_aliases"),
			MaxGroupsToRetrieve: String(MapData(m), "max_groups_to_retrieve"),
			TenantDomain:        String(MapData(m), "tenant_domain"),
			UseWsfed:            Bool(MapData(m), "use_wsfed"),
			WaadProtocol:        String(MapData(m), "waad_protocol"),
			WaadCommonEndpoint:  Bool(MapData(m), "waad_common_endpoint"),

			// Twilio
			Name:                String(MapData(m), "name"),
			TwilioSid:           String(MapData(m), "twilio_sid"),
			TwilioToken:         String(MapData(m), "twilio_token"),
			From:                String(MapData(m), "from"),
			Syntax:              String(MapData(m), "syntax"),
			Template:            String(MapData(m), "template"),
			MessagingServiceSid: String(MapData(m), "messaging_service_sid"),
			Totp: &management.ConnectionOptionsTotp{
				TimeStep: Int(MapData(m), "time_step"),
				Length:   Int(MapData(m), "length"),
			},

			// adfs
			AdfsServer: String(MapData(m), "adfs_server"),
		}

		List(MapData(m), "password_history").First(func(v interface{}) {

			m := v.(map[string]interface{})

			c.Options.PasswordHistory = make(map[string]interface{})
			c.Options.PasswordHistory["enable"] = Bool(MapData(m), "enable")
			c.Options.PasswordHistory["size"] = Int(MapData(m), "size")
		})

		List(MapData(m), "password_no_personal_info").First(func(v interface{}) {

			m := v.(map[string]interface{})

			c.Options.PasswordNoPersonalInfo = make(map[string]interface{})
			c.Options.PasswordNoPersonalInfo["enable"] = Bool(MapData(m), "enable")
		})
	})

	return c
}

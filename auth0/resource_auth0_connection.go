package auth0

import (
	"sort"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	auth0 "github.com/yieldr/go-auth0"
	"github.com/yieldr/go-auth0/management"
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
							Type:     schema.TypeMap,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
						"password_no_personal_info": {
							Type:     schema.TypeMap,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
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
					},
				},
			},
			"enabled_clients": {
				Type:     schema.TypeList,
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
	d.Set("strategy", c.Strategy)
	d.Set("options", []map[string]interface{}{
		{
			"validation":                     c.Options.Validation,
			"password_policy":                c.Options.PasswordPolicy,
			"password_history":               c.Options.PasswordHistory,
			"password_no_personal_info":      c.Options.PasswordNoPersonalInfo,
			"password_dictionary":            c.Options.PasswordDictionary,
			"api_enable_users":               c.Options.APIEnableUsers,
			"basic_profile":                  c.Options.BasicProfile,
			"ext_admin":                      c.Options.ExtAdmin,
			"ext_is_suspended":               c.Options.ExtIsSuspended,
			"ext_agreed_terms":               c.Options.ExtAgreedTerms,
			"ext_groups":                     c.Options.ExtGroups,
			"ext_assigned_plans":             c.Options.ExtAssignedPlans,
			"ext_profile":                    c.Options.ExtProfile,
			"enabled_database_customization": c.Options.EnabledDatabaseCustomization,
			"brute_force_protection":         c.Options.BruteForceProtection,
			"import_mode":                    c.Options.ImportMode,
			"disable_signup":                 c.Options.DisableSignup,
			"requires_username":              c.Options.RequiresUsername,
		},
	})

	sort.Slice(c.EnabledClients, func(i, j int) bool {
		return c.EnabledClients[i].(string) > c.EnabledClients[j].(string)
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
		Name:           String(d, "name"),
		Strategy:       String(d, "strategy"),
		EnabledClients: Slice(d, "enabled_clients"),
		Realms:         Slice(d, "realms"),
	}

	if v, ok := d.GetOk("options"); ok {

		vL := v.([]interface{})
		for _, v := range vL {

			if options, ok := v.(map[string]interface{}); ok {
				c.Options = &management.ConnectionOptions{
					Validation:                   options["validation"].(map[string]interface{}),
					PasswordPolicy:               auth0.String(options["password_policy"].(string)),
					PasswordHistory:              options["password_history"].(map[string]interface{}),
					PasswordNoPersonalInfo:       options["password_no_personal_info"].(map[string]interface{}),
					PasswordDictionary:           options["password_dictionary"].(map[string]interface{}),
					APIEnableUsers:               auth0.Bool(options["api_enable_users"].(bool)),
					BasicProfile:                 auth0.Bool(options["basic_profile"].(bool)),
					ExtAdmin:                     auth0.Bool(options["ext_admin"].(bool)),
					ExtIsSuspended:               auth0.Bool(options["ext_is_suspended"].(bool)),
					ExtAgreedTerms:               auth0.Bool(options["ext_agreed_terms"].(bool)),
					ExtGroups:                    auth0.Bool(options["ext_groups"].(bool)),
					ExtAssignedPlans:             auth0.Bool(options["ext_assigned_plans"].(bool)),
					ExtProfile:                   auth0.Bool(options["ext_profile"].(bool)),
					EnabledDatabaseCustomization: auth0.Bool(options["enabled_database_customization"].(bool)),
					BruteForceProtection:         auth0.Bool(options["brute_force_protection"].(bool)),
					ImportMode:                   auth0.Bool(options["import_mode"].(bool)),
					DisableSignup:                auth0.Bool(options["disable_signup"].(bool)),
					RequiresUsername:             auth0.Bool(options["requires_username"].(bool)),
				}
			}
		}
	}

	return c
}

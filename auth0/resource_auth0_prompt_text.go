package auth0

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"gopkg.in/auth0.v5/management"
)

func newPromptText() *schema.Resource {
	return &schema.Resource{

		Create: createPromptText,
		Read:   readPromptText,
		Update: updatePromptText,
		Delete: deletePromptText,

		Schema: map[string]*schema.Schema{
			"language": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"ar", "bg", "bs", "cs", "da", "de", "el", "en", "es", "et", "fi", "fr",
					"fr-CA", "fr-FR", "he", "hi", "hr", "hu", "id", "is", "it", "ja", "ko",
					"lt", "lv", "nb", "nl", "pl", "pt", "pt-BR", "pt-PT", "ro", "ru", "sk",
					"sl", "sr", "sv", "th", "tr", "uk", "vi", "zh-CN",
				}, false),
			},
			"prompt_name": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"login", "login_id", "login_password", "login_email_verification", "signup",
					"signup_id", "signup_password", "reset_password", "consent", "mfa_push",
					"mfa_otp", "mfa_voice", "mfa_phone", "mfa_webauthn", "mfa_sms", "mfa_email",
					"mfa_recovery_code", "mfa", "status", "device_flow", "email_verification",
					"email_otp_challenge", "organizations", "invitation", "common",
				}, false),
			},
			"prompt_content": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Only the reset-password prompt screens and fields are included right now
						"reset_password_request": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"invalid_email_format": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"auth0_users_expired_ticket": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"custom_script_error_code": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"auth0_users_used_ticket": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"auth0_users_validation": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"reset_password_error": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"too_many_email": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"too_many_requests": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"no_email": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"no_username": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"page_title": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"title": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"back_to_login_link_text": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"button_text": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"description_email": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"description_username": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"placeholder_email": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"placeholder_username": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"logo_alt_text": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"reset_password_email": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"page_title": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"title": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"email_description": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"resend_link_text": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"username_description": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"reset_password": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"page_title": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"title": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"description": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"button_text": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"password_placeholder": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"re_enterpassword_placeholder": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"password_security_text": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"logo_alt_text": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"auth0_users_expired_ticket": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"custom_script_error_code": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"auth0_users_used_ticket": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"auth0_users_validation": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"no_re_enter_password": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"reset_password_success": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"page_title": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"event_title": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"description": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"button_text": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"reset_password_error": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"page_title": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"back_to_login_link_text": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"description_expired": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"description_generic": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"description_used": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"event_title_expired": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"event_title_generic": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"event_title_used": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"auth0_users_expired_ticket": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"custom_script_error_code": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"auth0_users_used_ticket": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"auth0_users_validation": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"reset_password_error": {
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

func createPromptText(d *schema.ResourceData, m interface{}) error {
	d.SetId(resource.UniqueId())
	return updatePromptText(d, m)
}

func readPromptText(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	pn, l := getPromptNameAndLanguage(d)
	pc, err := api.Prompt.CustomText(pn, l)
	if err != nil {
		if mErr, ok := err.(management.Error); ok {
			if mErr.Status() == http.StatusNotFound {
				d.SetId("")
				return nil
			}
		}
		return err
	}

	// Does not work (diffs aren't detected on the live resource itself)
	// Seeking guidance to get started implementing this properly
	d.Set("prompt_content", pc)
	return nil
}

func updatePromptText(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)

	pn, l := getPromptNameAndLanguage(d)
	pc := buildPromptContent(d)

	err := api.Prompt.SetCustomText(pn, l, pc)
	if err != nil {
		return err
	}

	return readPromptText(d, m)
}

func deletePromptText(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	pn, l := getPromptNameAndLanguage(d)

	// Set prompt to empty map
	err := api.Prompt.SetCustomText(pn, l, make(map[string]interface{}))
	d.SetId("")
	return err
}

func getPromptNameAndLanguage(d *schema.ResourceData) (string, string) {
	pn := convertSnakeToKebabCase(d.Get("prompt_name").(string))
	l := d.Get("language").(string)
	return pn, l
}

// Build map for submission to auth0 using the casing auth0 expects
func buildPromptContent(d *schema.ResourceData) map[string]interface{} {
	pt := d.Get("prompt_content").([]interface{})[0].(map[string]interface{})

	auth0_fmt_prompt := make(map[string]interface{})

	for screen_name, content := range pt {
		auth0_fmt_screen := make(map[string]interface{})
		for _, fields := range content.([]interface{}) {
			for field, text := range fields.(map[string]interface{}) {
				// exclude fields not supplied in the tf module resource
				if text != "" {
					if isCamelCaseField(field) {
						auth0_fmt_screen[convertSnakeToCamelCase(field)] = text
					} else if isKebabCaseField(field) {
						auth0_fmt_screen[convertSnakeToKebabCase(field)] = text
					}
				}
			}
			// All screens use kebab-case
			auth0_fmt_prompt[convertSnakeToKebabCase(screen_name)] = auth0_fmt_screen
		}
	}

	return auth0_fmt_prompt
}

func convertSnakeToKebabCase(sc string) string {
	return strings.ReplaceAll(sc, "_", "-")
}

func convertSnakeToCamelCase(sc string) string {
	var snakePattern = regexp.MustCompile("_[a-z]")
	return snakePattern.ReplaceAllStringFunc(sc, func(s string) string {
		return strings.ToTitle(strings.Replace(s, "_", "", 1))
	})
}

// Return true if Auth0 expects field in kebab-case format
func isKebabCaseField(s string) bool {
	switch s {
	case
		"no_email",
		"custom_script_error_code",
		"auth0_users_validation",
		"auth0_users_used_ticket",
		"too_many_requests",
		"invalid_email_format",
		"reset_password_error",
		"too_many_email",
		"auth0_users_expired_ticket",
		"no_re_enter_password",
		"no_username":
		return true
	}
	return false
}

// Return true if Auth0 expects field in camelCase format
func isCamelCaseField(s string) bool {
	switch s {
	case
		"page_title",
		"event_title_used",
		"description_username",
		"resend_link_text",
		"password_security_text",
		"description_used",
		"description",
		"description_email",
		"description_generic",
		"password_placeholder",
		"event_title_generic",
		"back_to_login_link_text",
		"email_description",
		"event_title",
		"placeholder_email",
		"button_text",
		"re_enterpassword_placeholder",
		"placeholder_username",
		"logo_alt_text",
		"username_description",
		"title",
		"description_expired",
		"event_title_expired":
		return true
	}
	return false
}

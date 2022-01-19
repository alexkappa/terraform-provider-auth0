package auth0

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/structure"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"gopkg.in/auth0.v5/management"
)

var (
	availablePrompts = []string{
		"login", "login-id", "login-password", "login-email-verification", "signup", "signup-id", "signup-password",
		"reset-password", "consent", "mfa-push", "mfa-otp", "mfa-voice", "mfa-phone", "mfa-webauthn", "mfa-sms",
		"mfa-email", "mfa-recovery-code", "mfa", "status", "device-flow", "email-verification", "email-otp-challenge",
		"organizations", "invitation", "common",
	}
	availableLanguages = []string{
		"ar", "bg", "bs", "cs", "da", "de", "el", "en", "es", "et", "fi", "fr", "fr-CA", "fr-FR", "he", "hi", "hr",
		"hu", "id", "is", "it", "ja", "ko", "lt", "lv", "nb", "nl", "pl", "pt", "pt-BR", "pt-PT", "ro", "ru", "sk",
		"sl", "sr", "sv", "th", "tr", "uk", "vi", "zh-CN", "zh-TW",
	}
	errEmptyPromptCustomTextID         = fmt.Errorf("ID cannot be empty")
	errInvalidPromptCustomTextIDFormat = fmt.Errorf("ID must be formated as prompt:language")
)

func newPromptCustomText() *schema.Resource {
	return &schema.Resource{
		Create: createPromptCustomText,
		Read:   readPromptCustomText,
		Update: updatePromptCustomText,
		Delete: deletePromptCustomText,
		Importer: &schema.ResourceImporter{
			State: importPromptCustomText,
		},
		Schema: map[string]*schema.Schema{
			"prompt": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(availablePrompts, false),
			},
			"language": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(availableLanguages, false),
			},
			"body": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateFunc:     validation.StringIsJSON,
				DiffSuppressFunc: structure.SuppressJsonDiff,
			},
		},
	}
}

func importPromptCustomText(d *schema.ResourceData, _ interface{}) ([]*schema.ResourceData, error) {
	prompt, language, err := getPromptAndLanguage(d)

	d.SetId(d.Id())
	d.Set("prompt", prompt)
	d.Set("language", language)

	return []*schema.ResourceData{d}, err
}

func createPromptCustomText(d *schema.ResourceData, m interface{}) error {
	d.SetId(d.Get("prompt").(string) + ":" + d.Get("language").(string))
	return updatePromptCustomText(d, m)
}

func readPromptCustomText(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	customText, err := api.Prompt.CustomText(d.Get("prompt").(string), d.Get("language").(string))
	if err != nil {
		if mErr, ok := err.(management.Error); ok {
			if mErr.Status() == http.StatusNotFound {
				d.SetId("")
				return nil
			}
		}
		return err
	}

	body, err := marshalCustomTextBody(customText)
	if err != nil {
		return err
	}

	d.Set("body", body)
	return nil
}

func updatePromptCustomText(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	prompt, language, err := getPromptAndLanguage(d)
	if err != nil {
		return err
	}

	if d.Get("body").(string) != "" {
		var body map[string]interface{}
		if err := json.Unmarshal([]byte(d.Get("body").(string)), &body); err != nil {
			return err
		}

		err := api.Prompt.SetCustomText(prompt, language, body)
		if err != nil {
			return err
		}
	}

	return readPromptCustomText(d, m)
}

func deletePromptCustomText(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	d.Set("body", "{}")
	return nil
}

func getPromptAndLanguage(d *schema.ResourceData) (string, string, error) {
	rawID := d.Id()
	if rawID == "" {
		return "", "", errEmptyPromptCustomTextID
	}

	if !strings.Contains(rawID, ":") {
		return "", "", errInvalidPromptCustomTextIDFormat
	}

	idPair := strings.Split(rawID, ":")
	if len(idPair) != 2 {
		return "", "", errInvalidPromptCustomTextIDFormat
	}

	return idPair[0], idPair[1], nil
}

func marshalCustomTextBody(b map[string]interface{}) (string, error) {
	bodyBytes, err := json.Marshal(b)
	if err != nil {
		return "", fmt.Errorf("Failed to serialize the custom texts to JSON: %w", err)
	}

	var buffer bytes.Buffer
	const jsonIndentation = "    "
	if err := json.Indent(&buffer, bodyBytes, "", jsonIndentation); err != nil {
		return "", fmt.Errorf("Failed to format the custom texts JSON: %w", err)
	}

	return buffer.String(), nil
}

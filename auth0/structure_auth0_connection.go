package auth0

import (
	"gopkg.in/auth0.v4/management"
	"log"
)

func flattenConnectionOptions(options interface{}) []map[string]interface{} {
	var m map[string]interface{}

	switch o := options.(type) {
	case *management.ConnectionOptions:
		m = flattenConnectionOptionsAuth0(o)
	case *management.ConnectionOptionsGoogleOAuth2:
		m = flattenConnectionOptionsGoogleOAuth2(o)
		// case *management.ConnectionOptionsFacebook:
	// 	m = flattenConnectionOptionsFacebook(o)
	// case *management.ConnectionOptionsApple:
	// 	m = flattenConnectionOptionsApple(o)
	// case *management.ConnectionOptionsLinkedin:
	// 	m = flattenConnectionOptionsLinkedin(o)
	// case *management.ConnectionOptionsGitHub:
	// 	m = flattenConnectionOptionsGitHub(o)
	// case *management.ConnectionOptionsWindowsLive:
	// 	m = flattenConnectionOptionsWindowsLive(o)
	case *management.ConnectionOptionsSalesforce:
		m = flattenConnectionOptionsSalesforce(o)
	case *management.ConnectionOptionsEmail:
		log.Printf("[1ST] Reading email options")
		m = flattenConnectionOptionsEmail(o)
	case *management.ConnectionOptionsSMS:
		m = flattenConnectionOptionsSMS(o)
		// case *management.ConnectionOptionsOIDC:
	// 	m = flattenConnectionOptionsOIDC(o)
	case *management.ConnectionOptionsAD:
		m = flattenConnectionOptionsAD(o)
	case *management.ConnectionOptionsAzureAD:
		m = flattenConnectionOptionsAzureAD(o)
	}

	return []map[string]interface{}{m}
}

func flattenConnectionOptionsAuth0(o *management.ConnectionOptions) map[string]interface{} {
	return map[string]interface{}{
		"validation":                     o.Validation,
		"password_policy":                o.GetPasswordPolicy(),
		"password_history":               toMapSlice(o.PasswordHistory),
		"password_no_personal_info":      toMapSlice(o.PasswordNoPersonalInfo),
		"password_dictionary":            toMapSlice(o.PasswordDictionary),
		"password_complexity_options":    toMapSlice(o.PasswordComplexityOptions),
		"enabled_database_customization": o.GetEnabledDatabaseCustomization(),
		"brute_force_protection":         o.GetBruteForceProtection(),
		"import_mode":                    o.GetImportMode(),
		"disable_signup":                 o.GetDisableSignup(),
		"requires_username":              o.GetRequiresUsername(),
		"custom_scripts":                 o.CustomScripts,
		"configuration":                  o.Configuration,
	}
}

func flattenConnectionOptionsGoogleOAuth2(o *management.ConnectionOptionsGoogleOAuth2) map[string]interface{} {
	return map[string]interface{}{
		"client_id":         o.GetClientID(),
		"client_secret":     o.GetClientSecret(),
		"allowed_audiences": o.AllowedAudiences,
		"scopes":            o.Scopes(),
	}
}

func flattenConnectionOptionsSalesforce(o *management.ConnectionOptionsSalesforce) map[string]interface{} {
	return map[string]interface{}{
		"client_id":          o.GetClientID(),
		"client_secret":      o.GetClientSecret(),
		"community_base_url": o.GetCommunityBaseURL(),
		"scopes":             o.Scopes(),
	}
}

func flattenConnectionOptionsSMS(o *management.ConnectionOptionsSMS) map[string]interface{} {
	return map[string]interface{}{
		"name":                   o.GetName(),
		"from":                   o.GetFrom(),
		"syntax":                 o.GetSyntax(),
		"template":               o.GetTemplate(),
		"twilio_sid":             o.GetTwilioSID(),
		"twilio_token":           o.GetTwilioToken(),
		"messaging_service_sid":  o.GetMessagingServiceSID(),
		"disable_signup":         o.GetDisableSignup(),
		"brute_force_protection": o.GetBruteForceProtection(),
		"auth_params":            o.AuthParams,
		"totp": toMapSlice(map[string]interface{}{
			"time_step": o.GetOTP().GetTimeStep(),
			"length":    o.GetOTP().GetLength(),
		}),
	}
}

func flattenConnectionOptionsEmail(o *management.ConnectionOptionsEmail) map[string]interface{} {
	return map[string]interface{}{
		"name":                   o.GetName(),
		"from":                   o.GetEmail().GetFrom(),
		"syntax":                 o.GetEmail().GetSyntax(),
		"subject":                o.GetEmail().GetSubject(),
		"template":               o.GetEmail().GetBody(),
		"disable_signup":         o.GetDisableSignup(),
		"brute_force_protection": o.GetBruteForceProtection(),
		"auth_params":            o.AuthParams,
		"totp": toMapSlice(map[string]interface{}{
			"time_step": o.GetOTP().GetTimeStep(),
			"length":    o.GetOTP().GetLength(),
		}),
	}
}

func flattenConnectionOptionsAD(o *management.ConnectionOptionsAD) map[string]interface{} {
	return map[string]interface{}{
		"tenant_domain":          o.GetTenantDomain(),
		"domain_aliases":         o.DomainAliases,
		"icon_url":               o.GetLogoURL(),
		"ips":                    o.IPs,
		"use_cert_auth":          o.GetCertAuth(),
		"use_kerberos":           o.GetKerberos(),
		"disable_cache":          o.GetDisableCache(),
		"brute_force_protection": o.GetBruteForceProtection(),
	}
}

func flattenConnectionOptionsAzureAD(o *management.ConnectionOptionsAzureAD) map[string]interface{} {
	return map[string]interface{}{
		"client_id":              o.GetClientID(),
		"client_secret":          o.GetClientSecret(),
		"app_id":                 o.GetAppID(),
		"tenant_domain":          o.GetTenantDomain(),
		"domain":                 o.GetDomain(),
		"domain_aliases":         o.DomainAliases,
		"icon_url":               o.GetLogoURL(),
		"identity_api":           o.GetIdentityAPI(),
		"waad_protocol":          o.GetWAADProtocol(),
		"waad_common_endpoint":   o.GetUseCommonEndpoint(),
		"use_wsfed":              o.GetUseWSFederation(),
		"api_enable_users":       o.GetEnableUsersAPI(),
		"max_groups_to_retrieve": o.GetMaxGroupsToRetrieve(),
		"scopes":                 o.Scopes(),
	}
}

func expandConnection(d Data) *management.Connection {

	c := &management.Connection{
		Name:               String(d, "name"),
		IsDomainConnection: Bool(d, "is_domain_connection"),
		Strategy:           String(d, "strategy"),
		EnabledClients:     Set(d, "enabled_clients").List(),
		Realms:             Slice(d, "realms"),
	}

	s := d.Get("strategy").(string)

	// Elem func will only be called if the resource is new or the 'options' element has changed so
	// it is safe to assume in any "expand*" func that all keys should be set (ignoring if changes happen or not)
	List(d, "options").Elem(func(d Data) {
		switch s {
		case management.ConnectionStrategyAuth0:
			c.Options = expandConnectionOptionsAuth0(d)
		case management.ConnectionStrategyGoogleOAuth2:
			c.Options = expandConnectionOptionsGoogleOAuth2(d)
			// case management.ConnectionStrategyFacebook
		// 	management.ConnectionStrategyApple
		// 	management.ConnectionStrategyLinkedin
		// 	management.ConnectionStrategyGitHub
		// 	management.ConnectionStrategyWindowsLive:
		case management.ConnectionStrategySalesforce,
			management.ConnectionStrategySalesforceCommunity,
			management.ConnectionStrategySalesforceSandbox:
			c.Options = expandConnectionOptionsSalesforce(d)
		case management.ConnectionStrategyEmail:
			c.Options = expandConnectionOptionsEmail(d)
		case management.ConnectionStrategySMS:
			c.Options = expandConnectionOptionsSMS(d)
			// case management.ConnectionStrategyOIDC:
		case management.ConnectionStrategyAD:
			c.Options = expandConnectionOptionsAD(d)
		case management.ConnectionStrategyAzureAD:
			c.Options = expandConnectionOptionsAzureAD(d)
		}
	})

	return c
}

func expandConnectionOptionsAuth0(d Data) *management.ConnectionOptions {
	o := &management.ConnectionOptions{
		Validation:     MapIfExists(d, "validation"),
		PasswordPolicy: StringIfExists(d, "password_policy"),
	}

	ListIfExists(d, "password_history").Elem(func(d Data) {
		o.PasswordHistory = make(map[string]interface{})
		o.PasswordHistory["enable"] = BoolIfExists(d, "enable")
		o.PasswordHistory["size"] = IntIfExists(d, "size")
	})

	ListIfExists(d, "password_no_personal_info").Elem(func(d Data) {
		o.PasswordNoPersonalInfo = make(map[string]interface{})
		o.PasswordNoPersonalInfo["enable"] = BoolIfExists(d, "enable")
	})

	ListIfExists(d, "password_dictionary").Elem(func(d Data) {
		o.PasswordDictionary = make(map[string]interface{})
		o.PasswordDictionary["enable"] = BoolIfExists(d, "enable")
		o.PasswordDictionary["dictionary"] = SetIfExists(d, "dictionary").List()
	})

	ListIfExists(d, "password_complexity_options").Elem(func(d Data) {
		o.PasswordComplexityOptions = make(map[string]interface{})
		o.PasswordComplexityOptions["min_length"] = IntIfExists(d, "min_length")
	})

	o.EnabledDatabaseCustomization = BoolIfExists(d, "enabled_database_customization")
	o.BruteForceProtection = BoolIfExists(d, "brute_force_protection")
	o.ImportMode = BoolIfExists(d, "import_mode")
	o.DisableSignup = BoolIfExists(d, "disable_signup")
	o.RequiresUsername = BoolIfExists(d, "requires_username")
	o.CustomScripts = MapIfExists(d, "custom_scripts")
	o.Configuration = MapIfExists(d, "configuration")

	return o
}

func expandConnectionOptionsGoogleOAuth2(d Data) *management.ConnectionOptionsGoogleOAuth2 {

	o := &management.ConnectionOptionsGoogleOAuth2{
		ClientID:         StringIfExists(d, "client_id"),
		ClientSecret:     StringIfExists(d, "client_secret"),
		AllowedAudiences: SetIfExists(d, "allowed_audiences").List(),
	}

	expandConnectionOptionsScopes(d, o)

	return o
}

func expandConnectionOptionsSalesforce(d Data) *management.ConnectionOptionsSalesforce {

	o := &management.ConnectionOptionsSalesforce{
		ClientID:         StringIfExists(d, "client_id"),
		ClientSecret:     StringIfExists(d, "client_secret"),
		CommunityBaseURL: StringIfExists(d, "community_base_url"),
	}

	expandConnectionOptionsScopes(d, o)

	return o
}

func expandConnectionOptionsSMS(d Data) *management.ConnectionOptionsSMS {

	o := &management.ConnectionOptionsSMS{
		Name:                 StringIfExists(d, "name"),
		From:                 StringIfExists(d, "from"),
		Syntax:               StringIfExists(d, "syntax"),
		Template:             StringIfExists(d, "template"),
		TwilioSID:            StringIfExists(d, "twilio_sid"),
		TwilioToken:          StringIfExists(d, "twilio_token"),
		MessagingServiceSID:  StringIfExists(d, "messaging_service_sid"),
		DisableSignup:        BoolIfExists(d, "disable_signup"),
		BruteForceProtection: BoolIfExists(d, "brute_force_protection"),
		AuthParams:           StringMapIfExists(d, "auth_params"),
	}

	ListIfExists(d, "totp").Elem(func(d Data) {
		o.OTP = &management.ConnectionOptionsOTP{
			TimeStep: IntIfExists(d, "time_step"),
			Length:   IntIfExists(d, "length"),
		}
	})

	return o
}

func expandConnectionOptionsEmail(d Data) *management.ConnectionOptionsEmail {
	o := &management.ConnectionOptionsEmail{
		Name:                 StringIfExists(d, "name"),
		DisableSignup:        BoolIfExists(d, "disable_signup"),
		BruteForceProtection: BoolIfExists(d, "brute_force_protection"),
		AuthParams:           StringMapIfExists(d, "auth_params"),
		Email: &management.ConnectionOptionsEmailSettings{
			From:    StringIfExists(d, "from"),
			Syntax:  StringIfExists(d, "syntax"),
			Body:    StringIfExists(d, "template"),
			Subject: StringIfExists(d, "subject"),
		},
	}
	ListIfExists(d, "totp").Elem(func(d Data) {
		o.OTP = &management.ConnectionOptionsOTP{
			TimeStep: IntIfExists(d, "time_step"),
			Length:   IntIfExists(d, "length"),
		}
	})
	return o
}

func expandConnectionOptionsAD(d Data) *management.ConnectionOptionsAD {

	o := &management.ConnectionOptionsAD{
		DomainAliases:        SetIfExists(d, "domain_aliases").List(),
		TenantDomain:         StringIfExists(d, "tenant_domain"),
		LogoURL:              StringIfExists(d, "icon_url"),
		IPs:                  SetIfExists(d, "ips").List(),
		CertAuth:             BoolIfExists(d, "use_cert_auth"),
		Kerberos:             BoolIfExists(d, "use_kerberos"),
		DisableCache:         BoolIfExists(d, "disable_cache"),
		BruteForceProtection: BoolIfExists(d, "brute_force_protection"),
	}
	return o
}

func expandConnectionOptionsAzureAD(d Data) *management.ConnectionOptionsAzureAD {

	o := &management.ConnectionOptionsAzureAD{
		ClientID:            StringIfExists(d, "client_id"),
		ClientSecret:        StringIfExists(d, "client_secret"),
		AppID:               StringIfExists(d, "app_id"),
		Domain:              StringIfExists(d, "domain"),
		DomainAliases:       SetIfExists(d, "domain_aliases").List(),
		TenantDomain:        StringIfExists(d, "tenant_domain"),
		MaxGroupsToRetrieve: StringIfExists(d, "max_groups_to_retrieve"),
		UseWSFederation:     BoolIfExists(d, "use_wsfed"),
		WAADProtocol:        StringIfExists(d, "waad_protocol"),
		UseCommonEndpoint:   BoolIfExists(d, "waad_common_endpoint"),
		EnableUsersAPI:      BoolIfExists(d, "api_enable_users"),
		LogoURL:             StringIfExists(d, "icon_url"),
		IdentityAPI:         StringIfExists(d, "identity_api"),
	}

	add, rm := Diff(d, "scopes")
	for _, scope := range add {
		o.SetScopes(true, scope.(string))
	}
	for _, scope := range rm {
		o.SetScopes(false, scope.(string))
	}

	return o
}

type scoper interface {
	Scopes() []string
	SetScopes(enable bool, scopes ...string)
}

func expandConnectionOptionsScopes(d Data, s scoper) {
	add, rm := Diff(d, "scopes")
	for _, scope := range add {
		s.SetScopes(true, scope.(string))
	}
	for _, scope := range rm {
		s.SetScopes(false, scope.(string))
	}
}

// If the given map contains any keys, return a single value slice of the map otherwise nil
func toMapSlice(m map[string]interface{}) []map[string]interface{} {
	if len(m) > 0 {
		return []map[string]interface{}{m}
	}
	return nil
}

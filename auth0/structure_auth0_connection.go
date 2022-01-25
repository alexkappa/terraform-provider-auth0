package auth0

import (
	"log"

	"gopkg.in/auth0.v5"
	"gopkg.in/auth0.v5/management"
)

func flattenConnectionOptions(d ResourceData, options interface{}) []interface{} {

	var m interface{}

	switch o := options.(type) {
	case *management.ConnectionOptions:
		m = flattenConnectionOptionsAuth0(d, o)
	case *management.ConnectionOptionsGoogleOAuth2:
		m = flattenConnectionOptionsGoogleOAuth2(o)
	case *management.ConnectionOptionsOAuth2:
		m = flattenConnectionOptionsOAuth2(o)
	case *management.ConnectionOptionsFacebook:
		m = flattenConnectionOptionsFacebook(o)
	case *management.ConnectionOptionsApple:
		m = flattenConnectionOptionsApple(o)
	case *management.ConnectionOptionsLinkedin:
		m = flattenConnectionOptionsLinkedin(o)
	case *management.ConnectionOptionsGitHub:
		m = flattenConnectionOptionsGitHub(o)
	case *management.ConnectionOptionsWindowsLive:
		m = flattenConnectionOptionsWindowsLive(o)
	case *management.ConnectionOptionsSalesforce:
		m = flattenConnectionOptionsSalesforce(o)
	case *management.ConnectionOptionsEmail:
		m = flattenConnectionOptionsEmail(o)
	case *management.ConnectionOptionsSMS:
		m = flattenConnectionOptionsSMS(o)
	case *management.ConnectionOptionsOIDC:
		m = flattenConnectionOptionsOIDC(o)
	case *management.ConnectionOptionsAD:
		m = flattenConnectionOptionsAD(o)
	case *management.ConnectionOptionsAzureAD:
		m = flattenConnectionOptionsAzureAD(o)
	case *management.ConnectionOptionsADFS:
		m = flattenConnectionOptionsADFS(o)
	case *management.ConnectionOptionsSAML:
		m = flattenConnectionOptionsSAML(o)
	}

	return []interface{}{m}
}

func flattenConnectionOptionsGitHub(o *management.ConnectionOptionsGitHub) interface{} {
	return map[string]interface{}{
		"client_id":                o.GetClientID(),
		"client_secret":            o.GetClientSecret(),
		"set_user_root_attributes": o.GetSetUserAttributes(),
		"non_persistent_attrs":     o.GetNonPersistentAttrs(),
		"scopes":                   o.Scopes(),
	}
}

func flattenConnectionOptionsWindowsLive(o *management.ConnectionOptionsWindowsLive) interface{} {
	return map[string]interface{}{
		"client_id":                o.GetClientID(),
		"client_secret":            o.GetClientSecret(),
		"scopes":                   o.Scopes(),
		"set_user_root_attributes": o.GetSetUserAttributes(),
		"non_persistent_attrs":     o.GetNonPersistentAttrs(),
		"strategy_version":         o.GetStrategyVersion(),
	}
}

func flattenConnectionOptionsAuth0(d ResourceData, o *management.ConnectionOptions) interface{} {
	return map[string]interface{}{
		"validation":                     o.Validation,
		"password_policy":                o.GetPasswordPolicy(),
		"password_history":               o.PasswordHistory,
		"password_no_personal_info":      o.PasswordNoPersonalInfo,
		"password_dictionary":            o.PasswordDictionary,
		"password_complexity_options":    o.PasswordComplexityOptions,
		"enabled_database_customization": o.GetEnabledDatabaseCustomization(),
		"brute_force_protection":         o.GetBruteForceProtection(),
		"import_mode":                    o.GetImportMode(),
		"disable_signup":                 o.GetDisableSignup(),
		"requires_username":              o.GetRequiresUsername(),
		"custom_scripts":                 o.CustomScripts,
		"mfa":                            o.MFA,
		"configuration":                  Map(d, "configuration"), // does not get read back
		"non_persistent_attrs":           o.GetNonPersistentAttrs(),
	}
}

func flattenConnectionOptionsGoogleOAuth2(o *management.ConnectionOptionsGoogleOAuth2) interface{} {
	return map[string]interface{}{
		"client_id":                o.GetClientID(),
		"client_secret":            o.GetClientSecret(),
		"allowed_audiences":        o.AllowedAudiences,
		"scopes":                   o.Scopes(),
		"set_user_root_attributes": o.GetSetUserAttributes(),
		"non_persistent_attrs":     o.GetNonPersistentAttrs(),
	}
}

func flattenConnectionOptionsOAuth2(o *management.ConnectionOptionsOAuth2) interface{} {
	return map[string]interface{}{
		"client_id":                o.GetClientID(),
		"client_secret":            o.GetClientSecret(),
		"scopes":                   o.Scopes(),
		"token_endpoint":           o.GetTokenURL(),
		"authorization_endpoint":   o.GetAuthorizationURL(),
		"scripts":                  o.Scripts,
		"set_user_root_attributes": o.GetSetUserAttributes(),
		"non_persistent_attrs":     o.GetNonPersistentAttrs(),
	}
}

func flattenConnectionOptionsFacebook(o *management.ConnectionOptionsFacebook) interface{} {
	return map[string]interface{}{
		"client_id":                o.GetClientID(),
		"client_secret":            o.GetClientSecret(),
		"scopes":                   o.Scopes(),
		"set_user_root_attributes": o.GetSetUserAttributes(),
		"non_persistent_attrs":     o.GetNonPersistentAttrs(),
	}
}

func flattenConnectionOptionsApple(o *management.ConnectionOptionsApple) interface{} {
	return map[string]interface{}{
		"client_id":                o.GetClientID(),
		"client_secret":            o.GetClientSecret(),
		"team_id":                  o.GetTeamID(),
		"key_id":                   o.GetKeyID(),
		"scopes":                   o.Scopes(),
		"set_user_root_attributes": o.GetSetUserAttributes(),
		"non_persistent_attrs":     o.GetNonPersistentAttrs(),
	}
}

func flattenConnectionOptionsLinkedin(o *management.ConnectionOptionsLinkedin) interface{} {
	return map[string]interface{}{
		"client_id":                o.GetClientID(),
		"client_secret":            o.GetClientSecret(),
		"strategy_version":         o.GetStrategyVersion(),
		"scopes":                   o.Scopes(),
		"set_user_root_attributes": o.GetSetUserAttributes(),
		"non_persistent_attrs":     o.GetNonPersistentAttrs(),
	}
}

func flattenConnectionOptionsSalesforce(o *management.ConnectionOptionsSalesforce) interface{} {
	return map[string]interface{}{
		"client_id":                o.GetClientID(),
		"client_secret":            o.GetClientSecret(),
		"community_base_url":       o.GetCommunityBaseURL(),
		"scopes":                   o.Scopes(),
		"set_user_root_attributes": o.GetSetUserAttributes(),
		"non_persistent_attrs":     o.GetNonPersistentAttrs(),
	}
}

func flattenConnectionOptionsSMS(o *management.ConnectionOptionsSMS) interface{} {
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
		"totp": map[string]interface{}{
			"time_step": o.OTP.GetTimeStep(),
			"length":    o.OTP.GetLength(),
		},
	}
}

func flattenConnectionOptionsOIDC(o *management.ConnectionOptionsOIDC) interface{} {
	return map[string]interface{}{
		"client_id":      o.GetClientID(),
		"client_secret":  o.GetClientSecret(),
		"icon_url":       o.GetLogoURL(),
		"tenant_domain":  o.GetTenantDomain(),
		"domain_aliases": o.DomainAliases,

		"type":                     o.GetType(),
		"scopes":                   o.Scopes(),
		"issuer":                   o.GetIssuer(),
		"jwks_uri":                 o.GetJWKSURI(),
		"discovery_url":            o.GetDiscoveryURL(),
		"token_endpoint":           o.GetTokenEndpoint(),
		"userinfo_endpoint":        o.GetUserInfoEndpoint(),
		"authorization_endpoint":   o.GetAuthorizationEndpoint(),
		"set_user_root_attributes": o.GetSetUserAttributes(),
		"non_persistent_attrs":     o.GetNonPersistentAttrs(),
	}
}

func flattenConnectionOptionsEmail(o *management.ConnectionOptionsEmail) interface{} {
	return map[string]interface{}{
		"name":                   o.GetName(),
		"from":                   o.GetEmail().GetFrom(),
		"syntax":                 o.GetEmail().GetSyntax(),
		"subject":                o.GetEmail().GetSubject(),
		"template":               o.GetEmail().GetBody(),
		"disable_signup":         o.GetDisableSignup(),
		"brute_force_protection": o.GetBruteForceProtection(),
		"totp": map[string]interface{}{
			"time_step": o.OTP.GetTimeStep(),
			"length":    o.OTP.GetLength(),
		},
		"set_user_root_attributes": o.GetSetUserAttributes(),
		"non_persistent_attrs":     o.GetNonPersistentAttrs(),
	}
}

func flattenConnectionOptionsAD(o *management.ConnectionOptionsAD) interface{} {
	return map[string]interface{}{
		"tenant_domain":            o.GetTenantDomain(),
		"domain_aliases":           o.DomainAliases,
		"icon_url":                 o.GetLogoURL(),
		"ips":                      o.IPs,
		"use_cert_auth":            o.GetCertAuth(),
		"use_kerberos":             o.GetKerberos(),
		"disable_cache":            o.GetDisableCache(),
		"brute_force_protection":   o.GetBruteForceProtection(),
		"set_user_root_attributes": o.GetSetUserAttributes(),
		"non_persistent_attrs":     o.GetNonPersistentAttrs(),
	}
}

func flattenConnectionOptionsAzureAD(o *management.ConnectionOptionsAzureAD) interface{} {
	return map[string]interface{}{
		"client_id":                              o.GetClientID(),
		"client_secret":                          o.GetClientSecret(),
		"app_id":                                 o.GetAppID(),
		"tenant_domain":                          o.GetTenantDomain(),
		"domain":                                 o.GetDomain(),
		"domain_aliases":                         o.DomainAliases,
		"icon_url":                               o.GetLogoURL(),
		"identity_api":                           o.GetIdentityAPI(),
		"waad_protocol":                          o.GetWAADProtocol(),
		"waad_common_endpoint":                   o.GetUseCommonEndpoint(),
		"use_wsfed":                              o.GetUseWSFederation(),
		"api_enable_users":                       o.GetEnableUsersAPI(),
		"max_groups_to_retrieve":                 o.GetMaxGroupsToRetrieve(),
		"scopes":                                 o.Scopes(),
		"set_user_root_attributes":               o.GetSetUserAttributes(),
		"non_persistent_attrs":                   o.GetNonPersistentAttrs(),
		"should_trust_email_verified_connection": o.GetTrustEmailVerified(),
	}
}

func flattenConnectionOptionsADFS(o *management.ConnectionOptionsADFS) interface{} {
	return map[string]interface{}{
		"tenant_domain":            o.GetTenantDomain(),
		"domain_aliases":           o.DomainAliases,
		"icon_url":                 o.GetLogoURL(),
		"adfs_server":              o.GetADFSServer(),
		"api_enable_users":         o.GetEnableUsersAPI(),
		"set_user_root_attributes": o.GetSetUserAttributes(),
		"non_persistent_attrs":     o.GetNonPersistentAttrs(),
	}
}

func flattenConnectionOptionsSAML(o *management.ConnectionOptionsSAML) interface{} {
	return map[string]interface{}{
		"signing_cert":     o.GetSigningCert(),
		"protocol_binding": o.GetProtocolBinding(),
		"debug":            o.GetDebug(),
		"idp_initiated": map[string]interface{}{
			"client_id":              o.IdpInitiated.GetClientID(),
			"client_protocol":        o.IdpInitiated.GetClientProtocol(),
			"client_authorize_query": o.IdpInitiated.GetClientAuthorizeQuery(),
		},
		"tenant_domain":            o.GetTenantDomain(),
		"domain_aliases":           o.DomainAliases,
		"sign_in_endpoint":         o.GetSignInEndpoint(),
		"sign_out_endpoint":        o.GetSignOutEndpoint(),
		"signature_algorithm":      o.GetSignatureAlgorithm(),
		"digest_algorithm":         o.GetDigestAglorithm(),
		"fields_map":               o.FieldsMap,
		"sign_saml_request":        o.GetSignSAMLRequest(),
		"icon_url":                 o.GetLogoURL(),
		"request_template":         o.GetRequestTemplate(),
		"user_id_attribute":        o.GetUserIDAttribute(),
		"set_user_root_attributes": o.GetSetUserAttributes(),
		"non_persistent_attrs":     o.GetNonPersistentAttrs(),
		"entity_id":                o.GetEntityID(),
	}
}

func expandConnection(d ResourceData) *management.Connection {

	c := &management.Connection{
		Name:               String(d, "name", IsNewResource()),
		DisplayName:        String(d, "display_name"),
		Strategy:           String(d, "strategy", IsNewResource()),
		IsDomainConnection: Bool(d, "is_domain_connection"),
		EnabledClients:     Set(d, "enabled_clients").List(),
		Realms:             Slice(d, "realms", IsNewResource(), HasChange()),
		ShowAsButton:       Bool(d, "show_as_button"),
	}

	s := d.Get("strategy").(string)

	List(d, "options").Elem(func(d ResourceData) {
		switch s {
		case management.ConnectionStrategyAuth0:
			c.Options = expandConnectionOptionsAuth0(d)
		case management.ConnectionStrategyGoogleOAuth2:
			c.Options = expandConnectionOptionsGoogleOAuth2(d)
		case management.ConnectionStrategyOAuth2:
			c.Options = expandConnectionOptionsOAuth2(d)
		case management.ConnectionStrategyFacebook:
			c.Options = expandConnectionOptionsFacebook(d)
		case management.ConnectionStrategyApple:
			c.Options = expandConnectionOptionsApple(d)
		case management.ConnectionStrategyLinkedin:
			c.Options = expandConnectionOptionsLinkedin(d)
		case management.ConnectionStrategyGitHub:
			c.Options = expandConnectionOptionsGitHub(d)
		case management.ConnectionStrategyWindowsLive:
			c.Options = expandConnectionOptionsWindowsLive(d)
		case management.ConnectionStrategySalesforce,
			management.ConnectionStrategySalesforceCommunity,
			management.ConnectionStrategySalesforceSandbox:
			c.Options = expandConnectionOptionsSalesforce(d)
		case management.ConnectionStrategySMS:
			c.Options = expandConnectionOptionsSMS(d)
		case management.ConnectionStrategyOIDC:
			c.Options = expandConnectionOptionsOIDC(d)
		case management.ConnectionStrategyAD:
			c.Options = expandConnectionOptionsAD(d)
		case management.ConnectionStrategyAzureAD:
			c.Options = expandConnectionOptionsAzureAD(d)
		case management.ConnectionStrategyEmail:
			c.Options = expandConnectionOptionsEmail(d)
		case management.ConnectionStrategySAML:
			c.Options = expandConnectionOptionsSAML(d)
		case management.ConnectionStrategyADFS:
			c.Options = expandConnectionOptionsADFS(d)
		default:
			log.Printf("[WARN]: Unsupported connection strategy %s", s)
			log.Printf("[WARN]: Raise an issue with the auth0 provider in order to support it:")
			log.Printf("[WARN]: 	https://github.com/alexkappa/terraform-provider-auth0/issues/new")
		}
	})

	return c
}

func expandConnectionOptionsGitHub(d ResourceData) *management.ConnectionOptionsGitHub {

	o := &management.ConnectionOptionsGitHub{
		ClientID:           String(d, "client_id"),
		ClientSecret:       String(d, "client_secret"),
		SetUserAttributes:  String(d, "set_user_root_attributes"),
		NonPersistentAttrs: castToListOfStrings(Set(d, "non_persistent_attrs").List()),
	}

	expandConnectionOptionsScopes(d, o)

	return o
}

func expandConnectionOptionsAuth0(d ResourceData) *management.ConnectionOptions {

	o := &management.ConnectionOptions{
		PasswordPolicy:     String(d, "password_policy"),
		NonPersistentAttrs: castToListOfStrings(Set(d, "non_persistent_attrs").List()),
	}

	List(d, "validation").Elem(func(d ResourceData) {
		o.Validation = make(map[string]interface{})
		List(d, "username").Elem(func(d ResourceData) {
			usernameValidation := make(map[string]*int)
			usernameValidation["min"] = Int(d, "min")
			usernameValidation["max"] = Int(d, "max")
			o.Validation["username"] = usernameValidation
		})
	})

	List(d, "password_history").Elem(func(d ResourceData) {
		o.PasswordHistory = make(map[string]interface{})
		o.PasswordHistory["enable"] = Bool(d, "enable")
		o.PasswordHistory["size"] = Int(d, "size")
	})

	List(d, "password_no_personal_info").Elem(func(d ResourceData) {
		o.PasswordNoPersonalInfo = make(map[string]interface{})
		o.PasswordNoPersonalInfo["enable"] = Bool(d, "enable")
	})

	List(d, "password_dictionary").Elem(func(d ResourceData) {
		o.PasswordDictionary = make(map[string]interface{})
		o.PasswordDictionary["enable"] = Bool(d, "enable")
		o.PasswordDictionary["dictionary"] = Set(d, "dictionary").List()
	})

	List(d, "password_complexity_options").Elem(func(d ResourceData) {
		o.PasswordComplexityOptions = make(map[string]interface{})
		o.PasswordComplexityOptions["min_length"] = Int(d, "min_length")
	})

	List(d, "mfa").Elem(func(d ResourceData) {
		o.MFA = make(map[string]interface{})
		o.MFA["active"] = Bool(d, "active")
		o.MFA["return_enroll_settings"] = Bool(d, "return_enroll_settings")
	})

	o.EnabledDatabaseCustomization = Bool(d, "enabled_database_customization")
	o.BruteForceProtection = Bool(d, "brute_force_protection")
	o.ImportMode = Bool(d, "import_mode")
	o.DisableSignup = Bool(d, "disable_signup")
	o.RequiresUsername = Bool(d, "requires_username")
	o.CustomScripts = Map(d, "custom_scripts")
	o.Configuration = Map(d, "configuration")

	return o
}

func expandConnectionOptionsGoogleOAuth2(d ResourceData) *management.ConnectionOptionsGoogleOAuth2 {

	o := &management.ConnectionOptionsGoogleOAuth2{
		ClientID:           String(d, "client_id"),
		ClientSecret:       String(d, "client_secret"),
		AllowedAudiences:   Set(d, "allowed_audiences").List(),
		SetUserAttributes:  String(d, "set_user_root_attributes"),
		NonPersistentAttrs: castToListOfStrings(Set(d, "non_persistent_attrs").List()),
	}

	expandConnectionOptionsScopes(d, o)

	return o
}
func expandConnectionOptionsOAuth2(d ResourceData) *management.ConnectionOptionsOAuth2 {

	o := &management.ConnectionOptionsOAuth2{
		ClientID:           String(d, "client_id"),
		ClientSecret:       String(d, "client_secret"),
		AuthorizationURL:   String(d, "authorization_endpoint"),
		TokenURL:           String(d, "token_endpoint"),
		SetUserAttributes:  String(d, "set_user_root_attributes"),
		NonPersistentAttrs: castToListOfStrings(Set(d, "non_persistent_attrs").List()),
	}
	o.Scripts = Map(d, "scripts")

	expandConnectionOptionsScopes(d, o)

	return o
}

func expandConnectionOptionsFacebook(d ResourceData) *management.ConnectionOptionsFacebook {

	o := &management.ConnectionOptionsFacebook{
		ClientID:           String(d, "client_id"),
		ClientSecret:       String(d, "client_secret"),
		SetUserAttributes:  String(d, "set_user_root_attributes"),
		NonPersistentAttrs: castToListOfStrings(Set(d, "non_persistent_attrs").List()),
	}

	expandConnectionOptionsScopes(d, o)

	return o
}

func expandConnectionOptionsApple(d ResourceData) *management.ConnectionOptionsApple {

	o := &management.ConnectionOptionsApple{
		ClientID:           String(d, "client_id"),
		ClientSecret:       String(d, "client_secret"),
		TeamID:             String(d, "team_id"),
		KeyID:              String(d, "key_id"),
		SetUserAttributes:  String(d, "set_user_root_attributes"),
		NonPersistentAttrs: castToListOfStrings(Set(d, "non_persistent_attrs").List()),
	}

	expandConnectionOptionsScopes(d, o)

	return o
}

func expandConnectionOptionsLinkedin(d ResourceData) *management.ConnectionOptionsLinkedin {

	o := &management.ConnectionOptionsLinkedin{
		ClientID:           String(d, "client_id"),
		ClientSecret:       String(d, "client_secret"),
		StrategyVersion:    Int(d, "strategy_version"),
		SetUserAttributes:  String(d, "set_user_root_attributes"),
		NonPersistentAttrs: castToListOfStrings(Set(d, "non_persistent_attrs").List()),
	}

	expandConnectionOptionsScopes(d, o)

	return o
}

func expandConnectionOptionsSalesforce(d ResourceData) *management.ConnectionOptionsSalesforce {

	o := &management.ConnectionOptionsSalesforce{
		ClientID:           String(d, "client_id"),
		ClientSecret:       String(d, "client_secret"),
		CommunityBaseURL:   String(d, "community_base_url"),
		SetUserAttributes:  String(d, "set_user_root_attributes"),
		NonPersistentAttrs: castToListOfStrings(Set(d, "non_persistent_attrs").List()),
	}

	expandConnectionOptionsScopes(d, o)

	return o
}

func expandConnectionOptionsWindowsLive(d ResourceData) *management.ConnectionOptionsWindowsLive {

	o := &management.ConnectionOptionsWindowsLive{
		ClientID:           String(d, "client_id"),
		ClientSecret:       String(d, "client_secret"),
		StrategyVersion:    Int(d, "strategy_version"),
		SetUserAttributes:  String(d, "set_user_root_attributes"),
		NonPersistentAttrs: castToListOfStrings(Set(d, "non_persistent_attrs").List()),
	}

	expandConnectionOptionsScopes(d, o)

	return o
}

func expandConnectionOptionsSMS(d ResourceData) *management.ConnectionOptionsSMS {

	o := &management.ConnectionOptionsSMS{
		Name:                 String(d, "name"),
		From:                 String(d, "from"),
		Syntax:               String(d, "syntax"),
		Template:             String(d, "template"),
		TwilioSID:            String(d, "twilio_sid"),
		TwilioToken:          String(d, "twilio_token"),
		MessagingServiceSID:  String(d, "messaging_service_sid"),
		DisableSignup:        Bool(d, "disable_signup"),
		BruteForceProtection: Bool(d, "brute_force_protection"),
	}

	List(d, "totp").Elem(func(d ResourceData) {
		o.OTP = &management.ConnectionOptionsOTP{
			TimeStep: Int(d, "time_step"),
			Length:   Int(d, "length"),
		}
	})

	return o
}

func expandConnectionOptionsEmail(d ResourceData) *management.ConnectionOptionsEmail {

	o := &management.ConnectionOptionsEmail{
		Name:          String(d, "name"),
		DisableSignup: Bool(d, "disable_signup"),
		Email: &management.ConnectionOptionsEmailSettings{
			Syntax:  String(d, "syntax"),
			From:    String(d, "from"),
			Subject: String(d, "subject"),
			Body:    String(d, "template"),
		},
		BruteForceProtection: Bool(d, "brute_force_protection"),
		SetUserAttributes:    String(d, "set_user_root_attributes"),
		NonPersistentAttrs:   castToListOfStrings(Set(d, "non_persistent_attrs").List()),
	}

	List(d, "totp").Elem(func(d ResourceData) {
		o.OTP = &management.ConnectionOptionsOTP{
			TimeStep: Int(d, "time_step"),
			Length:   Int(d, "length"),
		}
	})

	return o
}

func expandConnectionOptionsAD(d ResourceData) *management.ConnectionOptionsAD {

	o := &management.ConnectionOptionsAD{
		DomainAliases:      Set(d, "domain_aliases").List(),
		TenantDomain:       String(d, "tenant_domain"),
		LogoURL:            String(d, "icon_url"),
		IPs:                Set(d, "ips").List(),
		CertAuth:           Bool(d, "use_cert_auth"),
		Kerberos:           Bool(d, "use_kerberos"),
		DisableCache:       Bool(d, "disable_cache"),
		SetUserAttributes:  String(d, "set_user_root_attributes"),
		NonPersistentAttrs: castToListOfStrings(Set(d, "non_persistent_attrs").List()),
	}

	// `brute_force_protection` will default to true by the API if we don't
	// specify it. Therefore if it's not specified we'll set it to false
	// ourselves.
	v, ok := d.GetOk("brute_force_protection")
	if !ok {
		v = false
	}
	o.BruteForceProtection = auth0.Bool(v.(bool))

	return o
}

func expandConnectionOptionsAzureAD(d ResourceData) *management.ConnectionOptionsAzureAD {

	o := &management.ConnectionOptionsAzureAD{
		ClientID:            String(d, "client_id"),
		ClientSecret:        String(d, "client_secret"),
		AppID:               String(d, "app_id"),
		Domain:              String(d, "domain"),
		DomainAliases:       Set(d, "domain_aliases").List(),
		TenantDomain:        String(d, "tenant_domain"),
		MaxGroupsToRetrieve: String(d, "max_groups_to_retrieve"),
		UseWSFederation:     Bool(d, "use_wsfed"),
		WAADProtocol:        String(d, "waad_protocol"),
		UseCommonEndpoint:   Bool(d, "waad_common_endpoint"),
		EnableUsersAPI:      Bool(d, "api_enable_users"),
		LogoURL:             String(d, "icon_url"),
		IdentityAPI:         String(d, "identity_api"),
		SetUserAttributes:   String(d, "set_user_root_attributes"),
		NonPersistentAttrs:  castToListOfStrings(Set(d, "non_persistent_attrs").List()),
		TrustEmailVerified:  String(d, "should_trust_email_verified_connection"),
	}

	expandConnectionOptionsScopes(d, o)

	return o
}

func expandConnectionOptionsOIDC(d ResourceData) *management.ConnectionOptionsOIDC {

	o := &management.ConnectionOptionsOIDC{
		ClientID:              String(d, "client_id"),
		ClientSecret:          String(d, "client_secret"),
		TenantDomain:          String(d, "tenant_domain"),
		DomainAliases:         Set(d, "domain_aliases").List(),
		LogoURL:               String(d, "icon_url"),
		DiscoveryURL:          String(d, "discovery_url"),
		AuthorizationEndpoint: String(d, "authorization_endpoint"),
		Issuer:                String(d, "issuer"),
		JWKSURI:               String(d, "jwks_uri"),
		Type:                  String(d, "type"),
		UserInfoEndpoint:      String(d, "userinfo_endpoint"),
		TokenEndpoint:         String(d, "token_endpoint"),
		SetUserAttributes:     String(d, "set_user_root_attributes"),
		NonPersistentAttrs:    castToListOfStrings(Set(d, "non_persistent_attrs").List()),
	}

	expandConnectionOptionsScopes(d, o)

	return o
}

func expandConnectionOptionsSAML(d ResourceData) *management.ConnectionOptionsSAML {
	o := &management.ConnectionOptionsSAML{
		Debug:              Bool(d, "debug"),
		SigningCert:        String(d, "signing_cert"),
		ProtocolBinding:    String(d, "protocol_binding"),
		TenantDomain:       String(d, "tenant_domain"),
		DomainAliases:      Set(d, "domain_aliases").List(),
		SignInEndpoint:     String(d, "sign_in_endpoint"),
		SignOutEndpoint:    String(d, "sign_out_endpoint"),
		SignatureAlgorithm: String(d, "signature_algorithm"),
		DigestAglorithm:    String(d, "digest_algorithm"),
		FieldsMap:          Map(d, "fields_map"),
		SignSAMLRequest:    Bool(d, "sign_saml_request"),
		RequestTemplate:    String(d, "request_template"),
		UserIDAttribute:    String(d, "user_id_attribute"),
		LogoURL:            String(d, "icon_url"),
		SetUserAttributes:  String(d, "set_user_root_attributes"),
		NonPersistentAttrs: castToListOfStrings(Set(d, "non_persistent_attrs").List()),
		EntityID:           String(d, "entity_id"),
	}

	List(d, "idp_initiated").Elem(func(d ResourceData) {
		o.IdpInitiated = &management.ConnectionOptionsSAMLIdpInitiated{
			ClientID:             String(d, "client_id"),
			ClientProtocol:       String(d, "client_protocol"),
			ClientAuthorizeQuery: String(d, "client_authorize_query"),
		}
	})

	return o
}

func expandConnectionOptionsADFS(d ResourceData) *management.ConnectionOptionsADFS {
	return &management.ConnectionOptionsADFS{
		TenantDomain:       String(d, "tenant_domain"),
		DomainAliases:      Slice(d, "domain_aliases"),
		LogoURL:            String(d, "icon_url"),
		ADFSServer:         String(d, "adfs_server"),
		EnableUsersAPI:     Bool(d, "api_enable_users"),
		SetUserAttributes:  String(d, "set_user_root_attributes"),
		NonPersistentAttrs: castToListOfStrings(Set(d, "non_persistent_attrs").List()),
	}
}

type scoper interface {
	Scopes() []string
	SetScopes(enable bool, scopes ...string)
}

func expandConnectionOptionsScopes(d ResourceData, s scoper) {
	add := Set(d, "scopes").List()
	_, rm := Diff(d, "scopes")
	for _, scope := range add {
		s.SetScopes(true, scope.(string))
	}
	for _, scope := range rm.List() {
		s.SetScopes(false, scope.(string))
	}
}

func castToListOfStrings(interfaces []interface{}) *[]string {
	var strings []string
	for _, v := range interfaces {
		strings = append(strings, v.(string))
	}
	return &strings
}

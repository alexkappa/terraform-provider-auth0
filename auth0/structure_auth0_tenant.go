package auth0

import "gopkg.in/auth0.v4/management"

func flattenTenantChangePassword(changePassword *management.TenantChangePassword) []interface{} {
	m := make(map[string]interface{})
	if changePassword != nil {
		m["enabled"] = changePassword.Enabled
		m["html"] = changePassword.HTML
	}
	return []interface{}{m}
}

func flattenTenantGuardianMFAPage(mfa *management.TenantGuardianMFAPage) []interface{} {
	m := make(map[string]interface{})
	if mfa != nil {
		m["enabled"] = mfa.Enabled
		m["html"] = mfa.HTML
	}
	return []interface{}{m}
}

func flattenTenantErrorPage(errorPage *management.TenantErrorPage) []interface{} {
	m := make(map[string]interface{})
	if errorPage != nil {
		m["html"] = errorPage.HTML
		m["show_log_link"] = errorPage.ShowLogLink
		m["url"] = errorPage.URL
	}
	return []interface{}{m}
}

func flattenTenantFlags(flags *management.TenantFlags) []interface{} {
	m := make(map[string]interface{})
	if flags != nil {
		m["change_pwd_flow_v1"] = flags.ChangePasswordFlowV1
		m["enable_client_connections"] = flags.EnableClientConnections
		m["enable_apis_section"] = flags.EnableAPIsSection
		m["enable_pipeline2"] = flags.EnablePipeline2
		m["enable_dynamic_client_registration"] = flags.EnableDynamicClientRegistration
		m["enable_custom_domain_in_emails"] = flags.EnableCustomDomainInEmails
		m["universal_login"] = flags.UniversalLogin
		m["enable_legacy_logs_search_v2"] = flags.EnableLegacyLogsSearchV2
		m["disable_clickjack_protection_headers"] = flags.DisableClickjackProtectionHeaders
		m["enable_public_signup_user_exists_error"] = flags.EnablePublicSignupUserExistsError
		m["use_scope_descriptions_for_consent"] = flags.UseScopeDescriptionsForConsent
	}
	return []interface{}{m}
}

func flattenTenantUniversalLogin(universalLogin *management.TenantUniversalLogin) []interface{} {
	m := make(map[string]interface{})
	if universalLogin != nil && universalLogin.Colors != nil {
		m["colors"] = []interface{}{
			map[string]interface{}{
				"primary":         universalLogin.Colors.Primary,
				"page_background": universalLogin.Colors.PageBackground,
			},
		}
	}
	return []interface{}{m}
}

func expandTenantChangePassword(d Data) (changePassword *management.TenantChangePassword) {
	List(d, "change_password").Elem(func(d Data) {
		changePassword = &management.TenantChangePassword{
			Enabled: BoolIfExists(d, "enabled"),
			HTML:    StringIfExists(d, "html"),
		}
	})
	return
}

func expandTenantGuardianMFAPage(d Data) (mfa *management.TenantGuardianMFAPage) {
	List(d, "guardian_mfa_page").Elem(func(d Data) {
		mfa = &management.TenantGuardianMFAPage{
			Enabled: BoolIfExists(d, "enabled"),
			HTML:    StringIfExists(d, "html"),
		}
	})
	return
}

func expandTenantErrorPage(d Data) (errorPage *management.TenantErrorPage) {
	List(d, "error_page").Elem(func(d Data) {
		errorPage = &management.TenantErrorPage{
			HTML:        StringIfExists(d, "html"),
			ShowLogLink: BoolIfExists(d, "show_log_link"),
			URL:         StringIfExists(d, "url"),
		}
	})
	return
}

func expandTenantFlags(d Data) (flags *management.TenantFlags) {
	List(d, "flags").Elem(func(d Data) {
		flags = &management.TenantFlags{
			ChangePasswordFlowV1:              Bool(d, "change_pwd_flow_v1"),
			EnableClientConnections:           Bool(d, "enable_client_connections"),
			EnableAPIsSection:                 Bool(d, "enable_apis_section"),
			EnablePipeline2:                   Bool(d, "enable_pipeline2"),
			EnableDynamicClientRegistration:   Bool(d, "enable_dynamic_client_registration"),
			EnableCustomDomainInEmails:        Bool(d, "enable_custom_domain_in_emails"),
			UniversalLogin:                    Bool(d, "universal_login"),
			EnableLegacyLogsSearchV2:          Bool(d, "enable_legacy_logs_search_v2"),
			DisableClickjackProtectionHeaders: Bool(d, "disable_clickjack_protection_headers"),
			EnablePublicSignupUserExistsError: Bool(d, "enable_public_signup_user_exists_error"),
			UseScopeDescriptionsForConsent:    Bool(d, "use_scope_descriptions_for_consent"),
		}
	})
	return
}

func expandTenantUniversalLogin(d Data) (universalLogin *management.TenantUniversalLogin) {
	List(d, "universal_login").Elem(func(d Data) {
		List(d, "colors").Elem(func(d Data) {
			universalLogin = &management.TenantUniversalLogin{
				Colors: &management.TenantUniversalLoginColors{
					Primary:        StringIfExists(d, "primary"),
					PageBackground: StringIfExists(d, "page_background"),
				},
			}
		})
	})
	return
}

package auth0

import "gopkg.in/auth0.v5/management"

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

func expandTenantChangePassword(d ResourceData) (changePassword *management.TenantChangePassword) {
	List(d, "change_password").Elem(func(d ResourceData) {
		changePassword = &management.TenantChangePassword{
			Enabled: Bool(d, "enabled"),
			HTML:    String(d, "html"),
		}
	})
	return
}

func expandTenantGuardianMFAPage(d ResourceData) (mfa *management.TenantGuardianMFAPage) {
	List(d, "guardian_mfa_page").Elem(func(d ResourceData) {
		mfa = &management.TenantGuardianMFAPage{
			Enabled: Bool(d, "enabled"),
			HTML:    String(d, "html"),
		}
	})
	return
}

func expandTenantErrorPage(d ResourceData) (errorPage *management.TenantErrorPage) {
	List(d, "error_page").Elem(func(d ResourceData) {
		errorPage = &management.TenantErrorPage{
			HTML:        String(d, "html"),
			ShowLogLink: Bool(d, "show_log_link"),
			URL:         String(d, "url"),
		}
	})
	return
}

func expandTenantFlags(d ResourceData, conditions ...Condition) (flags *management.TenantFlags) {
	List(d, "flags").Elem(func(d ResourceData) {
		flags = &management.TenantFlags{
			ChangePasswordFlowV1:              Bool(d, "change_pwd_flow_v1", conditions...),
			EnableClientConnections:           Bool(d, "enable_client_connections", conditions...),
			EnableAPIsSection:                 Bool(d, "enable_apis_section", conditions...),
			EnablePipeline2:                   Bool(d, "enable_pipeline2", conditions...),
			EnableDynamicClientRegistration:   Bool(d, "enable_dynamic_client_registration", conditions...),
			EnableCustomDomainInEmails:        Bool(d, "enable_custom_domain_in_emails", conditions...),
			UniversalLogin:                    Bool(d, "universal_login", conditions...),
			EnableLegacyLogsSearchV2:          Bool(d, "enable_legacy_logs_search_v2", conditions...),
			DisableClickjackProtectionHeaders: Bool(d, "disable_clickjack_protection_headers", conditions...),
			EnablePublicSignupUserExistsError: Bool(d, "enable_public_signup_user_exists_error", conditions...),
			UseScopeDescriptionsForConsent:    Bool(d, "use_scope_descriptions_for_consent", conditions...),
		}
	})
	return
}

func expandTenantUniversalLogin(d ResourceData) (universalLogin *management.TenantUniversalLogin) {
	List(d, "universal_login").Elem(func(d ResourceData) {
		List(d, "colors").Elem(func(d ResourceData) {
			universalLogin = &management.TenantUniversalLogin{
				Colors: &management.TenantUniversalLoginColors{
					Primary:        String(d, "primary"),
					PageBackground: String(d, "page_background"),
				},
			}
		})
	})
	return
}

package auth0

import "gopkg.in/auth0.v4/management"

func flattenTenantChangePassword(cp *management.TenantChangePassword) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"enabled": cp.Enabled,
			"html":    cp.HTML,
		},
	}
}

func flattenTenantGuardianMFAPage(mfa *management.TenantGuardianMFAPage) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"enabled": mfa.Enabled,
			"html":    mfa.HTML,
		},
	}
}

func flattenTenantErrorPage(ep *management.TenantErrorPage) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"html":          ep.HTML,
			"show_log_link": ep.ShowLogLink,
			"url":           ep.URL,
		},
	}
}

func flattenTenantFlags(f *management.TenantFlags) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"change_pwd_flow_v1":                     f.ChangePasswordFlowV1,
			"enable_client_connections":              f.EnableClientConnections,
			"enable_apis_section":                    f.EnableAPIsSection,
			"enable_pipeline2":                       f.EnablePipeline2,
			"enable_dynamic_client_registration":     f.EnableDynamicClientRegistration,
			"enable_custom_domain_in_emails":         f.EnableCustomDomainInEmails,
			"universal_login":                        f.UniversalLogin,
			"enable_legacy_logs_search_v2":           f.EnableLegacyLogsSearchV2,
			"disable_clickjack_protection_headers":   f.DisableClickjackProtectionHeaders,
			"enable_public_signup_user_exists_error": f.EnablePublicSignupUserExistsError,
			"use_scope_descriptions_for_consent":     f.UseScopeDescriptionsForConsent,
		},
	}
}

func flattenTenantUniversalLogin(ul *management.TenantUniversalLogin) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"colors": []interface{}{
				map[string]interface{}{
					"primary":         ul.Colors.Primary,
					"page_background": ul.Colors.PageBackground,
				},
			},
		},
	}
}

func expandTenantChangePassword(d Data) (c *management.TenantChangePassword) {
	List(d, "change_password").Elem(func(d Data) {
		c = &management.TenantChangePassword{
			Enabled: Bool(d, "enabled"),
			HTML:    String(d, "html"),
		}
	})
	return
}

func expandTenantGuardianMFAPage(d Data) (mfa *management.TenantGuardianMFAPage) {
	List(d, "guardian_mfa_page").Elem(func(d Data) {
		mfa = &management.TenantGuardianMFAPage{
			Enabled: Bool(d, "enabled"),
			HTML:    String(d, "html"),
		}
	})
	return
}

func expandTenantErrorPage(d Data) (e *management.TenantErrorPage) {
	List(d, "error_page").Elem(func(d Data) {
		e = &management.TenantErrorPage{
			HTML:        String(d, "html"),
			ShowLogLink: Bool(d, "show_log_link"),
			URL:         String(d, "url"),
		}
	})
	return
}

func expandTenantFlags(d Data) (f *management.TenantFlags) {
	List(d, "flags").Elem(func(d Data) {
		f = &management.TenantFlags{
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

func expandTenantUniversalLogin(d Data) (u *management.TenantUniversalLogin) {
	List(d, "universal_login").Elem(func(d Data) {
		List(d, "colors").Elem(func(d Data) {
			u = &management.TenantUniversalLogin{
				Colors: &management.TenantUniversalLoginColors{
					Primary:        String(d, "primary"),
					PageBackground: String(d, "page_background"),
				},
			}
		})
	})
	return
}

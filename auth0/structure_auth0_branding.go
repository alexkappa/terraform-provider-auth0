package auth0

import (
	"gopkg.in/auth0.v5/management"
)

func flattenBrandingColors(colors *management.BrandingColors) []interface{} {
	m := make(map[string]interface{})

	if colors != nil {
		m["primary"] = colors.Primary
		m["page_background"] = colors.PageBackground

		if colors.PageBackgroundGradient != nil {
			m["page_background_gradient"] = []interface{}{
				map[string]interface{}{
					"type":      colors.PageBackgroundGradient.Type,
					"start":     colors.PageBackgroundGradient.Start,
					"end":       colors.PageBackgroundGradient.End,
					"angle_deg": colors.PageBackgroundGradient.AngleDegree,
				},
			}
		}
	}

	return []interface{}{m}
}

func flattenBrandingFont(font *management.BrandingFont) []interface{} {
	m := make(map[string]interface{})

	if font != nil {
		m["url"] = font.URL
	}

	return []interface{}{m}
}

func expandBrandingColors(d ResourceData) (colors *management.BrandingColors) {
	List(d, "colors").Elem(func(d ResourceData) {
		colors = &management.BrandingColors{
			Primary:        String(d, "primary"),
			PageBackground: String(d, "page_background"),
		}

		List(d, "page_background_gradient").Elem(func(d ResourceData) {
			colors.PageBackground = nil
			colors.PageBackgroundGradient = &management.BrandingPageBackgroundGradient{
				Type:        String(d, "type"),
				Start:       String(d, "start"),
				End:         String(d, "end"),
				AngleDegree: Int(d, "angle_deg"),
			}
		})
	})
	return
}

func expandBrandingFont(d ResourceData) (font *management.BrandingFont) {
	List(d, "font").Elem(func(d ResourceData) {
		url := String(d, "url")

		// Avoid 400 Bad Request if we send the font object with an empty URL
		if url != nil {
			font = &management.BrandingFont{
				URL: String(d, "url"),
			}
		}

	})

	return
}

package management

type Branding struct {
	// Change password page settings
	Colors *BrandingColors `json:"colors,omitempty"`

	// URL for the favicon. Must use HTTPS.
	FaviconURL *string `json:"favicon_url,omitempty"`

	// URL for the logo. Must use HTTPS.
	LogoURL *string `json:"logo_url,omitempty"`

	Font *BrandingFont `json:"font,omitempty"`
}

type BrandingColors struct {
	// Accent color
	Primary *string `json:"primary,omitempty"`
	// Page background color
	PageBackground *string `json:"page_background,omitempty"`
}

type BrandingFont struct {
	// URL for the custom font. Must use HTTPS.
	URL *string `json:"url,omitempty"`
}

type BrandingManager struct {
	*Management
}

func newBrandingManager(m *Management) *BrandingManager {
	return &BrandingManager{m}
}

// Retrieve various settings related to branding.
//
// See: https://auth0.com/docs/api/management/v2#!/Branding/get_branding
func (m *BrandingManager) Read() (*Branding, error) {
	branding := new(Branding)
	err := m.get(m.uri("branding"), branding)
	return branding, err
}

// Update various fields related to branding.
//
// See: https://auth0.com/docs/api/management/v2#!/Branding/patch_branding
func (m *BrandingManager) Update(t *Branding) (err error) {
	return m.patch(m.uri("branding"), t)
}

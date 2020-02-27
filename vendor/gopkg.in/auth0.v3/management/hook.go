package management

type Hook struct {

	// The hook's identifier.
	ID *string `json:"id,omitempty"`

	// The name of the hook. Can only contain alphanumeric characters, spaces
	// and '-'. Can neither start nor end with '-' or spaces.
	Name *string `json:"name,omitempty"`

	// A script that contains the hook's code.
	Script *string `json:"script,omitempty"`

	// The extensibility point name
	// Can currently be any of the following:
	// "credentials-exchange", "pre-user-registration",
	// "post-user-registration", "post-change-password"
	TriggerID *string `json:"triggerId,omitempty"`

	// Used to store additional metadata
	Dependencies *map[string]interface{} `json:"dependencies,omitempty"`

	// Enabled should be set to true if the hook is enabled, false otherwise.
	Enabled *bool `json:"enabled,omitempty"`
}

type HookList struct {
	List
	Hooks []*Hook `json:"hooks"`
}

type HookManager struct {
	*Management
}

func newHookManager(m *Management) *HookManager {
	return &HookManager{m}
}

// Create a new hook.
//
// Note: Changing a hook's trigger changes the signature of the script and should be done with caution.
//
// See: https://auth0.com/docs/api/management/v2#!/Hooks/post_hooks
func (m *HookManager) Create(r *Hook) error {
	return m.post(m.uri("hooks"), r)
}

// Retrieve hook details. Accepts a list of fields to include or exclude in the result.
//
// See: https://auth0.com/docs/api/management/v2/#!/Hooks/get_hooks_by_id
func (m *HookManager) Read(id string) (r *Hook, err error) {
	err = m.get(m.uri("hooks", id), &r)
	return
}

// Update an existing hook.
//
// See: https://auth0.com/docs/api/management/v2/#!/Hooks/patch_hooks_by_id
func (m *HookManager) Update(id string, r *Hook) error {
	return m.patch(m.uri("hooks", id), r)
}

// Delete a hook.
//
// See: https://auth0.com/docs/api/management/v2/#!/Hooks/delete_hooks_by_id
func (m *HookManager) Delete(id string) error {
	return m.delete(m.uri("hooks", id))
}

// List all hooks.
//
// See: https://auth0.com/docs/api/management/v2/#!/Hooks/get_hooks
func (m *HookManager) List(opts ...ListOption) (r *HookList, err error) {
	opts = m.defaults(opts)
	err = m.get(m.uri("roles")+m.q(opts), &r)
	return
}

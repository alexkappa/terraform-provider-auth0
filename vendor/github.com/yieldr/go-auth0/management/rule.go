package management

type Rule struct {

	// The rule's identifier.
	ID string `json:"id,omitempty"`

	// The name of the rule. Can only contain alphanumeric characters, spaces
	// and '-'. Can neither start nor end with '-' or spaces.
	Name string `json:"name,omitempty"`

	// A script that contains the rule's code.
	Script string `json:"script,omitempty"`

	// The rule's order in relation to other rules. A rule with a lower order
	// than another rule executes first. If no order is provided it will
	// automatically be one greater than the current maximum.
	Order int `json:"order,omitempty"`

	// Enabled should be set to true if the rule is enabled, false otherwise.
	Enabled bool `json:"enabled,omitempty"`
}

type RuleManager struct {
	m *Management
}

func NewRuleManager(m *Management) *RuleManager {
	return &RuleManager{m}
}

func (rm *RuleManager) Create(r *Rule) error {
	return rm.m.post(rm.m.getURI("rules"), r)
}

func (rm *RuleManager) Read(id string) (*Rule, error) {
	r := new(Rule)
	err := rm.m.get(rm.m.getURI("rules", id), r)
	return r, err
}

func (rm *RuleManager) Update(id string, r *Rule) (err error) {
	return rm.m.patch(rm.m.getURI("rules", id), r)
}

func (rm *RuleManager) Delete(id string) (err error) {
	return rm.m.delete(rm.m.getURI("rules", id))
}

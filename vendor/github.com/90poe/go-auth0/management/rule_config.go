package management

type RuleConfig struct {

	// The key for a RuleConfigs config
	Key string `json:"key,omitempty"`

	// The value for the rules config
	Value string `json:"value,omitempty"`
}

type RuleConfigManager struct {
	m *Management
}

func NewRuleConfigManager(m *Management) *RuleConfigManager {
	return &RuleConfigManager{m}
}

func (rm *RuleConfigManager) Upsert(key string, r *RuleConfig) (err error) {
	return rm.m.put(rm.m.getURI("rules-configs", key), r)
}

func (rm *RuleConfigManager) Read(key string) (*RuleConfig, error) {
	var rs []*RuleConfig
	err := rm.m.get(rm.m.getURI("rules-configs"), &rs)
	if err != nil {
		return nil, err
	}
	for _, r := range rs {
		if r.Key == key {
			return r, nil
		}
	}
	return nil, &managementError{
		StatusCode: 404,
		Err:        "Not Found",
		Message:    "Rule config not found",
	}
}

func (rm *RuleConfigManager) Delete(key string) (err error) {
	return rm.m.delete(rm.m.getURI("rules-configs", key))
}

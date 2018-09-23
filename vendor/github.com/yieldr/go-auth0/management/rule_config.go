package management

import (
	"encoding/json"

	auth0 "github.com/yieldr/go-auth0"
)

type RuleConfig struct {

	// The key for a RuleConfigs config
	Key *string `json:"key,omitempty"`

	// The value for the rules config
	Value *string `json:"value,omitempty"`
}

func (r *RuleConfig) String() string {
	b, _ := json.MarshalIndent(r, "", "  ")
	return string(b)
}

type RuleConfigManager struct {
	m *Management
}

func NewRuleConfigManager(m *Management) *RuleConfigManager {
	return &RuleConfigManager{m}
}

func (rm *RuleConfigManager) Upsert(key string, r *RuleConfig) (err error) {
	return rm.m.put(rm.m.uri("rules-configs", key), r)
}

func (rm *RuleConfigManager) Read(key string) (*RuleConfig, error) {
	var rs []*RuleConfig
	err := rm.m.get(rm.m.uri("rules-configs"), &rs)
	if err != nil {
		return nil, err
	}
	for _, r := range rs {
		rkey := auth0.StringValue(r.Key)
		if rkey == key {
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
	return rm.m.delete(rm.m.uri("rules-configs", key))
}

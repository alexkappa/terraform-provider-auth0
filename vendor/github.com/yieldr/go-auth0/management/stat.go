package management

import (
	"encoding/json"
	"time"
)

type StatManager struct {
	m *Management
}

func NewStatManager(m *Management) *StatManager {
	return &StatManager{m}
}

func (sm *StatManager) ActiveUsers() (int, error) {
	var i int
	err := sm.m.get(sm.m.uri("stats/active-users"), &i)
	return i, err
}

type DailyStat struct {
	Date            *time.Time `json:"date"`
	Logins          *int       `json:"logins"`
	Signups         *int       `json:"signups"`
	LeakedPasswords *int       `json:"leaked_passwords"`
	UpdatedAt       *time.Time `json:"updated_at"`
	CreatedAt       *time.Time `json:"created_at"`
}

func (ds *DailyStat) String() string {
	b, _ := json.MarshalIndent(ds, "", "  ")
	return string(b)
}

func (sm *StatManager) Daily(opts ...reqOption) ([]*DailyStat, error) {
	var ds []*DailyStat
	err := sm.m.get(sm.m.uri("stats/daily")+sm.m.q(opts), &ds)
	return ds, err
}

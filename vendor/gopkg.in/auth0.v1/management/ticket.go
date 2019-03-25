package management

import "encoding/json"

type Ticket struct {
	// The user will be redirected to this endpoint once the ticket is used
	ResultURL *string `json:"result_url,omitempty"`

	// The UserID for which the ticket is to be created
	UserID *string `json:"user_id,omitempty"`

	// The ticket's lifetime in seconds starting from the moment of creation.
	// After expiration the ticket can not be used to verify the users's email.
	// If not specified or if you send 0 the Auth0 default lifetime will be
	// applied.
	TTLSec *int `json:"ttl_sec,omitempty"`

	// The connection that provides the identity for which the password is to be
	// changed. If sending this parameter, the email is also required and the
	// UserID is invalid.
	//
	// Requires: Email
	// Conflicts with: UserID
	ConnectionID *string `json:"connection_id,omitempty"`

	// The user's email
	//
	// Requires: ConnectionID
	// Conflicts with: UserID
	Email *string `json:"email,omitempty"`

	// The URL that represents the ticket
	Ticket *string `json:"ticket,omitempty"`
}

func (t *Ticket) String() string {
	b, _ := json.MarshalIndent(t, "", "  ")
	return string(b)
}

type TicketManager struct {
	m *Management
}

func NewTicketManager(m *Management) *TicketManager {
	return &TicketManager{m}
}

func (tm *TicketManager) VerifyEmail(t *Ticket) (*Ticket, error) {
	err := tm.m.post(tm.m.uri("tickets/email-verification"), t)
	return t, err
}

func (tm *TicketManager) ChangePassword(t *Ticket) (*Ticket, error) {
	err := tm.m.post(tm.m.uri("tickets/password-change"), t)
	return t, err
}

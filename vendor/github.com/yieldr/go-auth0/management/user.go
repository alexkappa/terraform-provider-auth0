package management

import (
	"encoding/json"
	"time"
)

type User struct {

	// The users identifier.
	ID *string `json:"user_id,omitempty"`

	// The connection the user belongs to.
	Connection *string `json:"connection,omitempty"`

	// The user's email
	Email *string `json:"email,omitempty"`

	// The user's username. Only valid if the connection requires a username
	Username *string `json:"username,omitempty"`

	// The user's password (mandatory for non SMS connections)
	Password *string `json:"password,omitempty"`

	// The user's phone number (following the E.164 recommendation), only valid
	// for users to be added to SMS connections.
	PhoneNumber *string `json:"phone_number,omitempty"`

	// The time the user is created.
	CreatedAt *time.Time `json:"created_at,omitempty"`

	// The last time the user is updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`

	// The last time the user has logged in.
	LastLogin *time.Time `json:"last_login,omitempty"`

	// UserMetadata holds data that the user has read/write access to (e.g.
	// color_preference, blog_url, etc).
	UserMetadata map[string]interface{} `json:"user_metadata,omitempty"`

	// True if the user's email is verified, false otherwise. If it is true then
	// the user will not receive a verification email, unless verify_email: true
	// was specified.
	EmailVerified *bool `json:"email_verified,omitempty"`

	// If true, the user will receive a verification email after creation, even
	// if created with email_verified set to true. If false, the user will not
	// receive a verification email, even if created with email_verified set to
	// false. If unspecified, defaults to the behavior determined by the value
	// of email_verified.
	VerifyEmail *bool `json:"verify_email,omitempty"`

	// True if the user's phone number is verified, false otherwise. When the
	// user is added to a SMS connection, they will not receive an verification
	// SMS if this is true.
	PhoneVerified *bool `json:"phone_verified,omitempty"`

	// AppMetadata holds data that the user has read-only access to (e.g. roles,
	// permissions, vip, etc).
	AppMetadata map[string]interface{} `json:"app_metadata,omitempty"`
}

func (u *User) String() string {
	b, _ := json.MarshalIndent(u, "", "  ")
	return string(b)
}

type UserManager struct {
	m *Management
}

func NewUserManager(m *Management) *UserManager {
	return &UserManager{m}
}

func (um *UserManager) Create(u *User) error {
	return um.m.post(um.m.uri("users"), u)
}

func (um *UserManager) Read(id string, opts ...reqOption) (*User, error) {
	u := new(User)
	err := um.m.get(um.m.uri("users", id)+um.m.q(opts), u)
	return u, err
}

func (um *UserManager) Update(id string, u *User) (err error) {
	return um.m.patch(um.m.uri("users", id), u)
}

func (um *UserManager) Delete(id string) (err error) {
	return um.m.delete(um.m.uri("users", id))
}

func (um *UserManager) List(opts ...reqOption) (us []*User, err error) {
	err = um.m.get(um.m.uri("users")+um.m.q(opts), &us)
	return
}

func (um *UserManager) Search(opts ...reqOption) (us []*User, err error) {
	return um.List(opts...)
}

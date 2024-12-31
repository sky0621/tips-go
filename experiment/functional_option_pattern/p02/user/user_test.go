package user

import (
	"testing"
)

type UserOption func(*User)

func WithID(id int) UserOption {
	return func(u *User) {
		u.ID = id
	}
}

func createTestUser(options ...UserOption) User {
	u := User{
		ID:      1,
		Name:    "TestUser",
		Age:     20,
		Hobbies: []string{"Hobby1", "Hobby2", "Hobby3"},
	}
	for _, option := range options {
		option(&u)
	}
	return u
}

func TestAddUser(t *testing.T) {
	tests := []struct {
		testName  string
		user      User
		wantError bool
	}{
		{
			testName:  "正常系",
			user:      createTestUser(),
			wantError: false,
		},
		{
			testName:  "準正常系（ID不正）",
			user:      createTestUser(WithID(0)),
			wantError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			err := AddUser(tt.user)
			if (err != nil) != tt.wantError {
				t.Errorf("AddUser() error = %v, wantErr %v", err, tt.wantError)
			}
		})
	}
}

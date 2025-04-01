package user

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetUser(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		want    *User
		wantErr bool
	}{
		{
			name:    "1",
			args:    args{id: 1},
			want:    &User{ID: 1, Name: "User11", Age: 7},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUser(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				/*
					=== RUN   TestGetUser/1
					    user_test.go:34: GetUser() diff =   &user.User{
					          	ID:   1,
					        - 	Name: "User01",
					        + 	Name: "User11",
					        - 	Age:  6,
					        + 	Age:  7,
					          }
					--- FAIL: TestGetUser/1 (0.00s)
				*/
				t.Errorf("GetUser() diff = %v", diff)
			}
		})
	}
}

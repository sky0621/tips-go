package user

import (
	"log"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("before all test")
	m.Run()
	log.Println("after all test")
}

func TestGetUser(t *testing.T) {
	log.Println("start TestGetUser")
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
			name:    "get user",
			args:    args{id: 1},
			want:    &User{ID: 1, Name: "User01"},
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUser() got = %v, want %v", got, tt.want)
			}
		})
	}
	log.Println("end TestGetUser")
}

func TestAddUser(t *testing.T) {
	log.Println("start TestAddUser")
	type args struct {
		name int
		id   int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "add user",
			args:    args{id: 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddUser(tt.args.name, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("AddUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	log.Println("end TestAddUser")
}

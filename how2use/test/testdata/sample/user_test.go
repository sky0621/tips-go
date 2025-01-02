package sample

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestAddUser(t *testing.T) {
	f, err := os.Open("testdata/user.json")
	if err != nil {
		t.Fatal(err)
	}
	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			t.Fatal(err)
		}
	}(f)

	bytes, err := io.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}

	var users []User
	if err := json.Unmarshal(bytes, &users); err != nil {
		t.Fatal(err)
	}

	type args struct {
		u User
	}
	type sut struct {
		name    string
		args    args
		wantErr bool
	}
	var tests []sut

	for idx, user := range users {
		tests = append(tests, sut{
			name:    fmt.Sprintf("test no.%02d", idx+1),
			args:    args{u: user},
			wantErr: false,
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddUser(tt.args.u); (err != nil) != tt.wantErr {
				t.Errorf("AddUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

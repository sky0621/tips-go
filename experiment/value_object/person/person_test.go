package person

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		id   int64
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    Person
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.id, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPersonID(t *testing.T) {
	type args struct {
		v int64
	}
	tests := []struct {
		name    string
		args    args
		want    ID
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewID(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPersonName(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name    string
		args    args
		want    Name
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{v: "John Doe"},
			want: Name{v: func() ValueObject[string] {
				n, _ := NewValueObject("John Doe")
				return n
			}()},
			wantErr: false,
		},
		{
			name:    "準正常系：ブランク",
			args:    args{v: ""},
			want:    Name{},
			wantErr: true,
		},
		{
			name:    "準正常系：ハイフン",
			args:    args{v: "-"},
			want:    Name{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewName(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewName() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPerson_ID(t *testing.T) {
	type fields struct {
		id   ID
		name Name
	}
	tests := []struct {
		name   string
		fields fields
		want   ID
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Person{
				id:   tt.fields.id,
				name: tt.fields.name,
			}
			if got := p.ID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPerson_IDValue(t *testing.T) {
	type fields struct {
		id   ID
		name Name
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Person{
				id:   tt.fields.id,
				name: tt.fields.name,
			}
			if got := p.IDValue(); got != tt.want {
				t.Errorf("IDValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPerson_Name(t *testing.T) {
	type fields struct {
		id   ID
		name Name
	}
	tests := []struct {
		name   string
		fields fields
		want   Name
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Person{
				id:   tt.fields.id,
				name: tt.fields.name,
			}
			if got := p.Name(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPerson_NameValue(t *testing.T) {
	type fields struct {
		id   ID
		name Name
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Person{
				id:   tt.fields.id,
				name: tt.fields.name,
			}
			if got := p.NameValue(); got != tt.want {
				t.Errorf("NameValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPersonName(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPersonName(tt.args.v); got != tt.want {
				t.Errorf("isPersonName() = %v, want %v", got, tt.want)
			}
		})
	}
}

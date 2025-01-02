package sample

import "testing"

func TestSimple(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{i: 1},
			want: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Simple(tt.args.i); got != tt.want {
				t.Errorf("Simple() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping in short mode")
	}
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{i: 1},
			want: "plus",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Complex(tt.args.i); got != tt.want {
				t.Errorf("Complex() = %v, want %v", got, tt.want)
			}
		})
	}
}

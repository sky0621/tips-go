package sample

import "testing"

func TestForParallel(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "1",
			want: "done",
		},
		{
			name: "2",
			want: "done",
		},
		{
			name: "3",
			want: "done",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := ForParallel(); got != tt.want {
				t.Errorf("ForParallel() = %v, want %v", got, tt.want)
			}
		})
	}
}

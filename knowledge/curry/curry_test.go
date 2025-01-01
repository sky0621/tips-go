package main

import (
	"testing"
)

func Test_fn(t *testing.T) {
	tests := []struct {
		name   string
		inputB bool
		inputI int
		inputS string
		want   string
	}{
		{
			name:   "T+S",
			inputB: true, inputI: 1, inputS: "S",
			want: "[T][+][S]",
		},
		{
			name:   "T-S2",
			inputB: true, inputI: -1, inputS: "S2",
			want: "[T][-][S2]",
		},
		{
			name:   "F+S3",
			inputB: false, inputI: 1, inputS: "S3",
			want: "[F][+][S3]",
		},
		{
			name:   "F-S4",
			inputB: false, inputI: -1, inputS: "S4",
			want: "[F][-][S4]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := fn(tt.inputB)(tt.inputI)(tt.inputS)

			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

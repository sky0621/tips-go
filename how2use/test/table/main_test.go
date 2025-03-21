package main

import "testing"

func TestGetCentury(t *testing.T) {
	tests := []struct {
		name    string
		year    int
		want    int
		wantErr bool
	}{
		{"1000年", 1000, 10, false},
		{"1999年", 1999, 20, false},
		{"2000年", 2000, 20, false},
		{"2021年", 2021, 21, false},
		{"999年", 999, -1, true},
		{"10000年", 10000, -1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getCentury(tt.year)
			if (err != nil) != tt.wantErr {
				t.Errorf("getCentury() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getCentury() = %v, want %v", got, tt.want)
			}
		})
	}
}

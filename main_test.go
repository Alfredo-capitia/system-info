
package main

import "testing"

func TestCheckPort(t *testing.T) {
	tests := []struct {
		name    string
		port    int
		wantErr bool
	}{
		{name: "valid port", port: 8080, wantErr: false},
		{name: "minimum valid port", port: 1, wantErr: false},
		{name: "maximum valid port", port: 65535, wantErr: false},
		{name: "port zero", port: 0, wantErr: true},
		{name: "negative port", port: -1, wantErr: true},
		{name: "port above maximum", port: 65536, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkPort(tt.port)

			if (err != nil) != tt.wantErr {
				t.Errorf(
					"checkPort(%d) error = %v, wantErr = %v",
					tt.port,
					err,
					tt.wantErr,
				)
			}
		})
	}
}

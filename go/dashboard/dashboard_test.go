package dashboard

import (
	"testing"
)

func Test_validateTimezone(t *testing.T) {
	tests := []struct {
		name     string
		timezone string
		wantErr  bool
	}{
		{
			name:     "empty string is valid",
			timezone: "",
			wantErr:  false,
		},
		{
			name:     "local is valid",
			timezone: "local",
			wantErr:  false,
		},
		{
			name:     "valid IANA timezone",
			timezone: "America/New_York",
			wantErr:  false,
		},
		{
			name:     "valid UTC",
			timezone: "UTC",
			wantErr:  false,
		},
		{
			name:     "invalid timezone string",
			timezone: "Not/A_Real_Place_At_Moon",
			wantErr:  true,
		},
		{
			name:     "garbage input",
			timezone: "12345",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateTimezone(tt.timezone)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateTimezone() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

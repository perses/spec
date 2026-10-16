package dashboard

import (
	"encoding/json"
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

func Test_validateQueryBatching(t *testing.T) {
	tests := []struct {
		name    string
		q       *QueryBatchingSpec
		wantErr bool
	}{
		{name: "nil is valid", q: nil, wantErr: false},
		{name: "empty mode", q: &QueryBatchingSpec{}, wantErr: false},
		{name: "mode off", q: &QueryBatchingSpec{Mode: "off"}, wantErr: false},
		{name: "mode panel", q: &QueryBatchingSpec{Mode: "panel", MaxPerRequest: 16}, wantErr: false},
		{name: "mode viewport", q: &QueryBatchingSpec{Mode: "viewport"}, wantErr: false},
		{name: "mode dashboard", q: &QueryBatchingSpec{Mode: "dashboard"}, wantErr: false},
		{name: "invalid mode", q: &QueryBatchingSpec{Mode: "all"}, wantErr: true},
		{name: "negative maxPerRequest", q: &QueryBatchingSpec{Mode: "panel", MaxPerRequest: -1}, wantErr: true},
		{name: "zero maxPerRequest ok", q: &QueryBatchingSpec{Mode: "panel", MaxPerRequest: 0}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateQueryBatching(tt.q)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateQueryBatching() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSpec_UnmarshalJSON_QueryBatching(t *testing.T) {
	okJSON := []byte(`{
		"duration": "1h",
		"panels": {},
		"layouts": [],
		"queryBatching": { "mode": "panel", "maxPerRequest": 8 }
	}`)
	var ok Spec
	if err := json.Unmarshal(okJSON, &ok); err != nil {
		t.Fatalf("expected valid unmarshal: %v", err)
	}
	if ok.QueryBatching == nil || ok.QueryBatching.Mode != "panel" || ok.QueryBatching.MaxPerRequest != 8 {
		t.Fatalf("unexpected queryBatching: %+v", ok.QueryBatching)
	}

	badJSON := []byte(`{
		"duration": "1h",
		"panels": {},
		"layouts": [],
		"queryBatching": { "mode": "nope" }
	}`)
	var bad Spec
	if err := json.Unmarshal(badJSON, &bad); err == nil {
		t.Fatal("expected error for invalid mode")
	}
}

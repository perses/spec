package dashboard

import (
	"encoding/json"
	"testing"

	"github.com/perses/spec/go/datasource"
	"github.com/perses/spec/go/plugin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestUnmarshalSpec(t *testing.T) {
	tests := []struct {
		name    string
		json    string
		yaml    string
		want    Spec
		wantErr string
	}{
		{
			name: "panel and datasource",
			json: `{"datasources":{"prom":{"default":true,"plugin":{"kind":"PrometheusDatasource","spec":{}}}},"panels":{"cpu":{"kind":"Panel","spec":{"plugin":{"kind":"TimeSeriesChart","spec":{}}}}}}`,
			yaml: "{datasources: {prom: {default: true, plugin: {kind: PrometheusDatasource, spec: {}}}}, panels: {cpu: {kind: Panel, spec: {plugin: {kind: TimeSeriesChart, spec: {}}}}}}",
			want: Spec{
				Datasources: map[string]*datasource.Spec{
					"prom": {Default: true, Plugin: plugin.Plugin{Kind: "PrometheusDatasource", Spec: map[string]any{}}},
				},
				Panels: map[string]*Panel{
					"cpu": {Kind: "Panel", Spec: PanelSpec{Plugin: plugin.Plugin{Kind: "TimeSeriesChart", Spec: map[string]any{}}}},
				},
				Duration: "1h",
			},
		},
		{
			name: "no panel",
			json: `{"panels":{}}`,
			yaml: "{panels: {}}",
			want: Spec{Panels: map[string]*Panel{}, Duration: "1h"},
		},
		{
			name:    "null panel",
			json:    `{"panels":{"cpu":null}}`,
			yaml:    "{panels: {cpu: null}}",
			wantErr: `panel "cpu" must not be null`,
		},
		{
			name:    "null datasource",
			json:    `{"datasources":{"prom":null},"panels":{}}`,
			yaml:    "datasources:\n  prom:\npanels: {}\n",
			wantErr: `datasource "prom" must not be null`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, format := range []struct {
				name string
				raw  string
				call func([]byte, any) error
			}{
				{"json", tt.json, json.Unmarshal},
				{"yaml", tt.yaml, yaml.Unmarshal},
			} {
				t.Run(format.name, func(t *testing.T) {
					var spec Spec
					err := format.call([]byte(format.raw), &spec)
					if tt.wantErr != "" {
						require.EqualError(t, err, tt.wantErr)
						return
					}
					require.NoError(t, err)
					assert.Equal(t, tt.want, spec)
				})
			}
		})
	}
}

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

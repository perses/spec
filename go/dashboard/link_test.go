// Copyright The Perses Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dashboard

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestUnmarshalLink(t *testing.T) {
	tests := []struct {
		name    string
		json    string
		yaml    string
		want    Link
		wantErr string
	}{
		{
			name: "url",
			json: `{"name":" Docs ","url":" https://perses.dev ","tooltip":"docs","renderVariables":true,"targetBlank":true}`,
			yaml: "{name: ' Docs ', url: ' https://perses.dev ', tooltip: docs, renderVariables: true, targetBlank: true}",
			want: Link{Name: "Docs", URL: "https://perses.dev", Tooltip: "docs", RenderVariables: true, TargetBlank: true},
		},
		{
			name: "tags",
			json: `{"name":" Service detail ","url":"   ","tags":[" traffic ","billing"],"keepTime":true,"includeVariables":true}`,
			yaml: "{name: ' Service detail ', url: '   ', tags: [' traffic ', billing], keepTime: true, includeVariables: true}",
			want: Link{Name: "Service detail", Tags: []string{"traffic", "billing"}, KeepTime: true, IncludeVariables: true},
		},
		{
			name:    "neither",
			json:    `{"name":"Docs"}`,
			yaml:    "{name: Docs}",
			wantErr: "link must set a url or at least one tag",
		},
		{
			name:    "blank url",
			json:    `{"url":"   "}`,
			yaml:    "{url: '   '}",
			wantErr: "link must set a url or at least one tag",
		},
		{
			name:    "both",
			json:    `{"name":"Docs","url":" https://perses.dev ","tags":["traffic"]}`,
			yaml:    "{name: Docs, url: ' https://perses.dev ', tags: [traffic]}",
			wantErr: "link must set either a url or tags, not both",
		},
		{
			name:    "empty tag",
			json:    `{"name":"Docs","tags":["traffic"," "]}`,
			yaml:    "{name: Docs, tags: [traffic, ' ']}",
			wantErr: "link tag must not be empty",
		},
		{
			name:    "tag without name",
			json:    `{"name":"  ","tags":["traffic"]}`,
			yaml:    "{name: '  ', tags: [traffic]}",
			wantErr: "a tag link must set a name",
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
					var link Link
					err := format.call([]byte(format.raw), &link)
					if tt.wantErr != "" {
						require.ErrorContains(t, err, tt.wantErr)
						return
					}
					require.NoError(t, err)
					assert.Equal(t, tt.want, link)
					if len(link.Tags) > 0 {
						data, err := json.Marshal(link)
						require.NoError(t, err)
						assert.NotContains(t, string(data), `"url"`)
					}
				})
			}
		})
	}
}

func TestUnmarshalPanelLink(t *testing.T) {
	const raw = `{
	  "duration": "1h",
	  "layouts": [],
	  "panels": {
	    "cpu": {
	      "kind": "Panel",
	      "spec": {
	        "plugin": {"kind": "TimeSeriesChart", "spec": {}},
	        "links": [{"name":" Detail ","url":"  ","tags":[" traffic "],"keepTime":true}]
	      }
	    }
	  }
	}`
	var spec Spec
	require.NoError(t, json.Unmarshal([]byte(raw), &spec))
	assert.Equal(t, Link{Name: "Detail", Tags: []string{"traffic"}, KeepTime: true}, spec.Panels["cpu"].Spec.Links[0])

	err := json.Unmarshal([]byte(`{
	  "duration": "1h",
	  "layouts": [],
	  "panels": {"cpu": {"kind": "Panel", "spec": {"plugin": {"kind": "TimeSeriesChart", "spec": {}}, "links": [{"tags": ["traffic"]}]}}}
	}`), &spec)
	require.ErrorContains(t, err, "a tag link must set a name")
}

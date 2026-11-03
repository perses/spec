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
			name: "neither",
			json: `{"name":"Docs"}`,
			yaml: "{name: Docs}",
			want: Link{Name: "Docs"},
		},
		{
			name: "blank url",
			json: `{"name":"todo","url":""}`,
			yaml: "{name: todo, url: ''}",
			want: Link{Name: "todo"},
		},
		{
			name: "whitespace url",
			json: `{"url":"   "}`,
			yaml: "{url: '   '}",
			want: Link{},
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
		{
			name: "tag target blank",
			json: `{"name":"Traffic","tags":["traffic"],"targetBlank":true}`,
			yaml: "{name: Traffic, tags: [traffic], targetBlank: true}",
			want: Link{Name: "Traffic", Tags: []string{"traffic"}, TargetBlank: true},
		},
		{
			name:    "too many tags",
			json:    `{"name":"Traffic","tags":["1","2","3","4","5","6","7","8","9","10","11","12","13","14","15","16","17","18","19","20","21"]}`,
			yaml:    "{name: Traffic, tags: ['1','2','3','4','5','6','7','8','9','10','11','12','13','14','15','16','17','18','19','20','21']}",
			wantErr: "cannot contain more than 20 tags",
		},
		{
			name:    "tag too long",
			json:    `{"name":"Traffic","tags":["aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"]}`,
			yaml:    "{name: Traffic, tags: ['aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa']}",
			wantErr: "cannot contain more than 50 characters",
		},
		{
			name:    "tag render variables",
			json:    `{"name":"Traffic","tags":["traffic"],"renderVariables":true}`,
			yaml:    "{name: Traffic, tags: [traffic], renderVariables: true}",
			wantErr: "renderVariables is only valid on a URL link",
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
					data, err := json.Marshal(link)
					require.NoError(t, err)
					yamlData, err := yaml.Marshal(link)
					require.NoError(t, err)
					if len(link.Tags) > 0 {
						assert.NotContains(t, string(data), `"url"`)
						assert.NotContains(t, string(yamlData), "url:")
						return
					}
					assert.Contains(t, string(data), `"url"`)
					assert.Contains(t, string(yamlData), "url:")
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

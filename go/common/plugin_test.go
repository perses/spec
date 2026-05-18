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

package common

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestPluginUnmarshallYAML(t *testing.T) {
	tests := []struct {
		title    string
		yml      string
		expected *Plugin
	}{
		{
			title: "empty spec",
			yml: `kind: TraceTable
spec: {}`,
			expected: &Plugin{
				Kind: "TraceTable",
				Spec: map[string]interface{}{},
			},
		},
		{
			title: "null spec",
			yml: `kind: TraceTable
spec: null`,
			expected: &Plugin{
				Kind: "TraceTable",
				Spec: nil,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.title, func(t *testing.T) {
			plg := &Plugin{}
			err := yaml.Unmarshal([]byte(tc.yml), plg)
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, plg)
		})
	}

}

func TestPluginUnmarshallJSON(t *testing.T) {
	jason := `{
"kind": "TraceTable",
"spec": {}
}`
	plg := &Plugin{}
	err := json.Unmarshal([]byte(jason), plg)
	assert.NoError(t, err)
	assert.Equal(t, "TraceTable", plg.Kind)
	assert.Equal(t, map[string]interface{}{}, plg.Spec)
}

func TestPluginMarshalYAML(t *testing.T) {
	tests := []struct {
		title    string
		plg      Plugin
		expected string
	}{
		{
			title: "empty spec",
			plg: Plugin{
				Kind: "TraceTable",
				Spec: map[string]interface{}{},
			},
			expected: `kind: TraceTable
spec: {}
`,
		},
		{
			title: "null spec",
			plg: Plugin{
				Kind: "TraceTable",
				Spec: nil,
			},
			expected: `kind: TraceTable
spec: {}
`,
		},
	}
	for _, tc := range tests {
		t.Run(tc.title, func(t *testing.T) {
			b, err := yaml.Marshal(tc.plg)
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, string(b))
		})
	}
}

func TestPluginMarshalJSON(t *testing.T) {
	tests := []struct {
		title    string
		plg      Plugin
		expected string
	}{
		{
			title: "empty spec",
			plg: Plugin{
				Kind: "TraceTable",
				Spec: map[string]interface{}{},
			},
			expected: `{
  "kind": "TraceTable",
  "spec": {}
}`,
		},
		{
			title: "null spec",
			plg: Plugin{
				Kind: "TraceTable",
				Spec: nil,
			},
			expected: `{
  "kind": "TraceTable",
  "spec": {}
}`,
		},
	}
	for _, tc := range tests {
		t.Run(tc.title, func(t *testing.T) {
			b, err := json.MarshalIndent(tc.plg, "", "  ")
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, string(b))
		})
	}
}

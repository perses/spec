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

// A panel carrying a panel-local annotation. This exercises the new PanelSpec.Annotations field,
// which mirrors the dashboard-level Spec.Annotations shape.
const panelWithAnnotationsJSON = `{
  "plugin": {"kind": "TimeSeriesChart", "spec": {}},
  "annotations": [
    {
      "display": {"name": "Deploys"},
      "plugin": {"kind": "PrometheusPromQLAnnotation", "spec": {}}
    }
  ]
}`

func TestPanelSpecAnnotationsRoundTripJSON(t *testing.T) {
	var spec PanelSpec
	require.NoError(t, json.Unmarshal([]byte(panelWithAnnotationsJSON), &spec))

	require.Len(t, spec.Annotations, 1)
	assert.Equal(t, "Deploys", spec.Annotations[0].Display.Name)
	assert.Equal(t, "PrometheusPromQLAnnotation", spec.Annotations[0].Plugin.Kind)

	// Marshal back and unmarshal again; the annotations must survive the round trip.
	data, err := json.Marshal(spec)
	require.NoError(t, err)

	var reparsed PanelSpec
	require.NoError(t, json.Unmarshal(data, &reparsed))
	assert.Equal(t, spec, reparsed)
}

func TestPanelSpecAnnotationsOmittedIsNil(t *testing.T) {
	// A panel without an annotations block leaves the field nil. Only dashboard-level annotations
	// apply to it, which is the default behavior.
	var spec PanelSpec
	require.NoError(t, json.Unmarshal([]byte(`{"plugin": {"kind": "TimeSeriesChart", "spec": {}}}`), &spec))
	assert.Nil(t, spec.Annotations)
}

func TestPanelSpecAnnotationsRejectsEmptyName(t *testing.T) {
	// The panel-local annotations reuse AnnotationSpec, whose UnmarshalJSON validates a non-empty
	// name. That validation fires automatically for each element, so no separate plumbing is needed
	// on PanelSpec.
	const invalid = `{
      "plugin": {"kind": "TimeSeriesChart", "spec": {}},
      "annotations": [
        {"display": {"name": ""}, "plugin": {"kind": "PrometheusPromQLAnnotation", "spec": {}}}
      ]
    }`
	var spec PanelSpec
	err := json.Unmarshal([]byte(invalid), &spec)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "annotation name cannot be empty")
}

func TestPanelSpecAnnotationsRoundTripYAML(t *testing.T) {
	const in = `
plugin:
  kind: TimeSeriesChart
  spec: {}
annotations:
  - display:
      name: Incidents
    plugin:
      kind: PrometheusPromQLAnnotation
      spec: {}
`
	var spec PanelSpec
	require.NoError(t, yaml.Unmarshal([]byte(in), &spec))
	require.Len(t, spec.Annotations, 1)
	assert.Equal(t, "Incidents", spec.Annotations[0].Display.Name)
}

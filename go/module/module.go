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

// Package module is here to define any struct necessary to build a plugin compatible with Perses.
// The definition of a plugin is different to the one you can find in the dashboard definition.
// The plugin definition in the dashboard is more about the configuration of a plugin, while the one in the module is more about the definition of a plugin.
// The module definition is stricter and is used to validate the plugin before it is loaded into Perses.
package module

import (
	"encoding/json"
	"fmt"

	"github.com/perses/spec/go/common"
)

type Kind string

const (
	KindVariable        Kind = "Variable"
	KindDatasource      Kind = "Datasource"
	KindPanel           Kind = "Panel"
	KindTimeSeriesQuery Kind = "TimeSeriesQuery"
	KindTraceQuery      Kind = "TraceQuery"
	KindProfileQuery    Kind = "ProfileQuery"
	KindLogQuery        Kind = "LogQuery"
	KindQuery           Kind = "Query"
	KindAlertsQuery     Kind = "AlertsQuery"
	KindSilencesQuery   Kind = "SilencesQuery"
	KindExplore         Kind = "Explore"
	KindAnnotation      Kind = "Annotation"
)

var kindMap = map[Kind]bool{
	KindVariable:        true,
	KindDatasource:      true,
	KindPanel:           true,
	KindTimeSeriesQuery: true,
	KindTraceQuery:      true,
	KindProfileQuery:    true,
	KindLogQuery:        true,
	KindQuery:           true,
	KindAlertsQuery:     true,
	KindSilencesQuery:   true,
	KindExplore:         true,
	KindAnnotation:      true,
}

func (k Kind) IsQuery() bool {
	return k == KindQuery ||
		k == KindTimeSeriesQuery ||
		k == KindTraceQuery ||
		k == KindProfileQuery ||
		k == KindLogQuery ||
		k == KindAlertsQuery ||
		k == KindSilencesQuery
}

type PluginSpec struct {
	Display *common.Display `json:"display" yaml:"display"`
	Name    string          `json:"name" yaml:"name"`
}

type Plugin struct {
	Kind Kind       `json:"kind" yaml:"kind"`
	Spec PluginSpec `json:"spec" yaml:"spec"`
}

func (p *Plugin) UnmarshalJSON(data []byte) error {
	var tmp Plugin
	type plain Plugin
	if err := json.Unmarshal(data, (*plain)(&tmp)); err != nil {
		return err
	}
	if err := (&tmp).validate(); err != nil {
		return err
	}
	*p = tmp
	return nil
}

func (p *Plugin) UnmarshalYAML(unmarshal func(any) error) error {
	var tmp Plugin
	type plain Plugin
	if err := unmarshal((*plain)(&tmp)); err != nil {
		return err
	}
	if err := (&tmp).validate(); err != nil {
		return err
	}
	*p = tmp
	return nil
}

func (p *Plugin) validate() error {
	if !kindMap[p.Kind] {
		return fmt.Errorf("invalid plugin kind %s", p.Kind)
	}
	return nil
}

type Status struct {
	IsLoaded bool   `json:"isLoaded" yaml:"isLoaded"`
	InDev    bool   `json:"inDev" yaml:"inDev"`
	Error    string `json:"error,omitempty" yaml:"error,omitempty"`
}

type Module struct {
	Metadata    *Metadata `json:"metadata,omitempty" yaml:"metadata,omitempty"`
	SchemasPath string    `json:"schemasPath" yaml:"schemasPath"`
	// ModuleName is deprecated and will be removed in the future. It is recommended to use the Metadata.Name field instead.
	ModuleName string `json:"moduleName,omitempty" yaml:"moduleName,omitempty"`
	// ModuleOrg is deprecated and will be removed in the future. It is recommended to use the Metadata.Registry field instead.
	ModuleOrg string   `json:"moduleOrg,omitempty" yaml:"moduleOrg,omitempty"`
	Plugins   []Plugin `json:"plugins" yaml:"plugins"`
}

func (m *Module) UnmarshalJSON(data []byte) error {
	var tmp Module
	type plain Module
	if err := json.Unmarshal(data, (*plain)(&tmp)); err != nil {
		return err
	}
	if err := (&tmp).validate(); err != nil {
		return err
	}
	*m = tmp
	return nil
}

func (m *Module) UnmarshalYAML(unmarshal func(any) error) error {
	var tmp Module
	type plain Module
	if err := unmarshal((*plain)(&tmp)); err != nil {
		return err
	}
	if err := (&tmp).validate(); err != nil {
		return err
	}
	*m = tmp
	return nil
}

func (m *Module) validate() error {
	if len(m.Plugins) == 0 {
		return fmt.Errorf("the module spec must have at least one plugin")
	}
	if len(m.SchemasPath) == 0 {
		m.SchemasPath = "schemas"
	}
	return nil
}

type Metadata struct {
	Name     string `json:"name" yaml:"name"`
	Version  string `json:"version" yaml:"version"`
	Registry string `json:"registry" yaml:"registry"`
}

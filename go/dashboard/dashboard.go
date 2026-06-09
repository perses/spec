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
	"fmt"
	"time"

	"github.com/perses/spec/go/common"
	"github.com/perses/spec/go/datasource"
	"github.com/perses/spec/go/plugin"
)

type Link struct {
	Name            string `json:"name,omitempty" yaml:"name,omitempty"`
	URL             string `json:"url" yaml:"url"`
	Tooltip         string `json:"tooltip,omitempty" yaml:"tooltip,omitempty"`
	RenderVariables bool   `json:"renderVariables,omitempty" yaml:"renderVariables,omitempty"`
	TargetBlank     bool   `json:"targetBlank,omitempty" yaml:"targetBlank,omitempty"`
}

type PanelDisplay struct {
	Name        string `json:"name,omitempty" yaml:"name,omitempty"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}

type PanelSpec struct {
	Display *PanelDisplay `json:"display,omitempty" yaml:"display,omitempty"`
	Plugin  plugin.Plugin `json:"plugin" yaml:"plugin"`
	Queries []Query       `json:"queries,omitempty" yaml:"queries,omitempty"`
	Links   []Link        `json:"links,omitempty" yaml:"links,omitempty"`
}

type Panel struct {
	Kind string    `json:"kind" yaml:"kind"`
	Spec PanelSpec `json:"spec" yaml:"spec"`
}

type Query struct {
	Kind string    `json:"kind" yaml:"kind"`
	Spec QuerySpec `json:"spec" yaml:"spec"`
}

type QuerySpec struct {
	Name   string        `json:"name,omitempty" yaml:"name,omitempty"`
	Plugin plugin.Plugin `json:"plugin" yaml:"plugin"`
}

type Spec struct {
	Display *common.Display `json:"display,omitempty" yaml:"display,omitempty"`
	// Datasources is an optional list of datasource definition.
	Datasources map[string]*datasource.Spec `json:"datasources,omitempty" yaml:"datasources,omitempty"`
	Variables   []Variable                  `json:"variables,omitempty" yaml:"variables,omitempty"`
	Annotations []AnnotationSpec            `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Panels      map[string]*Panel           `json:"panels" yaml:"panels"`
	Layouts     []Layout                    `json:"layouts" yaml:"layouts"`
	// Duration is the default time range to use when getting data to fill the dashboard
	Duration common.DurationString `json:"duration" yaml:"duration"`
	// RefreshInterval is the default refresh interval to use when landing on the dashboard
	RefreshInterval common.DurationString `json:"refreshInterval,omitempty" yaml:"refreshInterval,omitempty"`
	// Timezone is the dashboard level timezone
	Timezone string `json:"timezone,omitempty" yaml:"timezone,omitempty"`
	// Links is an optional list of links to display at the dashboard level
	Links []Link `json:"links,omitempty" yaml:"links,omitempty"`
}

func (d *Spec) UnmarshalJSON(data []byte) error {
	var tmp Spec
	type plain Spec
	if err := json.Unmarshal(data, (*plain)(&tmp)); err != nil {
		return err
	}
	if err := (&tmp).validate(); err != nil {
		return err
	}
	*d = tmp
	return nil
}

func (d *Spec) UnmarshalYAML(unmarshal func(any) error) error {
	var tmp Spec
	type plain Spec
	if err := unmarshal((*plain)(&tmp)); err != nil {
		return err
	}
	if err := (&tmp).validate(); err != nil {
		return err
	}
	*d = tmp
	return nil
}

func validateTimezone(timezone string) error {
	if timezone == "" || timezone == "local" {
		return nil
	}

	_, err := time.LoadLocation(timezone)
	if err != nil {
		return fmt.Errorf("%q is not a valid timezone: %w", timezone, err)
	}
	return nil
}

func (d *Spec) validate() error {
	variables := make(map[string]bool, len(d.Variables))
	for i, variable := range d.Variables {
		name := variable.Spec.GetName()
		if !variables[name] {
			variables[name] = true
		} else {
			return fmt.Errorf("variable %q (index %d) already exists", name, i)
		}
	}
	for panelKey := range d.Panels {
		if err := common.ValidateID(panelKey); err != nil {
			return err
		}
	}
	if len(d.Duration) == 0 {
		d.Duration = "1h"
	}
	if err := validateTimezone(d.Timezone); err != nil {
		return err
	}
	return nil
}

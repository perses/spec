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

	"github.com/perses/spec/go/common"
	"gopkg.in/yaml.v3"
)

type LayoutKind string

const (
	KindGridLayout LayoutKind = "Grid"
	KindTabLayout  LayoutKind = "Tabs"
)

var layoutKindMap = map[LayoutKind]bool{
	KindGridLayout: true,
	KindTabLayout:  true,
}

func (k *LayoutKind) UnmarshalJSON(data []byte) error {
	var tmp LayoutKind
	type plain LayoutKind
	if err := json.Unmarshal(data, (*plain)(&tmp)); err != nil {
		return err
	}
	if err := (&tmp).validate(); err != nil {
		return err
	}
	*k = tmp
	return nil
}

func (k *LayoutKind) UnmarshalYAML(unmarshal func(any) error) error {
	var tmp LayoutKind
	type plain LayoutKind
	if err := unmarshal((*plain)(&tmp)); err != nil {
		return err
	}
	if err := (&tmp).validate(); err != nil {
		return err
	}
	*k = tmp
	return nil
}

func (k *LayoutKind) validate() error {
	if len(*k) == 0 {
		return fmt.Errorf("layout.kind cannot be empty")
	}
	if _, ok := layoutKindMap[*k]; !ok {
		return fmt.Errorf("unknown layout.kind %q used", *k)
	}
	return nil
}

type RepeatVariableAlignment string

const (
	RepeatVariableAlignmentHorizontal RepeatVariableAlignment = "horizontal"
	RepeatVariableAlignmentVertical   RepeatVariableAlignment = "vertical"
)

var repeatVariableAlignmentMap = map[RepeatVariableAlignment]bool{
	RepeatVariableAlignmentHorizontal: true,
	RepeatVariableAlignmentVertical:   true,
}

func (a *RepeatVariableAlignment) UnmarshalJSON(data []byte) error {
	var tmp RepeatVariableAlignment
	type plain RepeatVariableAlignment
	if err := json.Unmarshal(data, (*plain)(&tmp)); err != nil {
		return err
	}
	if err := (&tmp).validate(); err != nil {
		return err
	}
	*a = tmp
	return nil
}

func (a *RepeatVariableAlignment) UnmarshalYAML(unmarshal func(any) error) error {
	var tmp RepeatVariableAlignment
	type plain RepeatVariableAlignment
	if err := unmarshal((*plain)(&tmp)); err != nil {
		return err
	}
	if err := (&tmp).validate(); err != nil {
		return err
	}
	*a = tmp
	return nil
}

func (a *RepeatVariableAlignment) validate() error {
	if _, ok := repeatVariableAlignmentMap[*a]; !ok {
		return fmt.Errorf("unknown repeatVariable.alignment %q used", *a)
	}
	return nil
}

type RepeatVariable struct {
	Value     string                  `json:"value" yaml:"value"`
	MaxPer    *int                    `json:"maxPer,omitempty" yaml:"maxPer,omitempty"`
	Alignment RepeatVariableAlignment `json:"alignment,omitempty" yaml:"alignment,omitempty"`
}

type GridItem struct {
	X              int             `json:"x" yaml:"x"`
	Y              int             `json:"y" yaml:"y"`
	Width          int             `json:"width" yaml:"width"`
	Height         int             `json:"height" yaml:"height"`
	Content        *common.JSONRef `json:"content" yaml:"content"`
	RepeatVariable *RepeatVariable `json:"repeatVariable,omitempty" yaml:"repeatVariable,omitempty"`
}

type GridLayoutCollapse struct {
	Open bool `json:"open" yaml:"open"`
}

type GridLayoutDisplay struct {
	Title string `json:"title" yaml:"title"`
	// If Collapse is defined, the grid layout will be rendered in a collapsible group.
	// If not defined, the grid layout will be rendered expanded without the ability to collapse it.
	Collapse *GridLayoutCollapse `json:"collapse,omitempty" yaml:"collapse,omitempty"`
}

type GridLayoutSpec struct {
	Display        *GridLayoutDisplay `json:"display,omitempty" yaml:"display,omitempty"`
	Items          []GridItem         `json:"items" yaml:"items"`
	RepeatVariable string             `json:"repeatVariable,omitempty" yaml:"repeatVariable,omitempty"`
}

type TabItem struct {
	// Display name of the tab shown in the tab bar.
	Name string `json:"name" yaml:"name"`
	// The grid layout items within this tab.
	Items []GridItem `json:"items" yaml:"items"`
}

type TabLayoutDisplay struct {
	Title string `json:"title" yaml:"title"`
	// If Collapse is defined, the tab group will be rendered in a collapsible container.
	// If not defined, the tab group will be rendered expanded without the ability to collapse it.
	Collapse *GridLayoutCollapse `json:"collapse,omitempty" yaml:"collapse,omitempty"`
}

type TabLayoutSpec struct {
	Display *TabLayoutDisplay `json:"display,omitempty" yaml:"display,omitempty"`
	// Ordered list of tabs. The first tab (or defaultTab index) is shown by default.
	Tabs []TabItem `json:"tabs" yaml:"tabs"`
	// Zero-based index of the tab to show by default. Defaults to 0.
	DefaultTab int `json:"defaultTab,omitempty" yaml:"defaultTab,omitempty"`
}

type LayoutSpec any

type tmpDashboardLayout struct {
	Kind LayoutKind     `json:"kind" yaml:"kind"`
	Spec map[string]any `json:"spec" yaml:"spec"`
}

type Layout struct {
	Kind LayoutKind `json:"kind" yaml:"kind"`
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:Schemaless
	Spec LayoutSpec `json:"spec" yaml:"spec"`
}

func (d *Layout) UnmarshalJSON(data []byte) error {
	jsonUnmarshalFunc := func(panel any) error {
		return json.Unmarshal(data, panel)
	}
	return d.unmarshal(jsonUnmarshalFunc, json.Marshal, json.Unmarshal)
}

func (d *Layout) UnmarshalYAML(unmarshal func(any) error) error {
	return d.unmarshal(unmarshal, yaml.Marshal, yaml.Unmarshal)
}

func (d *Layout) unmarshal(unmarshal func(any) error, staticMarshal func(any) ([]byte, error), staticUnmarshal func([]byte, any) error) error {
	var tmpLayout tmpDashboardLayout
	if err := unmarshal(&tmpLayout); err != nil {
		return err
	}
	d.Kind = tmpLayout.Kind

	if len(tmpLayout.Kind) == 0 {
		return fmt.Errorf("variable.kind cannot be empty")
	}

	rawParameter, err := staticMarshal(tmpLayout.Spec)
	if err != nil {
		return err
	}
	var spec any
	switch tmpLayout.Kind {
	case KindGridLayout:
		spec = &GridLayoutSpec{}
	case KindTabLayout:
		spec = &TabLayoutSpec{}
	}
	if err := staticUnmarshal(rawParameter, spec); err != nil {
		return err
	}
	d.Spec = spec
	return nil
}

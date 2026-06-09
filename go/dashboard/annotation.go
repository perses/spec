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
	"errors"

	"github.com/perses/spec/go/plugin"
)

type AnnotationDisplay struct {
	Name        string `json:"name" yaml:"name"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	Hidden      bool   `json:"hidden,omitempty" yaml:"hidden,omitempty"`
	Color       string `json:"color,omitempty" yaml:"color,omitempty"`
}

type AnnotationSpec struct {
	Display AnnotationDisplay `json:"display" yaml:"display"`
	Plugin  plugin.Plugin     `json:"plugin" yaml:"plugin"`
}

func (d *AnnotationSpec) UnmarshalJSON(data []byte) error {
	var tmp AnnotationSpec
	type plain AnnotationSpec
	if err := json.Unmarshal(data, (*plain)(&tmp)); err != nil {
		return err
	}
	if err := (&tmp).validate(); err != nil {
		return err
	}
	*d = tmp
	return nil
}

func (d *AnnotationSpec) UnmarshalYAML(unmarshal func(any) error) error {
	var tmp AnnotationSpec
	type plain AnnotationSpec
	if err := unmarshal((*plain)(&tmp)); err != nil {
		return err
	}
	if err := (&tmp).validate(); err != nil {
		return err
	}
	*d = tmp
	return nil
}

func (d *AnnotationSpec) validate() error {
	if len(d.Display.Name) == 0 {
		return errors.New("annotation name cannot be empty")
	}
	return nil
}

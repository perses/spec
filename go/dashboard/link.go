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
	"strings"
)

// Link is a URL or a same-project tag menu. Exactly one of URL or Tags is set.
type Link struct {
	Name             string   `json:"name,omitempty" yaml:"name,omitempty"`
	URL              string   `json:"url,omitempty" yaml:"url,omitempty"`
	Tooltip          string   `json:"tooltip,omitempty" yaml:"tooltip,omitempty"`
	RenderVariables  bool     `json:"renderVariables,omitempty" yaml:"renderVariables,omitempty"`
	TargetBlank      bool     `json:"targetBlank,omitempty" yaml:"targetBlank,omitempty"`
	Tags             []string `json:"tags,omitempty" yaml:"tags,omitempty"`
	KeepTime         bool     `json:"keepTime,omitempty" yaml:"keepTime,omitempty"`
	IncludeVariables bool     `json:"includeVariables,omitempty" yaml:"includeVariables,omitempty"`
}

func (l *Link) UnmarshalJSON(data []byte) error {
	var tmp Link
	type plain Link
	if err := json.Unmarshal(data, (*plain)(&tmp)); err != nil {
		return err
	}
	if err := (&tmp).validate(); err != nil {
		return err
	}
	*l = tmp
	return nil
}

func (l *Link) UnmarshalYAML(unmarshal func(any) error) error {
	var tmp Link
	type plain Link
	if err := unmarshal((*plain)(&tmp)); err != nil {
		return err
	}
	if err := (&tmp).validate(); err != nil {
		return err
	}
	*l = tmp
	return nil
}

func (l *Link) validate() error {
	l.URL = strings.TrimSpace(l.URL)
	l.Name = strings.TrimSpace(l.Name)

	var tags []string
	for _, tag := range l.Tags {
		tag = strings.TrimSpace(tag)
		if tag == "" {
			return errors.New("link tag must not be empty")
		}
		tags = append(tags, tag)
	}
	l.Tags = tags

	hasURL := l.URL != ""
	hasTags := len(l.Tags) > 0
	if hasURL == hasTags {
		if hasURL {
			return errors.New("link must set either a url or tags, not both")
		}
		return errors.New("link must set a url or at least one tag")
	}
	if hasTags && l.Name == "" {
		return errors.New("a tag link must set a name")
	}
	return nil
}

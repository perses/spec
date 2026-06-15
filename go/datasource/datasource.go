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

package datasource

import (
	"encoding/json"
	"fmt"

	"github.com/perses/spec/go/common"
	"github.com/perses/spec/go/datasource/proxy/http"
	"github.com/perses/spec/go/datasource/proxy/sql"
	"github.com/perses/spec/go/plugin"
)

type Spec struct {
	Display *common.Display `json:"display,omitempty" yaml:"display,omitempty"`
	Default bool            `json:"default" yaml:"default"`
	// Plugin will contain the datasource configuration.
	// The data typed is available in Cue.
	Plugin plugin.Plugin `json:"plugin" yaml:"plugin"`
}

// HTTPDatasourceSpec is the struct that can be used to define an HTTP Datasource plugin.
// To be used when implementing a plugin, and you want to provide the associated go-sdk.
// This struct is just here to avoid developer to redefine the same struct in their own plugin implementation,
// because most of the time developers do not need more field than the two proposed.
// If you need more, define your own struct and use it in your plugin implementation.
type HTTPDatasourceSpec struct {
	DirectURL string      `json:"directUrl,omitempty" yaml:"directUrl,omitempty"`
	Proxy     *http.Proxy `json:"proxy,omitempty" yaml:"proxy,omitempty"`
}

func (s *HTTPDatasourceSpec) UnmarshalJSON(data []byte) error {
	type plain HTTPDatasourceSpec
	var tmp HTTPDatasourceSpec
	if err := json.Unmarshal(data, (*plain)(&tmp)); err != nil {
		return err
	}
	if err := (&tmp).validate(); err != nil {
		return err
	}
	*s = tmp
	return nil
}

func (s *HTTPDatasourceSpec) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var tmp HTTPDatasourceSpec
	type plain HTTPDatasourceSpec
	if err := unmarshal((*plain)(&tmp)); err != nil {
		return err
	}
	if err := (&tmp).validate(); err != nil {
		return err
	}
	*s = tmp
	return nil
}

func (s *HTTPDatasourceSpec) validate() error {
	if len(s.DirectURL) == 0 && s.Proxy == nil {
		return fmt.Errorf("directUrl or proxy cannot be empty")
	}
	if len(s.DirectURL) > 0 && s.Proxy != nil {
		return fmt.Errorf("at most directUrl or proxy must be configured")
	}
	return nil
}

// SQLDatasourceSpec is the struct that can be used to define an SQL Datasource plugin.
// To be used when implementing a plugin, and you want to provide the associated go-sdk.
// This struct is just here to avoid developer to redefine the same struct in their own plugin implementation,
// because most of the time developers do not need more field than the two proposed.
// If you need more, define your own struct and use it in your plugin implementation.
type SQLDatasourceSpec struct {
	Proxy *sql.Proxy
}

func (s *SQLDatasourceSpec) UnmarshalJSON(data []byte) error {
	type plain SQLDatasourceSpec
	var tmp SQLDatasourceSpec
	if err := json.Unmarshal(data, (*plain)(&tmp)); err != nil {
		return err
	}
	if err := (&tmp).validate(); err != nil {
		return err
	}
	*s = tmp
	return nil
}

func (s *SQLDatasourceSpec) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var tmp SQLDatasourceSpec
	type plain SQLDatasourceSpec
	if err := unmarshal((*plain)(&tmp)); err != nil {
		return err
	}
	if err := (&tmp).validate(); err != nil {
		return err
	}
	*s = tmp
	return nil
}

func (s *SQLDatasourceSpec) validate() error {
	if s.Proxy == nil {
		return fmt.Errorf("proxy cannot be empty")
	}
	return nil
}

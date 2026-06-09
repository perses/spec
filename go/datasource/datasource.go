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
	"github.com/perses/spec/go/common"
	"github.com/perses/spec/go/plugin"
)

type Spec struct {
	Display *common.Display `json:"display,omitempty" yaml:"display,omitempty"`
	Default bool            `json:"default" yaml:"default"`
	// Plugin will contain the datasource configuration.
	// The data typed is available in Cue.
	Plugin plugin.Plugin `json:"plugin" yaml:"plugin"`
}

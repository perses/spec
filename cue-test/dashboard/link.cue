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

legacyLink: #Link & {name: "todo"}
emptyURLLink: #Link & {name: "todo", url: ""}
nbspURLLink: #Link & {url: "\u00a0"}
urlLink: #Link & {
	name:             "Docs"
	url:              "https://perses.dev"
	keepTime:         true
	includeVariables: true
	renderVariables:  true
	targetBlank:      true
}
tagLink: #Link & {
	name:             "Traffic"
	tags: ["traffic"]
	keepTime:         true
	includeVariables: true
	targetBlank:      true
}
nullTagsLink: #Link & {name: "Docs", url: "https://perses.dev", tags: null}

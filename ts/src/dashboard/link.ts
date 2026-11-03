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

interface LinkBase {
  tooltip?: string;
  /** Copy the current time range. Applies to a URL and to a tag menu. */
  keepTime?: boolean;
  /** Copy variables. A URL receives every current variable. A tag menu receives those the destination declares. */
  includeVariables?: boolean;
}

/** url is present, possibly empty. tags is absent. An empty URL is a placeholder stored before tag links. */
export interface URLLink extends LinkBase {
  name?: string;
  url: string;
  /** Interpolate $var inside the URL. Tag menus do not use this. */
  renderVariables?: boolean;
  targetBlank?: boolean;
  tags?: never;
}

/** Same-project menu of dashboards that carry every tag. url is absent. */
export interface TagLink extends LinkBase {
  name: string;
  tags: string[];
  /** Open the chosen dashboard in a new tab. */
  targetBlank?: boolean;
  url?: never;
  /** Tag menus have no URL to interpolate. */
  renderVariables?: never;
}

export type Link = URLLink | TagLink;

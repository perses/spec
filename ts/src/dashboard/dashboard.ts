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

import { Display, DurationString } from '../common';
import { DatasourceSpec } from '../datasource';
import { LayoutDefinition } from './layout';
import { VariableDefinition } from './variable';
import { Link } from './link';
import { PanelDefinition } from './panel';
import { AnnotationSpec } from './annotation';

export interface DashboardSpec {
  display?: Display;
  datasources?: Record<string, DatasourceSpec>;
  annotations?: AnnotationSpec[];
  duration: DurationString;
  refreshInterval?: DurationString;
  variables: VariableDefinition[];
  layouts: LayoutDefinition[];
  panels: Record<string, PanelDefinition>;
  timezone?: string;
  links?: Link[];
}

export interface DashboardSelector {
  project: string;
  dashboard: string;
  tags?: string[];
}

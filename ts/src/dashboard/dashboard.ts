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

import type { Display, DurationString } from '../common';
import type { DatasourceSpec } from '../datasource';
import type { AnnotationSpec } from './annotation';
import type { LayoutDefinition } from './layout';
import type { Link } from './link';
import type { PanelDefinition } from './panel';
import type { VariableDefinition } from './variable';

/** Optional PromQL query packing for multi-series panels. */
export type QueryBatchingMode = 'off' | 'panel' | 'viewport' | 'dashboard';

export interface QueryBatchingSpec {
  /**
   * off = disabled;
   * panel = batch queries within one panel;
   * viewport = batch across visible panels;
   * dashboard = batch across the whole dashboard (including non-visible panels when loaded).
   */
  mode?: QueryBatchingMode;
  /** Optional max queries per batch HTTP request (default runtime-defined). */
  maxPerRequest?: number;
}

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
  queryBatching?: QueryBatchingSpec;
}

export interface DashboardSelector {
  project: string;
  dashboard: string;
  tags?: string[];
}

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

import { Definition, UnknownSpec } from '../../common';
import { TimeSeriesData } from './time-series-data';
import { TraceData } from './trace-data';
import { LogData } from './log-data';
import { ProfileData } from './profile-data';
import { AlertsData } from './alerts-data';
import { SilencesData } from './silences-data';

interface QuerySpec<PluginSpec> {
  name?: string;
  plugin: Definition<PluginSpec>;
}
/**
 * A generic query definition interface that can be extended to support more than just TimeSeriesQuery
 */
// Kind needs to be `any` because otherwise typescript will complain 'unknown' is not assignable to type '"TimeSeriesQuery"' in a few places
// eslint-disable-next-line @typescript-eslint/no-explicit-any
export interface QueryDefinition<Kind = any, PluginSpec = UnknownSpec> {
  kind: Kind;
  spec: QuerySpec<PluginSpec>;
}

/**
 * Mapping the query plugin to the data type it returns
 */
export interface QueryType {
  TimeSeriesQuery: TimeSeriesData;
  TraceQuery: TraceData;
  ProfileQuery: ProfileData;
  LogQuery: LogData;
  AlertsQuery: AlertsData;
  SilencesQuery: SilencesData;
}

/**
 * Check if the given type is a valid {@link QueryPluginType} with compile time safety
 * @param type
 */
export function isValidQueryPluginType(type: string): type is QueryPluginType {
  return ['TimeSeriesQuery', 'TraceQuery', 'ProfileQuery', 'LogQuery', 'AlertsQuery', 'SilencesQuery'].includes(
    type as QueryPluginType
  );
}

/**
 * Extract the keys of QueryType
 * ex: 'TimeSeriesQuery'
 */
export type QueryPluginType = keyof QueryType;

/**
 * Values of QueryType
 * ex: 'TimeSeriesData'
 */
export type QueryDataType = QueryType[keyof QueryType];

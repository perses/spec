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

import type { AbsoluteTimeRange, UnixTimeMs } from '../../common';
import type { BaseMetadata } from './base-metadata';
import type { Labels, TimeSeriesValueTuple, TimeSeriesHistogramTuple } from './time-series-queries';

export interface TimeScale {
  startMs: number;
  endMs: number;
  stepMs: number;
  rangeMs: number;
}

export interface TimeSeriesData {
  timeRange?: AbsoluteTimeRange;
  stepMs?: number;
  series: TimeSeries[];
  /**
   * Exemplar data attached to the series of this query result.
   * Mirrors the Prometheus `query_exemplars` response shape. Optional and
   * backward-compatible: query plugins that don't support exemplars can omit it.
   */
  exemplars?: TimeSeriesExemplars[];
  metadata?: TimeSeriesMetadata;
}

export interface TimeSeries {
  name: string;
  values: TimeSeriesValueTuple[];
  histograms?: TimeSeriesHistogramTuple[];
  formattedName?: string;
  labels?: Labels;
}

export interface TimeSeriesMetadata extends BaseMetadata {
  [key: string]: unknown;
}

/**
 * A single exemplar: a specific data point with its own labels (e.g. a trace ID)
 * referencing a particular series at a particular timestamp.
 */
export interface Exemplar {
  labels: Labels;
  value: number;
  timestamp: UnixTimeMs;
}

/**
 * The exemplars belonging to one series, keyed by the series labels.
 * Mirrors an entry of the Prometheus `query_exemplars` response:
 * https://prometheus.io/docs/prometheus/latest/querying/api/#querying-exemplars
 */
export interface TimeSeriesExemplars {
  seriesLabels: Labels;
  exemplars: Exemplar[];
}

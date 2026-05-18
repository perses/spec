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

import { BaseMetadata } from './base-metadata';

export type SilenceState = 'active' | 'expired' | 'pending';

export interface SilenceMatcher {
  /** Label or field name to match against */
  name: string;
  /** Value to match */
  value: string;
  /** Whether this is an equality match (true) or negation (false). Defaults to true. */
  isEqual?: boolean;
  /** Whether the value is a regex pattern */
  isRegex?: boolean;
}

export interface Silence {
  /** Unique identifier */
  id: string;
  /** Current state of the silence */
  state: SilenceState;
  /** Matching criteria that determine which alerts this silence applies to */
  matchers: SilenceMatcher[];
  /** ISO 8601 timestamp when the silence becomes active */
  startsAt: string;
  /** ISO 8601 timestamp when the silence expires */
  endsAt: string;
  /** User or system that created the silence */
  createdBy: string;
  /** Human-readable reason for the silence */
  comment?: string;
  /** ISO 8601 timestamp of the last update */
  updatedAt?: string;
}

export interface SilencesData {
  silences: Silence[];
  metadata?: SilencesMetadata;
}

export interface SilencesMetadata extends BaseMetadata {
  [key: string]: unknown;
}

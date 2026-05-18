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

export type AlertState = 'inactive' | 'pending' | 'firing' | 'resolved';

export interface Alert {
  /** Unique identifier for this alert instance */
  id: string;
  /** Alert name or title */
  name: string;
  /** Current alert state in the generic lifecycle */
  state: AlertState;
  /** Key-value labels identifying the alert */
  labels: Record<string, string>;
  /** Key-value annotations with additional context */
  annotations: Record<string, string>;
  /** Normalized severity level */
  severity?: string;
  /** ISO 8601 timestamp when the alert started firing */
  startsAt: string;
  /** ISO 8601 timestamp when the alert was resolved or is expected to expire */
  endsAt?: string;
  /** ISO 8601 timestamp of the last update to this alert */
  updatedAt?: string;
  /** URL linking back to the alert source */
  sourceURL?: string;
  /** Whether this alert is suppressed */
  suppressed?: boolean;
  /** Typed references to the suppression sources */
  suppressedBy?: SuppressionRule[];
  /** Whether a responder has acknowledged this alert */
  acknowledged?: boolean;
  /** Notification targets this alert is routed to */
  receivers?: string[];
}

export interface SuppressionRule {
  /** Suppression category — defined by each provider plugin */
  type: string;
  /** Identifier of the suppression source */
  id: string;
}

export interface AlertsData {
  alerts: Alert[];
  metadata?: AlertsMetadata;
}

export interface AlertsMetadata extends BaseMetadata {
  [key: string]: unknown;
}

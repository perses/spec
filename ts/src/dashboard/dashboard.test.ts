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

import { describe, expect, it } from 'vitest';

import { DashboardSelector, DashboardSpec } from './dashboard';

// A realistic JSON payload exercising every field of DashboardSpec, mirroring the style of
// the Go unmarshalling tests (see go/dashboard/panel_test.go): a JSON string is parsed, and the
// resulting object is checked against the expected shape/values field by field.
const dashboardSpecJSON = `{
  "display": { "name": "My Dashboard", "description": "A test dashboard" },
  "datasources": {
    "promDemo": {
      "default": true,
      "plugin": { "kind": "PrometheusDatasource", "spec": { "directUrl": "https://prometheus.demo.do.prometheus.io" } }
    }
  },
  "annotations": [
    { "display": { "name": "Deploys" }, "plugin": { "kind": "PrometheusPromQLAnnotation", "spec": {} } }
  ],
  "duration": "1h",
  "refreshInterval": "30s",
  "variables": [
    {
      "kind": "TextVariable",
      "spec": { "name": "environment", "value": "production", "constant": true }
    },
    {
      "kind": "ListVariable",
      "spec": {
        "name": "instance",
        "display": { "name": "Instance", "hidden": false },
        "allowAllValue": true,
        "allowMultiple": true,
        "defaultValue": "$__all",
        "plugin": { "kind": "PrometheusPromQLVariable", "spec": { "expr": "up" } }
      }
    }
  ],
  "layouts": [
    {
      "kind": "Grid",
      "spec": {
        "display": { "title": "Overview" },
        "items": [
          {
            "x": 0,
            "y": 0,
            "width": 12,
            "height": 6,
            "content": { "$ref": "#/spec/panels/cpuUsage" },
            "repeatVariable": { "value": "instance", "maxPer": 4, "alignment": "horizontal" }
          }
        ]
      }
    }
  ],
  "panels": {
    "cpuUsage": {
      "kind": "Panel",
      "spec": {
        "display": { "name": "CPU Usage" },
        "plugin": { "kind": "TimeSeriesChart", "spec": {} },
        "queries": [
          {
            "kind": "TimeSeriesQuery",
            "spec": { "plugin": { "kind": "PrometheusTimeSeriesQuery", "spec": { "query": "rate(cpu[5m])" } } }
          }
        ],
        "links": [{ "name": "Runbook", "url": "https://runbook.example.com" }]
      }
    }
  },
  "timezone": "America/New_York",
  "links": [{ "name": "Docs", "url": "https://perses.dev", "renderVariables": true }]
}`;

// The same payload, hand-typed as a DashboardSpec literal. Because this is a plain object
// literal (not a JSON.parse result typed via a cast), the TypeScript compiler performs full
// structural/excess-property checking against DashboardSpec, so any interface field that no
// longer matches the JSON representation above will fail to compile.
const expectedDashboardSpec: DashboardSpec = {
  display: { name: 'My Dashboard', description: 'A test dashboard' },
  datasources: {
    promDemo: {
      default: true,
      plugin: { kind: 'PrometheusDatasource', spec: { directUrl: 'https://prometheus.demo.do.prometheus.io' } },
    },
  },
  annotations: [{ display: { name: 'Deploys' }, plugin: { kind: 'PrometheusPromQLAnnotation', spec: {} } }],
  duration: '1h',
  refreshInterval: '30s',
  variables: [
    {
      kind: 'TextVariable',
      spec: { name: 'environment', value: 'production', constant: true },
    },
    {
      kind: 'ListVariable',
      spec: {
        name: 'instance',
        display: { name: 'Instance', hidden: false },
        allowAllValue: true,
        allowMultiple: true,
        defaultValue: '$__all',
        plugin: { kind: 'PrometheusPromQLVariable', spec: { expr: 'up' } },
      },
    },
  ],
  layouts: [
    {
      kind: 'Grid',
      spec: {
        display: { title: 'Overview' },
        items: [
          {
            x: 0,
            y: 0,
            width: 12,
            height: 6,
            content: { $ref: '#/spec/panels/cpuUsage' },
            repeatVariable: { value: 'instance', maxPer: 4, alignment: 'horizontal' },
          },
        ],
      },
    },
  ],
  panels: {
    cpuUsage: {
      kind: 'Panel',
      spec: {
        display: { name: 'CPU Usage' },
        plugin: { kind: 'TimeSeriesChart', spec: {} },
        queries: [
          {
            kind: 'TimeSeriesQuery',
            spec: { plugin: { kind: 'PrometheusTimeSeriesQuery', spec: { query: 'rate(cpu[5m])' } } },
          },
        ],
        links: [{ name: 'Runbook', url: 'https://runbook.example.com' }],
      },
    },
  },
  timezone: 'America/New_York',
  links: [{ name: 'Docs', url: 'https://perses.dev', renderVariables: true }],
};

describe('DashboardSpec JSON representation', () => {
  it('unmarshals a full JSON payload into a matching DashboardSpec', () => {
    const spec = JSON.parse(dashboardSpecJSON) as DashboardSpec;
    expect(spec).toEqual(expectedDashboardSpec);
  });

  it('preserves every top-level field name and nested value', () => {
    const spec = JSON.parse(dashboardSpecJSON) as DashboardSpec;

    expect(spec.display).toEqual({ name: 'My Dashboard', description: 'A test dashboard' });
    expect(spec.datasources?.promDemo?.default).toBe(true);
    expect(spec.annotations?.[0]?.display.name).toBe('Deploys');
    expect(spec.duration).toBe('1h');
    expect(spec.refreshInterval).toBe('30s');
    expect(spec.variables).toHaveLength(2);
    expect(spec.variables[0]).toMatchObject({ kind: 'TextVariable', spec: { name: 'environment' } });
    expect(spec.variables[1]).toMatchObject({ kind: 'ListVariable', spec: { name: 'instance' } });
    expect(spec.layouts).toHaveLength(1);
    expect(spec.layouts[0]).toMatchObject({ kind: 'Grid' });
    expect(spec.panels.cpuUsage?.spec.display?.name).toBe('CPU Usage');
    expect(spec.timezone).toBe('America/New_York');
    expect(spec.links?.[0]).toEqual({ name: 'Docs', url: 'https://perses.dev', renderVariables: true });
  });

  it('round-trips through JSON without losing or renaming any fields', () => {
    const spec = JSON.parse(dashboardSpecJSON) as DashboardSpec;

    // Mirrors the Go marshal -> unmarshal round trip tests (e.g. TestPanelSpecAnnotationsRoundTripJSON):
    // serializing the parsed spec back to JSON and re-parsing it must reproduce the original payload.
    const roundTripped = JSON.parse(JSON.stringify(spec));
    expect(roundTripped).toEqual(JSON.parse(dashboardSpecJSON));
  });

  it('parses a minimal payload containing only the required fields', () => {
    const minimalJSON = `{
      "duration": "1h",
      "variables": [],
      "layouts": [],
      "panels": {}
    }`;

    const spec = JSON.parse(minimalJSON) as DashboardSpec;
    expect(spec).toEqual({ duration: '1h', variables: [], layouts: [], panels: {} });
    expect(spec.display).toBeUndefined();
    expect(spec.datasources).toBeUndefined();
    expect(spec.annotations).toBeUndefined();
    expect(spec.refreshInterval).toBeUndefined();
    expect(spec.timezone).toBeUndefined();
    expect(spec.links).toBeUndefined();
  });
});

describe('DashboardSelector JSON representation', () => {
  it('unmarshals a JSON payload including optional tags', () => {
    const selectorJSON = `{ "project": "observability", "dashboard": "cpu-usage", "tags": ["prod", "team-a"] }`;

    const expectedSelector: DashboardSelector = {
      project: 'observability',
      dashboard: 'cpu-usage',
      tags: ['prod', 'team-a'],
    };

    const selector = JSON.parse(selectorJSON) as DashboardSelector;
    expect(selector).toEqual(expectedSelector);
  });

  it('unmarshals a JSON payload omitting the optional tags field', () => {
    const selectorJSON = `{ "project": "observability", "dashboard": "cpu-usage" }`;

    const selector = JSON.parse(selectorJSON) as DashboardSelector;
    expect(selector).toEqual({ project: 'observability', dashboard: 'cpu-usage' });
    expect(selector.tags).toBeUndefined();
  });
});

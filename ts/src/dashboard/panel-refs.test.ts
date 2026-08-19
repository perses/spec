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

import { DashboardSpec } from './dashboard';
import { PanelDefinition, PanelRef } from './panel';
import { createPanelRef, getPanelKeyFromRef, resolvePanelRef } from './panel-refs';

function buildPanelDefinition(name: string): PanelDefinition {
  return {
    kind: 'Panel',
    spec: {
      display: { name },
      plugin: { kind: 'TimeSeriesChart', spec: {} },
    },
  };
}

function buildDashboardSpec(panels: Record<string, PanelDefinition>): DashboardSpec {
  return {
    duration: '1h',
    variables: [],
    layouts: [],
    panels,
  };
}

describe('createPanelRef', () => {
  it('creates a PanelRef using the standard panels prefix', () => {
    expect(createPanelRef('cpuUsage')).toEqual({ $ref: '#/spec/panels/cpuUsage' });
  });

  it('creates a PanelRef for an empty panel key', () => {
    expect(createPanelRef('')).toEqual({ $ref: '#/spec/panels/' });
  });
});

describe('getPanelKeyFromRef', () => {
  it('extracts the panel key from a PanelRef', () => {
    const panelRef: PanelRef = { $ref: '#/spec/panels/cpuUsage' };
    expect(getPanelKeyFromRef(panelRef)).toBe('cpuUsage');
  });

  it('round-trips with createPanelRef', () => {
    const panelRef = createPanelRef('memoryUsage');
    expect(getPanelKeyFromRef(panelRef)).toBe('memoryUsage');
  });

  it('returns an empty string when the ref has no key', () => {
    const panelRef: PanelRef = { $ref: '#/spec/panels/' };
    expect(getPanelKeyFromRef(panelRef)).toBe('');
  });
});

describe('resolvePanelRef', () => {
  it('resolves a PanelRef to its PanelDefinition', () => {
    const cpuPanel = buildPanelDefinition('CPU Usage');
    const spec = buildDashboardSpec({ cpuUsage: cpuPanel });

    expect(resolvePanelRef(spec, createPanelRef('cpuUsage'))).toBe(cpuPanel);
  });

  it('throws when the PanelRef does not match any panel in the spec', () => {
    const spec = buildDashboardSpec({ cpuUsage: buildPanelDefinition('CPU Usage') });
    const panelRef = createPanelRef('missingPanel');

    expect(() => resolvePanelRef(spec, panelRef)).toThrow(
      'Could not resolve panels reference #/spec/panels/missingPanel',
    );
  });

  it('throws when the spec has no panels at all', () => {
    const spec = buildDashboardSpec({});
    const panelRef = createPanelRef('cpuUsage');

    expect(() => resolvePanelRef(spec, panelRef)).toThrow('Could not resolve panels reference #/spec/panels/cpuUsage');
  });
});

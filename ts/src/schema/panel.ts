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

import { z } from 'zod';

import type { Link, PanelDefinition, PanelDisplay, PanelSpec, QueryDefinition } from '../dashboard';
import { annotationSpecSchema } from './annotation';
import { pluginSchema } from './plugin';
import type { PluginSchema } from './plugin';

export const panelDisplaySpec: z.ZodType<PanelDisplay, PanelDisplay> = z.object({
  name: z.string().optional(),
  description: z.string().optional(),
  hideHeader: z.boolean().optional(),
});

export const querySpecSchema: z.ZodType<QueryDefinition, QueryDefinition> = z.object({
  kind: z.string().min(1),
  spec: z.object({
    name: z.string().optional(),
    plugin: pluginSchema,
  }),
});

export const linkSchema: z.ZodType<Link, Link> = z.object({
  name: z.string().optional(),
  url: z.string().min(1),
  tooltip: z.string().optional(),
  renderVariables: z.boolean().optional(),
  targetBlank: z.boolean().optional(),
});

export const panelSpecSchema: z.ZodType<PanelSpec, PanelSpec> = z.object({
  display: panelDisplaySpec.optional(),
  plugin: pluginSchema,
  queries: z.array(querySpecSchema).optional(),
  links: z.array(linkSchema).optional(),
  annotations: z.array(annotationSpecSchema).optional(),
});

export function buildPanelSpecSchema(customPluginSchema: PluginSchema): z.ZodType<PanelSpec, PanelSpec> {
  return z.object({
    display: panelDisplaySpec.optional(),
    plugin: customPluginSchema,
    queries: z.array(querySpecSchema).optional(),
    links: z.array(linkSchema).optional(),
    annotations: z.array(annotationSpecSchema).optional(),
  });
}

export const panelDefinitionSchema: z.ZodType<PanelDefinition, PanelDefinition> = z.object({
  kind: z.literal('Panel'),
  spec: panelSpecSchema,
});

export function buildPanelDefinitionSchema(
  customPluginSchema: PluginSchema,
): z.ZodType<PanelDefinition, PanelDefinition> {
  return z.object({
    kind: z.literal('Panel'),
    spec: buildPanelSpecSchema(customPluginSchema),
  });
}

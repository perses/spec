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

import type {
  ListVariableDefinition,
  ListVariableSpec,
  TextVariableDefinition,
  TextVariableSpec,
  VariableDefinition,
  VariableDisplay,
} from '../dashboard';
import { pluginSchema } from './plugin';
import type { PluginSchema } from './plugin';

export const variableDisplaySchema: z.ZodType<VariableDisplay, VariableDisplay> = z.object({
  name: z.string().optional(),
  description: z.string().optional(),
  hidden: z.boolean().optional(),
});

export const variableListSpecSchema: z.ZodType<ListVariableSpec, ListVariableSpec> = z.object({
  name: z.string().min(1),
  display: variableDisplaySchema.optional(),
  defaultValue: z.string().or(z.array(z.string())).optional(),
  allowAllValue: z.boolean(),
  allowMultiple: z.boolean(),
  customAllValue: z.string().optional(),
  capturingRegexp: z.string().optional(),
  sort: z
    .enum([
      'none',
      'alphabetical-asc',
      'alphabetical-desc',
      'numerical-asc',
      'numerical-desc',
      'alphabetical-ci-asc',
      'alphabetical-ci-desc',
    ])
    .optional(),
  plugin: pluginSchema,
});

export function buildVariableListSpecSchema(
  customPluginSchema: PluginSchema,
): z.ZodType<ListVariableSpec, ListVariableSpec> {
  return z.object({
    name: z.string().min(1),
    display: variableDisplaySchema.optional(),
    defaultValue: z.string().or(z.array(z.string())).optional(),
    allowAllValue: z.boolean(),
    allowMultiple: z.boolean(),
    customAllValue: z.string().optional(),
    capturingRegexp: z.string().optional(),
    sort: z
      .enum([
        'none',
        'alphabetical-asc',
        'alphabetical-desc',
        'numerical-asc',
        'numerical-desc',
        'alphabetical-ci-asc',
        'alphabetical-ci-desc',
      ])
      .optional(),
    plugin: customPluginSchema,
  });
}

export const variableListSchema = z.object({
  kind: z.literal('ListVariable'),
  spec: variableListSpecSchema,
});

export function buildVariableListSchema(customPluginSchema: PluginSchema): typeof variableListSchema {
  return z.object({
    kind: z.literal('ListVariable'),
    spec: buildVariableListSpecSchema(customPluginSchema),
  });
}

export const variableTextSpecSchema: z.ZodType<TextVariableSpec, TextVariableSpec> = z.object({
  name: z.string().min(1),
  display: variableDisplaySchema.optional(),
  value: z.string(),
  constant: z.boolean().optional(),
});

export const variableTextSchema = z.object({
  kind: z.literal('TextVariable'),
  spec: variableTextSpecSchema,
});

export const variableSpecSchema: z.ZodType<
  TextVariableDefinition | ListVariableDefinition,
  TextVariableDefinition | ListVariableDefinition
> = z.discriminatedUnion('kind', [variableTextSchema, variableListSchema]);

export function buildVariableSpecSchema(
  customPluginSchema: PluginSchema,
): z.ZodType<VariableDefinition, VariableDefinition> {
  return z.union([variableTextSchema, buildVariableListSchema(customPluginSchema)]);
}

export const variableDefinitionSchema: z.ZodType<VariableDefinition, VariableDefinition> = variableSpecSchema;

export function buildVariableDefinitionSchema(
  customPluginSchema: PluginSchema,
): z.ZodType<VariableDefinition, VariableDefinition> {
  return z.discriminatedUnion('kind', [
    z.object({
      kind: z.literal('ListVariable'),
      spec: buildVariableListSpecSchema(customPluginSchema),
    }),
    z.object({
      kind: z.literal('TextVariable'),
      spec: variableTextSpecSchema,
    }),
  ]);
}

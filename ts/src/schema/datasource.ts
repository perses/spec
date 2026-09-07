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

import type { DatasourceSpec } from '../datasource';
import { displaySchema } from './display';
import { pluginSchema } from './plugin';
import type { PluginSchema } from './plugin';

export const datasourceSpecSchema: z.ZodType<DatasourceSpec, DatasourceSpec> = z.object({
  display: displaySchema.optional(),
  default: z.boolean(),
  plugin: pluginSchema,
});

export function buildDatasourceSpecSchema(customPluginSchema: PluginSchema): z.ZodType<DatasourceSpec, DatasourceSpec> {
  return z.object({
    display: displaySchema.optional(),
    default: z.boolean(),
    plugin: customPluginSchema,
  });
}

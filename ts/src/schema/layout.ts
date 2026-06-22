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

export const repeatVariableSchema = z.object({
  value: z.string().min(1),
  maxPer: z.number().int('Provide valid number.').positive().optional(),
  alignment: z.enum(['horizontal', 'vertical']).optional(),
});

export const layoutDefinitionSchema = z.object({
  repeatVariable: repeatVariableSchema.optional(),
  width: z.number().int().positive(),
  height: z.number().int().positive(),
});

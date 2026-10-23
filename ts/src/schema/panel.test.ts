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

import { linkSchema } from './panel';

describe('linkSchema', () => {
  it.each([
    {
      name: 'url',
      input: { name: ' Docs ', url: ' https://perses.dev ', tooltip: 'docs', renderVariables: true, targetBlank: true },
      output: { name: 'Docs', url: 'https://perses.dev', tooltip: 'docs', renderVariables: true, targetBlank: true },
    },
    {
      name: 'tags',
      input: {
        name: ' Service detail ',
        url: '   ',
        tags: [' traffic ', 'billing'],
        keepTime: true,
        includeVariables: true,
      },
      output: { name: 'Service detail', tags: ['traffic', 'billing'], keepTime: true, includeVariables: true },
    },
  ])('normalizes a $name link', ({ input, output }) => {
    const parsed = linkSchema.parse(input);
    expect(parsed).toEqual(output);
    expect(parsed).not.toHaveProperty(output.url ? 'tags' : 'url');
  });

  it.each([
    [{ name: 'Docs' }, 'link must set a url or at least one tag'],
    [{ url: '   ' }, 'link must set a url or at least one tag'],
    [{ name: 'Docs', url: ' https://perses.dev ', tags: ['traffic'] }, 'link must set either a url or tags, not both'],
    [{ name: 'Docs', tags: ['traffic', ' '] }, 'link tag must not be empty'],
    [{ name: '  ', tags: ['traffic'] }, 'a tag link must set a name'],
  ])('rejects %j', (input, message) => {
    expect(() => linkSchema.parse(input)).toThrow(message);
  });
});

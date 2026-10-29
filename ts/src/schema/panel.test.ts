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
      input: {
        name: ' Docs ',
        url: ' https://perses.dev ',
        tooltip: 'docs',
        renderVariables: true,
        targetBlank: true,
        keepTime: true,
        includeVariables: true,
      },
      output: {
        name: 'Docs',
        url: 'https://perses.dev',
        tooltip: 'docs',
        renderVariables: true,
        targetBlank: true,
        keepTime: true,
        includeVariables: true,
      },
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
    if ('tags' in output) {
      expect(parsed).not.toHaveProperty('url');
      return;
    }
    expect(parsed).not.toHaveProperty('tags');
    expect(parsed).toHaveProperty('url');
  });

  it.each([
    { name: 'neither', input: { name: 'Docs' }, output: { name: 'Docs', url: '' } },
    { name: 'blank url', input: { name: 'todo', url: '' }, output: { name: 'todo', url: '' } },
    { name: 'whitespace url', input: { url: '   ' }, output: { url: '' } },
  ])('keeps a $name placeholder', ({ input, output }) => {
    const parsed = linkSchema.parse(input);
    expect(parsed).toEqual(output);
    expect(parsed).not.toHaveProperty('tags');
    expect(Object.values(parsed).every((value) => value !== undefined)).toBe(true);
  });

  it('omits undefined keys', () => {
    expect(Object.keys(linkSchema.parse({ url: 'https://perses.dev' }))).toEqual(['url']);
    expect(Object.keys(linkSchema.parse({ name: 'Traffic', tags: ['traffic'], targetBlank: true })).sort()).toEqual([
      'name',
      'tags',
      'targetBlank',
    ]);
  });

  it.each([
    [{ name: 'Docs', url: ' https://perses.dev ', tags: ['traffic'] }, 'link must set either a url or tags, not both'],
    [{ name: 'Docs', tags: ['traffic', ' '] }, 'link tag must not be empty'],
    [{ name: '  ', tags: ['traffic'] }, 'a tag link must set a name'],
    [{ name: 'Traffic', tags: ['traffic'], renderVariables: true }, 'renderVariables is only valid on a URL link'],
    [
      { name: 'Traffic', tags: Array.from({ length: 21 }, (_, i) => String(i + 1)) },
      'cannot contain more than 20 tags',
    ],
    [{ name: 'Traffic', tags: ['a'.repeat(51)] }, 'cannot contain more than 50 characters'],
    [{ name: 'Traffic', tags: ['é'.repeat(51)] }, 'cannot contain more than 50 characters'],
  ])('rejects %j', (input, message) => {
    expect(() => linkSchema.parse(input)).toThrow(message);
  });
});

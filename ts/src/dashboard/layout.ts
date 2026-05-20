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

import { PanelRef } from './panel';

export type LayoutDefinition = GridDefinition | TabDefinition;

export interface GridDefinition {
  kind: 'Grid';
  spec: {
    display?: {
      title: string;
      collapse?: {
        open: boolean;
      };
    };
    items: GridItemDefinition[];
    repeatVariable?: string;
  };
}

export interface RepeatVariable {
  value: string;
  maxPer?: number;
  alignment?: 'horizontal' | 'vertical';
}

export interface GridItemDefinition {
  x: number;
  y: number;
  width: number;
  height: number;
  content: PanelRef;
  repeatVariable?: RepeatVariable;
}

export interface TabDefinition {
  kind: 'Tabs';
  spec: {
    display?: {
      title: string;
      collapse?: {
        open: boolean;
      };
    };
    tabs: TabItemDefinition[];
    defaultTab?: number;
  };
}

export interface TabItemDefinition {
  name: string;
  items: GridItemDefinition[];
}

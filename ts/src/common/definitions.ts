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

/**
 * Optional metadata attached to a plugin definition. It mirrors the `metadata` field of the backend `Plugin` model
 * and is mainly used to pin a plugin to a specific version and/or registry.
 */
export interface PluginDefinitionMetadata {
  /**
   * Version of the plugin to use. When omitted, the latest version available in the Perses instance is used.
   */
  version?: string;
  /**
   * Registry the plugin comes from. When omitted, the default registry is used.
   */
  registry?: string;
}

/**
 * Base type for definitions in JSON config resources.
 */
export interface Definition<Spec> {
  kind: string;
  metadata?: PluginDefinitionMetadata;
  spec: Spec;
}

/**
 * Type to represent specs at runtime in framework code where we don't know the spec's real type, probably because it's
 * part of a plugin.
 */
export type UnknownSpec = Record<string, unknown>;

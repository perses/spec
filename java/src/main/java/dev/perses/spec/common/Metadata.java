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

package dev.perses.spec.common;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;

/**
 * Optional metadata attached to a plugin definition. Mirrors the Go {@code plugin.Metadata} model.
 */
@JsonInclude(JsonInclude.Include.NON_NULL)
public class Metadata {
    /**
     * Version is optional. If not provided, it means the latest version available in the Perses instance.
     * Version needs to follow the semantic versioning format (e.g., "1.0.0" or "v1.0.0").
     */
    @JsonProperty("version")
    public String version;

    /**
     * Registry is optional. If not provided, it means the default registry is: "perses.dev".
     */
    @JsonProperty("registry")
    public String registry;

    public Metadata() {
    }
}

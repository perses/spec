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

package dev.perses.dashboard;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import dev.perses.common.duration.Duration;
import dev.perses.dashboard.panel.Link;
import dev.perses.dashboard.panel.Panel;
import dev.perses.common.Display;

import java.util.List;
import java.util.Map;

@JsonInclude(JsonInclude.Include.NON_NULL)
public class Spec {
    @JsonProperty("display")
    public Display display;

    @JsonProperty("datasources")
    public Map<String, dev.perses.datasource.Spec> datasources;

    @JsonProperty("variables")
    public List<Variable> variables;

    @JsonProperty(value = "panels", required = true)
    public Map<String, Panel> panels;

    @JsonProperty(value = "layouts", required = true)
    public List<Layout> layouts;

    @JsonProperty(value = "duration", required = true)
    public Duration duration;

    @JsonProperty("refreshInterval")
    public Duration refreshInterval;

    @JsonProperty("timezone")
    public String timezone;

    @JsonProperty("links")
    public List<Link> Links;

    public Spec() {
    }
}

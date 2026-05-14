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

import com.fasterxml.jackson.databind.ObjectMapper;
import org.junit.jupiter.api.Test;

import java.io.InputStream;
import java.nio.charset.StandardCharsets;

import static org.junit.jupiter.api.Assertions.*;

public class SpecUnmarshalTest {
    private final ObjectMapper mapper = new ObjectMapper();

    private String readResource(String resourcePath) throws Exception {
        try (InputStream is = getClass().getResourceAsStream(resourcePath)) {
            assertNotNull(is, "resource not found: " + resourcePath);
            return new String(is.readAllBytes(), StandardCharsets.UTF_8);
        }
    }

    @Test
    public void testUnmarshalTabsLayout() throws Exception {
        String json = readResource("/dev/perses/dashboard/simple_tabs_dashboard.json");

        Spec d = mapper.readValue(json, Spec.class);
        assertNotNull(d);
        assertEquals("Tabs Dashboard", d.display.name);
        assertEquals(1, d.layouts.size());

        Layout layout = d.layouts.get(0);
        assertEquals(Layout.LayoutKind.Tabs, layout.kind);
        assertNotNull(layout.spec);
    }

    @Test
    public void testUnmarshalFullDashboard() throws Exception {
        String json = readResource("/dev/perses/dashboard/simple_spec_dashboard.json");

        Spec d = mapper.readValue(json, Spec.class);
        assertNotNull(d);
        assertNotNull(d.display);
        assertEquals("My Dashboard", d.display.name);
        assertEquals("A test dashboard", d.display.description);
        assertNotNull(d.datasources);
        assertTrue(d.datasources.containsKey("prom"));
        assertNotNull(d.variables);
        assertEquals(0, d.variables.size());
    }
}

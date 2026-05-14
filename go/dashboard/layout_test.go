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

package dashboard

import (
	"encoding/json"
	"testing"

	"github.com/perses/spec/go/common"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestUnmarshalJSONLayout(t *testing.T) {
	testSuite := []struct {
		title  string
		jason  string
		result Layout
	}{
		{
			title: "grid layout",
			jason: `
{
  "kind": "Grid",
  "spec": {
    "items": [
      {
        "x": 0,
        "y": 0,
        "width": 3,
        "height": 4,
        "content": { "$ref": "#/panels/gaugeCpuBusy" }
      },
      {
        "x": 3,
        "y": 0,
        "width": 3,
        "height": 4,
        "content": { "$ref": "#/panels/gaugeSystemLoad" }
      }
    ]
  }
}
`,
			result: Layout{
				Kind: KindGridLayout,
				Spec: &GridLayoutSpec{
					Items: []GridItem{
						{
							X:      0,
							Y:      0,
							Width:  3,
							Height: 4,
							Content: &common.JSONRef{
								Ref:  "#/panels/gaugeCpuBusy",
								Path: []string{"panels", "gaugeCpuBusy"},
							},
						},
						{
							X:      3,
							Y:      0,
							Width:  3,
							Height: 4,
							Content: &common.JSONRef{
								Ref:  "#/panels/gaugeSystemLoad",
								Path: []string{"panels", "gaugeSystemLoad"},
							},
						},
					},
				},
			},
		},
		{
			title: "expand layout",
			jason: `
{
  "kind": "Grid",
  "spec": {
    "display": {
      "title": "My Expending Grid",
      "collapse": {
        "open": true
      }
    },
    "items": [
      {
        "x": 0,
        "y": 0,
        "width": 3,
        "height": 4,
        "content": { "$ref": "#/panels/gaugeCpuBusy" }
      },
      {
        "x": 3,
        "y": 0,
        "width": 3,
        "height": 4,
        "content": { "$ref": "#/panels/gaugeSystemLoad" }
      }
    ]
  }
}
`,
			result: Layout{
				Kind: KindGridLayout,
				Spec: &GridLayoutSpec{
					Display: &GridLayoutDisplay{
						Title:    "My Expending Grid",
						Collapse: &GridLayoutCollapse{Open: true},
					},
					Items: []GridItem{
						{
							X:      0,
							Y:      0,
							Width:  3,
							Height: 4,
							Content: &common.JSONRef{
								Ref:  "#/panels/gaugeCpuBusy",
								Path: []string{"panels", "gaugeCpuBusy"},
							},
						},
						{
							X:      3,
							Y:      0,
							Width:  3,
							Height: 4,
							Content: &common.JSONRef{
								Ref:  "#/panels/gaugeSystemLoad",
								Path: []string{"panels", "gaugeSystemLoad"},
							},
						},
					},
				},
			},
		},
		{
			title: "basic tabs layout",
			jason: `
{
  "kind": "Tabs",
  "spec": {
    "tabs": [
      {
        "name": "Overview",
        "items": [
          {
            "x": 0,
            "y": 0,
            "width": 6,
            "height": 4,
            "content": { "$ref": "#/panels/cpuUsage" }
          }
        ]
      },
      {
        "name": "Details",
        "items": [
          {
            "x": 0,
            "y": 0,
            "width": 12,
            "height": 6,
            "content": { "$ref": "#/panels/memoryUsage" }
          }
        ]
      }
    ]
  }
}
`,
			result: Layout{
				Kind: KindTabLayout,
				Spec: &TabLayoutSpec{
					Tabs: []TabItem{
						{
							Name: "Overview",
							Items: []GridItem{
								{
									X:      0,
									Y:      0,
									Width:  6,
									Height: 4,
									Content: &common.JSONRef{
										Ref:  "#/panels/cpuUsage",
										Path: []string{"panels", "cpuUsage"},
									},
								},
							},
						},
						{
							Name: "Details",
							Items: []GridItem{
								{
									X:      0,
									Y:      0,
									Width:  12,
									Height: 6,
									Content: &common.JSONRef{
										Ref:  "#/panels/memoryUsage",
										Path: []string{"panels", "memoryUsage"},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			title: "tabs layout with display and collapse",
			jason: `
{
  "kind": "Tabs",
  "spec": {
    "display": {
      "title": "My Tab Group",
      "collapse": {
        "open": true
      }
    },
    "tabs": [
      {
        "name": "Tab One",
        "items": [
          {
            "x": 0,
            "y": 0,
            "width": 6,
            "height": 4,
            "content": { "$ref": "#/panels/panel1" }
          }
        ]
      }
    ]
  }
}
`,
			result: Layout{
				Kind: KindTabLayout,
				Spec: &TabLayoutSpec{
					Display: &TabLayoutDisplay{
						Title:    "My Tab Group",
						Collapse: &GridLayoutCollapse{Open: true},
					},
					Tabs: []TabItem{
						{
							Name: "Tab One",
							Items: []GridItem{
								{
									X:      0,
									Y:      0,
									Width:  6,
									Height: 4,
									Content: &common.JSONRef{
										Ref:  "#/panels/panel1",
										Path: []string{"panels", "panel1"},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			title: "tabs layout with defaultTab",
			jason: `
{
  "kind": "Tabs",
  "spec": {
    "tabs": [
      {
        "name": "First",
        "items": []
      },
      {
        "name": "Second",
        "items": []
      }
    ],
    "defaultTab": 1
  }
}
`,
			result: Layout{
				Kind: KindTabLayout,
				Spec: &TabLayoutSpec{
					Tabs: []TabItem{
						{
							Name:  "First",
							Items: []GridItem{},
						},
						{
							Name:  "Second",
							Items: []GridItem{},
						},
					},
					DefaultTab: 1,
				},
			},
		},
	}
	for _, test := range testSuite {
		t.Run(test.title, func(t *testing.T) {
			result := Layout{}
			assert.NoError(t, json.Unmarshal([]byte(test.jason), &result))
			assert.Equal(t, test.result, result)
		})
	}
}

func TestUnmarshalYAMLLayout(t *testing.T) {
	testSuite := []struct {
		title  string
		yamele string
		result Layout
	}{
		{
			title: "grid layout",
			yamele: `
kind: "Grid"
spec:
  items:
  - x: 0
    y: 0
    width: 3
    height: 4
    content: 
      $ref: "#/panels/gaugeCpuBusy"
  - x: 3
    y: 0
    width: 3
    height: 4
    content: 
      $ref: "#/panels/gaugeSystemLoad"
`,
			result: Layout{
				Kind: KindGridLayout,
				Spec: &GridLayoutSpec{
					Items: []GridItem{
						{
							X:      0,
							Y:      0,
							Width:  3,
							Height: 4,
							Content: &common.JSONRef{
								Ref:  "#/panels/gaugeCpuBusy",
								Path: []string{"panels", "gaugeCpuBusy"},
							},
						},
						{
							X:      3,
							Y:      0,
							Width:  3,
							Height: 4,
							Content: &common.JSONRef{
								Ref:  "#/panels/gaugeSystemLoad",
								Path: []string{"panels", "gaugeSystemLoad"},
							},
						},
					},
				},
			},
		},
		{
			title: "expand layout",
			yamele: `
kind: "Grid"
spec:
  display:
    title: "My Expending Grid"
    collapse:
      open: true
  items:
  - x: 0
    y: 0
    width: 3
    height: 4
    content:
      $ref: "#/panels/gaugeCpuBusy"
  - x: 3
    y: 0
    width: 3
    height: 4
    content:
      $ref: "#/panels/gaugeSystemLoad"
`,
			result: Layout{
				Kind: KindGridLayout,
				Spec: &GridLayoutSpec{
					Display: &GridLayoutDisplay{
						Title:    "My Expending Grid",
						Collapse: &GridLayoutCollapse{Open: true},
					},
					Items: []GridItem{
						{
							X:      0,
							Y:      0,
							Width:  3,
							Height: 4,
							Content: &common.JSONRef{
								Ref:  "#/panels/gaugeCpuBusy",
								Path: []string{"panels", "gaugeCpuBusy"},
							},
						},
						{
							X:      3,
							Y:      0,
							Width:  3,
							Height: 4,
							Content: &common.JSONRef{
								Ref:  "#/panels/gaugeSystemLoad",
								Path: []string{"panels", "gaugeSystemLoad"},
							},
						},
					},
				},
			},
		},
		{
			title: "basic tabs layout",
			yamele: `
kind: "Tabs"
spec:
  tabs:
  - name: "Overview"
    items:
    - x: 0
      y: 0
      width: 6
      height: 4
      content:
        $ref: "#/panels/cpuUsage"
  - name: "Details"
    items:
    - x: 0
      y: 0
      width: 12
      height: 6
      content:
        $ref: "#/panels/memoryUsage"
`,
			result: Layout{
				Kind: KindTabLayout,
				Spec: &TabLayoutSpec{
					Tabs: []TabItem{
						{
							Name: "Overview",
							Items: []GridItem{
								{
									X:      0,
									Y:      0,
									Width:  6,
									Height: 4,
									Content: &common.JSONRef{
										Ref:  "#/panels/cpuUsage",
										Path: []string{"panels", "cpuUsage"},
									},
								},
							},
						},
						{
							Name: "Details",
							Items: []GridItem{
								{
									X:      0,
									Y:      0,
									Width:  12,
									Height: 6,
									Content: &common.JSONRef{
										Ref:  "#/panels/memoryUsage",
										Path: []string{"panels", "memoryUsage"},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			title: "tabs layout with display and collapse",
			yamele: `
kind: "Tabs"
spec:
  display:
    title: "My Tab Group"
    collapse:
      open: true
  tabs:
  - name: "Tab One"
    items:
    - x: 0
      y: 0
      width: 6
      height: 4
      content:
        $ref: "#/panels/panel1"
`,
			result: Layout{
				Kind: KindTabLayout,
				Spec: &TabLayoutSpec{
					Display: &TabLayoutDisplay{
						Title:    "My Tab Group",
						Collapse: &GridLayoutCollapse{Open: true},
					},
					Tabs: []TabItem{
						{
							Name: "Tab One",
							Items: []GridItem{
								{
									X:      0,
									Y:      0,
									Width:  6,
									Height: 4,
									Content: &common.JSONRef{
										Ref:  "#/panels/panel1",
										Path: []string{"panels", "panel1"},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			title: "tabs layout with defaultTab",
			yamele: `
kind: "Tabs"
spec:
  tabs:
  - name: "First"
    items: []
  - name: "Second"
    items: []
  defaultTab: 1
`,
			result: Layout{
				Kind: KindTabLayout,
				Spec: &TabLayoutSpec{
					Tabs: []TabItem{
						{
							Name:  "First",
							Items: []GridItem{},
						},
						{
							Name:  "Second",
							Items: []GridItem{},
						},
					},
					DefaultTab: 1,
				},
			},
		},
	}
	for _, test := range testSuite {
		t.Run(test.title, func(t *testing.T) {
			result := Layout{}
			assert.NoError(t, yaml.Unmarshal([]byte(test.yamele), &result))
			assert.Equal(t, test.result, result)
		})
	}
}

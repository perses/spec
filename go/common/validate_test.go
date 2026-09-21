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

package common

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateID(t *testing.T) {
	testSuite := []struct {
		title  string
		name   string
		errMsg string
	}{
		{
			title: "simple alphanumeric name",
			name:  "myDashboard1",
		},
		{
			title: "name with allowed special characters",
			name:  "my_dash-board.v2",
		},
		{
			title: "single character name",
			name:  "a",
		},
		{
			title: "name at max length",
			name:  strings.Repeat("a", keyMaxLength),
		},
		{
			title:  "empty name",
			name:   "",
			errMsg: "name cannot be empty",
		},
		{
			title:  "name exceeding max length",
			name:   strings.Repeat("a", keyMaxLength+1),
			errMsg: "cannot contain more than 75 characters",
		},
		{
			title:  "name with space",
			name:   "my dashboard",
			errMsg: `"my dashboard" is not a correct name. It should match the regexp: ^[a-zA-Z0-9_.-]+$`,
		},
		{
			title:  "name with slash",
			name:   "my/dashboard",
			errMsg: `"my/dashboard" is not a correct name. It should match the regexp: ^[a-zA-Z0-9_.-]+$`,
		},
		{
			title:  "name with unicode characters",
			name:   "tableau-é",
			errMsg: `"tableau-é" is not a correct name. It should match the regexp: ^[a-zA-Z0-9_.-]+$`,
		},
		{
			title:  "name containing double dot",
			name:   "my..dashboard",
			errMsg: `"my..dashboard" is not a correct name. It should not contain '..'`,
		},
		{
			title:  "name starting with dot",
			name:   ".dashboard",
			errMsg: `".dashboard" is not a correct name. It should not start or end with '.'`,
		},
		{
			title:  "name ending with dot",
			name:   "dashboard.",
			errMsg: `"dashboard." is not a correct name. It should not start or end with '.'`,
		},
		{
			title:  "name that is only a dot",
			name:   ".",
			errMsg: `"." is not a correct name. It should not start or end with '.'`,
		},
		// Path traversal attack vectors
		{
			title:  "path traversal with double dot only",
			name:   "..",
			errMsg: `".." is not a correct name. It should not contain '..'`,
		},
		{
			title:  "path traversal unix style",
			name:   "../../etc/passwd",
			errMsg: `"../../etc/passwd" is not a correct name. It should match the regexp: ^[a-zA-Z0-9_.-]+$`,
		},
		{
			title:  "path traversal windows style",
			name:   `..\..\windows\system32`,
			errMsg: `"..\\..\\windows\\system32" is not a correct name. It should match the regexp: ^[a-zA-Z0-9_.-]+$`,
		},
		{
			title:  "path traversal with absolute unix path",
			name:   "/etc/passwd",
			errMsg: `"/etc/passwd" is not a correct name. It should match the regexp: ^[a-zA-Z0-9_.-]+$`,
		},
		{
			title:  "path traversal with absolute windows path",
			name:   `C:\Windows\System32`,
			errMsg: `"C:\\Windows\\System32" is not a correct name. It should match the regexp: ^[a-zA-Z0-9_.-]+$`,
		},
		{
			title:  "path traversal url encoded slash",
			name:   "..%2F..%2Fetc%2Fpasswd",
			errMsg: `"..%2F..%2Fetc%2Fpasswd" is not a correct name. It should match the regexp: ^[a-zA-Z0-9_.-]+$`,
		},
		{
			title:  "path traversal url encoded dots",
			name:   "%2e%2e/%2e%2e/etc/passwd",
			errMsg: `"%2e%2e/%2e%2e/etc/passwd" is not a correct name. It should match the regexp: ^[a-zA-Z0-9_.-]+$`,
		},
		{
			title:  "path traversal with null byte",
			name:   "dashboard\x00../../etc/passwd",
			errMsg: `"dashboard\x00../../etc/passwd" is not a correct name. It should match the regexp: ^[a-zA-Z0-9_.-]+$`,
		},
		{
			title:  "path traversal with unicode overlong slash",
			name:   "..\u2215..\u2215etc",
			errMsg: `"..∕..∕etc" is not a correct name. It should match the regexp: ^[a-zA-Z0-9_.-]+$`,
		},
		{
			title:  "path traversal disguised with allowed characters only, triple dot",
			name:   "...dashboard",
			errMsg: `"...dashboard" is not a correct name. It should not contain '..'`,
		},
		{
			title:  "path traversal disguised with allowed characters only, embedded double dot",
			name:   "a..b..c",
			errMsg: `"a..b..c" is not a correct name. It should not contain '..'`,
		},
		{
			title:  "path traversal disguised with allowed characters only, trailing double dot",
			name:   "dashboard..",
			errMsg: `"dashboard.." is not a correct name. It should not contain '..'`,
		},
		{
			title:  "path traversal with dot-dash prefix",
			name:   ".-dashboard",
			errMsg: `".-dashboard" is not a correct name. It should not start or end with '.'`,
		},
		{
			title:  "path traversal with home directory tilde",
			name:   "~/.ssh/id_rsa",
			errMsg: `"~/.ssh/id_rsa" is not a correct name. It should match the regexp: ^[a-zA-Z0-9_.-]+$`,
		},
	}
	for _, test := range testSuite {
		t.Run(test.title, func(t *testing.T) {
			err := ValidateID(test.name)
			if test.errMsg == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, test.errMsg)
			}
		})
	}
}

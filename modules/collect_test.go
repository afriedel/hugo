// Copyright 2019 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package modules

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPathKey(t *testing.T) {
	assert := require.New(t)

	for _, test := range []struct {
		in     string
		expect string
	}{
		{"github.com/foo", "github.com/foo"},
		{"github.com/foo/v2", "github.com/foo"},
		{"github.com/foo/v12", "github.com/foo"},
		{"github.com/foo/v3d", "github.com/foo/v3d"},
		{"MyTheme", "mytheme"},
	} {
		assert.Equal(test.expect, pathKey(test.in))
	}

}

func TestFilterUnwantedMounts(t *testing.T) {

	mounts := []Mount{
		Mount{Source: "a", Target: "b", Lang: "en"},
		Mount{Source: "a", Target: "b", Lang: "en"},
		Mount{Source: "b", Target: "c", Lang: "en"},
	}

	filtered := filterUnwantedMounts(mounts)

	assert := require.New(t)
	assert.Len(filtered, 2)
	assert.Equal([]Mount{Mount{Source: "a", Target: "b", Lang: "en"}, Mount{Source: "b", Target: "c", Lang: "en"}}, filtered)

}

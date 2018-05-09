/*
Copyright 2018 Turbine Labs, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package codec

import (
	"testing"

	tbnflag "github.com/turbinelabs/nonstdlib/flag"
	"github.com/turbinelabs/test/assert"
)

func TestFromFlagsValidate(t *testing.T) {
	ff := &fromFlags{}
	assert.ErrorContains(t, ff.Validate(), "unsupported format: ")

	ff.format = "foo"
	assert.ErrorContains(t, ff.Validate(), "unsupported format: foo")

	ff.format = "json"
	assert.Nil(t, ff.Validate())
	assert.Equal(t, ff.Type(), "json")

	ff.format = "yaml"
	assert.Nil(t, ff.Validate())
	assert.Equal(t, ff.Type(), "yaml")
}

func TestNewFromFlagsJson(t *testing.T) {
	fs := tbnflag.NewTestFlagSet()
	ff := NewFromFlags(fs)
	fs.Parse([]string{"-format", "json"})

	testCodecEncodeDecodeJson(t, ff.Make())
}

func TestNewFromFlagsYaml(t *testing.T) {
	fs := tbnflag.NewTestFlagSet()
	ff := NewFromFlags(fs)
	fs.Parse([]string{"-format", "yaml"})

	testCodecEncodeDecodeYaml(t, ff.Make())
}

func TestNewFromFlagsDefault(t *testing.T) {
	fs := tbnflag.NewTestFlagSet()
	ff := NewFromFlags(fs)
	fs.Parse([]string{"-format", "garbage"})

	testCodecEncodeDecodeJson(t, ff.Make())
}

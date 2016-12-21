/*
Copyright 2017 Turbine Labs, Inc.

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
	"bytes"
	"errors"
	"io"
	"testing"

	"github.com/turbinelabs/test/assert"
)

type brokenRW struct {
	err error
}

func (_ brokenRW) Read(p []byte) (int, error) {
	return 0, io.ErrNoProgress
}

func (_ brokenRW) Write(p []byte) (int, error) {
	return 0, errors.New("Gah!")
}

func testCodecEncodeDecode(t *testing.T, c Codec, want string) {
	buf := &bytes.Buffer{}
	broke := brokenRW{errors.New("Gah!")}
	v := []int{1, 2, 3}
	got := []int{}

	err := c.Encode(v, broke)
	assert.DeepEqual(t, err, broke.err)

	err = c.Encode(v, buf)
	assert.Nil(t, err)
	assert.Equal(t, buf.String(), want)

	err = c.Decode(broke, &got)
	assert.Equal(t, err, io.ErrNoProgress)

	err = c.Decode(buf, &got)
	assert.Nil(t, err)
	assert.DeepEqual(t, got, v)
}

func TestCodecEncodeDecodeJson(t *testing.T) {
	testCodecEncodeDecodeJson(t, NewJson())
}

func TestCodecEncodeDecodeJsonMin(t *testing.T) {
	testCodecEncodeDecode(t, NewJsonMin(), `[1,2,3]`)
}

func testCodecEncodeDecodeJson(t *testing.T, e Codec) {
	testCodecEncodeDecode(t, e, `[
  1,
  2,
  3
]`)
}

func TestCodecEncodeDecodeYaml(t *testing.T) {
	testCodecEncodeDecodeYaml(t, NewYaml())
}

func testCodecEncodeDecodeYaml(t *testing.T, e Codec) {
	testCodecEncodeDecode(t, e, `- 1
- 2
- 3
`)
}

func TestEncodeToString(t *testing.T) {
	c := NewJsonMin()
	m := map[string]int{"one": 1, "two": 2}
	got, err := EncodeToString(c, m)
	assert.Nil(t, err)
	assert.Equal(t, got, `{"one":1,"two":2}`)
}

func TestDecodeFromString(t *testing.T) {
	c := NewJsonMin()
	want := map[string]int{"one": 1, "two": 2}
	in := `{"one":1,"two":2}`
	got := map[string]int{}
	err := DecodeFromString(c, in, &got)
	assert.Nil(t, err)
	assert.DeepEqual(t, got, want)
}

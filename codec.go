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

// Package codec provides a simple interface for encoding and decoding values
// with JSON and YAML implementations, along with a means to configure them
// with a flag.FlagSet.
package codec

//go:generate mockgen -source $GOFILE -destination mock_$GOFILE -package $GOPACKAGE

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/ghodss/yaml"
)

// Codec allows encoding of an interface{} to an io.Writer
// and decoding from an io.Reader. This is a useful alternative
// to the golang binary encoding interfaces--which typically
// go to/from byte slices--when reading to and from files or
// file descriptors
type Codec interface {
	// Encode a value to a writer, based on the --format flag
	Encode(interface{}, io.Writer) error
	// Decode a value from a reader, based on the --format flag
	Decode(io.Reader, interface{}) error
}

// NewJson produces a Codec that reads and writes to JSON. The JSON produced
// by Encode is prettified for human consumption.
func NewJson() Codec {
	return codec{
		func(v interface{}) ([]byte, error) {
			return json.MarshalIndent(v, "", "  ")
		},
		decodeFn(json.Unmarshal),
	}
}

// NewJsonMin returns a Codec that reads and writes to JSON. The JSON written
// is not indented.
func NewJsonMin() Codec {
	return codec{
		func(v interface{}) ([]byte, error) {
			return json.Marshal(v)
		},
		decodeFn(json.Unmarshal),
	}
}

// NewYaml produces an Codec that reads and writes YAML
func NewYaml() Codec {
	return codec{
		encodeFn(yaml.Marshal),
		decodeFn(yaml.Unmarshal),
	}
}

type encodeFn func(interface{}) ([]byte, error)
type decodeFn func([]byte, interface{}) error

type codec struct {
	encodeFn encodeFn
	decodeFn decodeFn
}

func (c codec) Encode(v interface{}, out io.Writer) error {
	data, err := c.encodeFn(v)

	if err != nil {
		return err
	}

	_, err = out.Write(data)
	return err
}

func (c codec) Decode(in io.Reader, v interface{}) error {
	data, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}

	return c.decodeFn(data, v)
}

func EncodeToString(codec Codec, obj interface{}) (string, error) {
	b := bytes.NewBuffer(nil)
	if err := codec.Encode(obj, b); err != nil {
		return "", err
	}
	return b.String(), nil
}

func DecodeFromString(codec Codec, src string, dest interface{}) error {
	r := bytes.NewReader([]byte(src))
	err := codec.Decode(r, dest)
	return err
}

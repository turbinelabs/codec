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

//go:generate mockgen -source $GOFILE -destination mock_$GOFILE -package $GOPACKAGE

import (
	"fmt"

	tbnflag "github.com/turbinelabs/nonstdlib/flag"
)

// FromFlags produces an Codec from a flag.FlagSet.
type FromFlags interface {
	// Validate verifies that all necessary flags are provided
	Validate() error

	// Make produces a new Codec from the provided flags
	Make() Codec

	// Type returns the type of the codec.
	Type() string
}

// NewFromFlags produces a FromFlags, and installs necessary
// configuration flags into the provided FlagSet
func NewFromFlags(flagset tbnflag.FlagSet) FromFlags {
	ff := &fromFlags{}
	flagset.StringVar(&ff.format, "format", "json", "The I/O format (json or yaml)")
	return ff
}

type fromFlags struct {
	format string
}

func (ff *fromFlags) Make() Codec {
	if ff.format == "yaml" {
		return NewYaml()
	}
	return NewJson()
}

func (ff *fromFlags) Validate() error {
	if ff.format != "json" && ff.format != "yaml" {
		return fmt.Errorf("unsupported format: %s", ff.format)
	}
	return nil
}

func (ff *fromFlags) Type() string {
	return ff.format
}

package codec

//go:generate mockgen -source $GOFILE -destination mock_$GOFILE -package $GOPACKAGE

import (
	"flag"
	"fmt"
)

// FromFlags produces an Codec from provided flags
type FromFlags interface {
	// Validate verifies that all necessary flags are provided
	Validate() error

	// Make produces a new Codec from the provided flags
	Make() Codec
}

// NewFromFlags produces a FromFlags, and installs necessary
// configuration flags into the provided FlagSet
func NewFromFlags(flagset *flag.FlagSet) FromFlags {
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

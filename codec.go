package codec

//go:generate mockgen -source $GOFILE -destination mock_$GOFILE -package $GOPACKAGE

import (
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

// NewYaml produces an Codec that reads and writes to JSON
func NewJson() Codec {
	return codec{
		func(v interface{}) ([]byte, error) {
			return json.MarshalIndent(v, "", "  ")
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

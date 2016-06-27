package codec

import (
	"flag"
	"testing"

	"github.com/turbinelabs/test/assert"
)

func TestFromFlagsValidate(t *testing.T) {
	ff := &fromFlags{}
	assert.ErrorContains(t, ff.Validate(), "unsupported format: ")

	ff.format = "foo"
	assert.ErrorContains(t, ff.Validate(), "unsupported format: foo")

	ff.format = "json"
	assert.Nil(t, ff.Validate())

	ff.format = "yaml"
	assert.Nil(t, ff.Validate())
}

func TestNewFromFlagsJson(t *testing.T) {
	var fs flag.FlagSet
	ff := NewFromFlags(&fs)
	fs.Parse([]string{"-format", "json"})

	testCodecEncodeDecodeJson(t, ff.Make())
}

func TestNewFromFlagsYaml(t *testing.T) {
	var fs flag.FlagSet
	ff := NewFromFlags(&fs)
	fs.Parse([]string{"-format", "yaml"})

	testCodecEncodeDecodeYaml(t, ff.Make())
}

func TestNewFromFlagsDefault(t *testing.T) {
	var fs flag.FlagSet
	ff := NewFromFlags(&fs)
	fs.Parse([]string{"-format", "garbage"})

	testCodecEncodeDecodeJson(t, ff.Make())
}

package clui

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/akesi/akesi/tools/test"
)

func TestGetEnv(t *testing.T) {
	// setup
	os.Setenv("ENVVAR", "DATA\n")
	os.Setenv("BLANK", "\n")

	// success
	data, err := GetEnv("ENVVAR")
	assert.Equal(t, "DATA", data)
	assert.NoError(t, err)

	// failure - does not exist
	data, err = GetEnv("NOPE")
	assert.Empty(t, data)
	test.AssertErr(t, err, `cannot access envvar "NOPE": does not exist`)

	// failure - is blank
	data, err = GetEnv("BLANK")
	assert.Empty(t, data)
	test.AssertErr(t, err, `cannot access envvar "BLANK": is blank`)
}

func TestParseArgs(t *testing.T) {
	// success - real argument
	amap, err := ParseArgs([]string{"name"}, []string{"argument"})
	assert.Equal(t, map[string]string{"name": "argument"}, amap)
	assert.NoError(t, err)

	// success - default argument
	amap, err = ParseArgs([]string{"name:default"}, nil)
	assert.Equal(t, map[string]string{"name": "default"}, amap)
	assert.NoError(t, err)

	// failure - not provided
	amap, err = ParseArgs([]string{"name"}, nil)
	assert.Empty(t, amap)
	test.AssertErr(t, err, `argument "name" not provided`)
}

func TestSplitArgs(t *testing.T) {
	// success - no arguments
	name, argus := SplitArgs(nil)
	assert.Empty(t, name)
	assert.Empty(t, argus)

	// success - one argument
	name, argus = SplitArgs([]string{"name"})
	assert.Equal(t, "name", name)
	assert.Empty(t, argus)

	// success - multiple arguments
	name, argus = SplitArgs([]string{"name", "argument"})
	assert.Equal(t, "name", name)
	assert.Equal(t, []string{"argument"}, argus)
}

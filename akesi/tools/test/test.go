// Package test implements unit testing helper function.
package test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockFiles is a base:body map of mock testing files.
var MockFiles = map[string]string{
	"alpha.extn":    "Alpha note.\n",
	"bravo.extn":    "Bravo note.\n",
	"charlie.trash": "Charlie note (deleted).\n",
}

// AssertDire asserts a directory's files are equal to a base:body map.
func AssertDire(t *testing.T, dire string, fmap map[string]string) {
	for base, body := range fmap {
		orig := filepath.Join(dire, base)
		AssertFile(t, orig, body)
	}
}

// AssertErr asserts an error's message matches a regular expression.
func AssertErr(t *testing.T, err error, regx string) {
	assert.Regexp(t, regx, err.Error())
}

// AssertFile asserts a file's body is equal to a string.
func AssertFile(t *testing.T, orig, body string) {
	bytes, err := os.ReadFile(orig)
	assert.Equal(t, body, string(bytes))
	assert.NoError(t, err)
}

// TempDire returns a temporary directory populated from MockFiles.
func TempDire(t *testing.T) string {
	dire := t.TempDir()
	for base, body := range MockFiles {
		dest := filepath.Join(dire, base)
		if err := os.WriteFile(dest, []byte(body), 0666); err != nil {
			panic(err)
		}
	}

	return dire
}

// TempFile returns a temporary file populated from a MockFiles value.
func TempFile(t *testing.T, base string) string {
	dire := t.TempDir()
	dest := filepath.Join(dire, base)
	if err := os.WriteFile(dest, []byte(MockFiles[base]), 0666); err != nil {
		panic(err)
	}

	return dest
}

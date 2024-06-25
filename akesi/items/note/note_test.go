package note

import (
	"io/fs"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/akesi/akesi/tools/test"
)

func xNote(t *testing.T) *Note {
	orig := test.TempFile(t, "alpha.extn")
	return New(orig, 0666)
}

func TestNew(t *testing.T) {
	// success
	note := xNote(t)
	assert.Contains(t, note.Orig, "alpha.extn")
	assert.Equal(t, fs.FileMode(0666), note.Mode)
}

func TestDelete(t *testing.T) {
	// setup
	note := xNote(t)
	dest := strings.Replace(note.Orig, ".extn", ".trash", 1)

	// success
	err := note.Delete()
	assert.NoFileExists(t, note.Orig)
	assert.FileExists(t, dest)
	assert.NoError(t, err)
}

func TestMatch(t *testing.T) {
	// setup
	note := xNote(t)

	// success - true
	ok := note.Match("ALPH")
	assert.True(t, ok)

	// success - false
	ok = note.Match("NOPE")
	assert.False(t, ok)
}

func TestName(t *testing.T) {
	// setup
	note := xNote(t)

	// success
	name := note.Name()
	assert.Equal(t, "alpha", name)
}

func TestRead(t *testing.T) {
	// setup
	note := xNote(t)

	// success
	body, err := note.Read()
	assert.Equal(t, "Alpha note.\n", body)
	assert.NoError(t, err)
}

func TestSearch(t *testing.T) {
	// setup
	note := xNote(t)

	// success - true
	ok, err := note.Search("ALPH")
	assert.True(t, ok)
	assert.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	// setup
	note := xNote(t)

	// success
	err := note.Update("body\n")
	test.AssertFile(t, note.Orig, "body\n")
	assert.NoError(t, err)
}

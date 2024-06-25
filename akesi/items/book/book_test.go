package book

import (
	"io/fs"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/akesi/akesi/items/note"
	"github.com/stvmln86/akesi/akesi/tools/test"
)

func xBook(t *testing.T) *Book {
	dire := test.TempDire(t)
	return New(dire, ".extn", fs.FileMode(0666))
}

func TestNew(t *testing.T) {
	// success
	book := xBook(t)
	assert.NotEmpty(t, book.Dire)
	assert.Equal(t, ".extn", book.Extn)
	assert.Equal(t, fs.FileMode(0666), book.Mode)
}

func TestCreate(t *testing.T) {
	// setup
	book := xBook(t)

	// success
	note, err := book.Create("name", "body\n")
	test.AssertFile(t, note.Orig, "body\n")
	assert.NoError(t, err)
}

func TestFilter(t *testing.T) {
	// setup
	book := xBook(t)

	// success
	notes, err := book.Filter(func(note *note.Note) (bool, error) {
		return note.Name() == "alpha", nil
	})
	assert.Len(t, notes, 1)
	assert.Equal(t, "alpha", notes[0].Name())
	assert.NoError(t, err)
}

func TestGet(t *testing.T) {
	// setup
	book := xBook(t)

	// success
	note, err := book.Get("alpha")
	assert.Contains(t, note.Orig, "alpha.extn")
	assert.NoError(t, err)

	// failure - does not exist
	note, err = book.Get("nope")
	assert.Nil(t, note)
	test.AssertErr(t, err, `cannot get file .*: does not exist`)
}

func TestList(t *testing.T) {
	// setup
	book := xBook(t)

	// success
	notes, err := book.List()
	assert.Len(t, notes, 2)
	assert.Equal(t, "alpha", notes[0].Name())
	assert.Equal(t, "bravo", notes[1].Name())
	assert.NoError(t, err)
}

func TestMatch(t *testing.T) {
	// setup
	book := xBook(t)

	// success
	notes, err := book.Match("ALPH")
	assert.Len(t, notes, 1)
	assert.Equal(t, "alpha", notes[0].Name())
	assert.NoError(t, err)
}

func TestSearch(t *testing.T) {
	// setup
	book := xBook(t)

	// success
	notes, err := book.Search("ALPH")
	assert.Len(t, notes, 1)
	assert.Equal(t, "alpha", notes[0].Name())
	assert.NoError(t, err)
}

// Package book implements the Book type and methods.
package book

import (
	"fmt"
	"io/fs"

	"github.com/stvmln86/akesi/akesi/items/note"
	"github.com/stvmln86/akesi/akesi/tools/file"
	"github.com/stvmln86/akesi/akesi/tools/neat"
	"github.com/stvmln86/akesi/akesi/tools/path"
)

// Book is a single directory of Note files.
type Book struct {
	Dire string
	Extn string
	Mode fs.FileMode
}

// New returns a new Book.
func New(dire, extn string, mode fs.FileMode) *Book {
	dire = neat.Path(dire)
	extn = neat.Extn(extn)
	return &Book{dire, extn, mode}
}

// Create creates and returns a new Note with a string.
func (b *Book) Create(name, body string) (*note.Note, error) {
	name = neat.Name(name)
	body = neat.Body(body)
	dest := path.Join(b.Dire, name, b.Extn)
	if err := file.Create(dest, body, b.Mode); err != nil {
		return nil, err
	}

	return note.New(dest, b.Mode), nil
}

// Filter returns all existing Notes passing a filter function.
func (b *Book) Filter(ffun func(*note.Note) (bool, error)) ([]*note.Note, error) {
	var goods []*note.Note
	notes, err := b.List()
	if err != nil {
		return nil, err
	}

	for _, note := range notes {
		ok, err := ffun(note)
		switch {
		case err != nil:
			return nil, err
		case ok:
			goods = append(goods, note)
		}
	}

	return goods, nil
}

// Get returns an existing Note by name.
func (b *Book) Get(name string) (*note.Note, error) {
	name = neat.Name(name)
	dest := path.Join(b.Dire, name, b.Extn)
	if !file.Exists(dest) {
		return nil, fmt.Errorf("cannot get file %q: does not exist", dest)
	}

	return note.New(dest, b.Mode), nil
}

// List returns all existing Notes in the Book.
func (b *Book) List() ([]*note.Note, error) {
	var notes []*note.Note
	origs, err := file.List(b.Dire, b.Extn)
	if err != nil {
		return nil, err
	}

	for _, orig := range origs {
		notes = append(notes, note.New(orig, b.Mode))
	}

	return notes, nil
}

// Match returns all existing Notes with names containing a substring.
func (b *Book) Match(term string) ([]*note.Note, error) {
	return b.Filter(func(note *note.Note) (bool, error) {
		return note.Match(term), nil
	})
}

// Search returns all existing Notes with bodies containing a substring.
func (b *Book) Search(term string) ([]*note.Note, error) {
	return b.Filter(func(note *note.Note) (bool, error) {
		return note.Search(term)
	})
}

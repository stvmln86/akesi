// Package file implements file system access functions.
package file

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/stvmln86/akesi/akesi/tools/path"
)

// Create writes a new file with a string..
func Create(dest, body string, mode fs.FileMode) error {
	if Exists(dest) {
		return fmt.Errorf("cannot create file %q: already exists", dest)
	}

	if err := os.WriteFile(dest, []byte(body), mode); err != nil {
		return fmt.Errorf("cannot create file %q: %w", dest, err)
	}

	return nil
}

// Delete renames a file to the ".trash" file extension.
func Delete(orig string) error {
	if !Exists(orig) {
		return fmt.Errorf("cannot delete file %q: does not exist", orig)
	}

	dire := path.Dire(orig)
	name := path.Name(orig)
	dest := path.Join(dire, name, ".trash")
	if err := os.Rename(orig, dest); err != nil {
		return fmt.Errorf("cannot delete file %q: %w", orig, err)
	}

	return nil
}

// Exists returns true if a file or directory path exists.
func Exists(orig string) bool {
	_, err := os.Stat(orig)
	return !errors.Is(err, fs.ErrNotExist)
}

// List returns all paths in a directory matching an extension.
func List(dire, extn string) ([]string, error) {
	if !Exists(dire) {
		return nil, fmt.Errorf("cannot list directory %q: does not exist", dire)
	}

	pttn := filepath.Join(dire, "*"+extn)
	elems, err := filepath.Glob(pttn)
	if err != nil {
		return nil, fmt.Errorf("cannot list directory %q: %w", dire, err)
	}

	sort.Strings(elems)
	return elems, nil
}

// Read returns a file's body as a string.
func Read(orig string) (string, error) {
	if !Exists(orig) {
		return "", fmt.Errorf("cannot read file %q: does not exist", orig)
	}

	bytes, err := os.ReadFile(orig)
	if err != nil {
		return "", fmt.Errorf("cannot read file %q: %w", orig, err)
	}

	return string(bytes), nil
}

// Search returns true if a file's body contains a substring.
func Search(orig, term string) (bool, error) {
	if !Exists(orig) {
		return false, fmt.Errorf("cannot search file %q: does not exist", orig)
	}

	bytes, err := os.ReadFile(orig)
	if err != nil {
		return false, fmt.Errorf("cannot search file %q: %w", orig, err)
	}

	body := strings.ToLower(string(bytes))
	term = strings.ToLower(term)
	return strings.Contains(body, term), nil
}

// Update overwrites an existing file's body with a string.
func Update(orig, body string, mode fs.FileMode) error {
	if !Exists(orig) {
		return fmt.Errorf("cannot update file %q: does not exist", orig)
	}

	if err := os.WriteFile(orig, []byte(body), mode); err != nil {
		return fmt.Errorf("cannot update file %q: %w", orig, err)
	}

	return nil
}

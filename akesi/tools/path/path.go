// Package path implements file path manipulation functions.
package path

import (
	"path/filepath"
	"strings"
)

// Base returns a path's file name and extension.
func Base(orig string) string {
	return filepath.Base(orig)
}

// Dire returns a path's parent directory.
func Dire(orig string) string {
	return filepath.Dir(orig)
}

// Extn returns a path's file extension with a leading dot.
func Extn(orig string) string {
	_, extn, _ := strings.Cut(orig, ".")
	return "." + extn
}

// Join returns a joined path from a directory, file name and extension.
func Join(dire, name, extn string) string {
	return filepath.Join(dire, name+extn)
}

// Match returns true if a path's file name contains a substring.
func Match(orig, term string) bool {
	name := Name(orig)
	name = strings.ToLower(orig)
	term = strings.ToLower(term)
	return strings.Contains(name, term)
}

// Name returns a path's file name.
func Name(orig string) string {
	base := filepath.Base(orig)
	name, _, _ := strings.Cut(base, ".")
	return name
}

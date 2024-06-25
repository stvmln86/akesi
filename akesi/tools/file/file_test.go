package file

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/akesi/akesi/tools/test"
)

func TestCreate(t *testing.T) {
	// setup
	dest := filepath.Join(t.TempDir(), "alpha.extn")

	// success
	err := Create(dest, "body\n", 0666)
	test.AssertFile(t, dest, "body\n")
	assert.NoError(t, err)

	// failure - already exists
	err = Create(dest, "body\n", 0666)
	test.AssertErr(t, err, "cannot create file .*: already exists")
}

func TestDelete(t *testing.T) {
	// setup
	orig := test.TempFile(t, "alpha.extn")
	dest := strings.Replace(orig, ".extn", ".trash", 1)

	// success
	err := Delete(orig)
	assert.NoFileExists(t, orig)
	assert.FileExists(t, dest)
	assert.NoError(t, err)

	// failure - does not exist
	err = Delete(orig)
	test.AssertErr(t, err, "cannot delete file .*: does not exist")
}

func TestExists(t *testing.T) {
	// setup
	orig := test.TempFile(t, "alpha.extn")

	// success - true
	ok := Exists(orig)
	assert.True(t, ok)

	// success - false
	ok = Exists("/nope")
	assert.False(t, ok)
}

func TestList(t *testing.T) {
	// setup
	dire := test.TempDire(t)

	// success
	elems, err := List(dire, ".extn")
	assert.Equal(t, []string{
		filepath.Join(dire, "alpha.extn"),
		filepath.Join(dire, "bravo.extn"),
	}, elems)
	assert.NoError(t, err)

	// failure - does not exist
	elems, err = List("/nope", ".extn")
	assert.Empty(t, elems)
	test.AssertErr(t, err, "cannot list directory .*: does not exist")
}

func TestRead(t *testing.T) {
	// setup
	orig := test.TempFile(t, "alpha.extn")

	// success
	body, err := Read(orig)
	assert.Equal(t, "Alpha note.\n", body)
	assert.NoError(t, err)

	// failure - does not exist
	body, err = Read("/nope")
	assert.Empty(t, body)
	test.AssertErr(t, err, "cannot read file .*: does not exist")
}

func TestSearch(t *testing.T) {
	// setup
	orig := test.TempFile(t, "alpha.extn")

	// success - true
	ok, err := Search(orig, "ALPH")
	assert.True(t, ok)

	// success - false
	ok, err = Search(orig, "NOPE")
	assert.False(t, ok)

	// failure - does not exist
	ok, err = Search("/nope", "NOPE")
	assert.False(t, ok)
	test.AssertErr(t, err, "cannot search file .*: does not exist")
}

func TestUpdate(t *testing.T) {
	// setup
	orig := test.TempFile(t, "alpha.extn")

	// success
	err := Update(orig, "body\n", 0666)
	test.AssertFile(t, orig, "body\n")
	assert.NoError(t, err)

	// failure - does not exist
	err = Update("/nope", "body\n", 0666)
	test.AssertErr(t, err, "cannot update file .*: does not exist")
}

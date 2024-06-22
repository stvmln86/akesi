package path

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBase(t *testing.T) {
	// success
	base := Base("/dire/name.extn")
	assert.Equal(t, "name.extn", base)
}

func TestDire(t *testing.T) {
	// success
	dire := Dire("/dire/name.extn")
	assert.Equal(t, "/dire", dire)
}

func TestExtn(t *testing.T) {
	// success
	extn := Extn("/dire/name.extn")
	assert.Equal(t, ".extn", extn)
}

func TestJoin(t *testing.T) {
	// success
	dest := Join("/dire", "name", ".extn")
	assert.Equal(t, "/dire/name.extn", dest)
}

func TestMatch(t *testing.T) {
	// success - true
	ok := Match("/dire/name.extn", "NAME")
	assert.True(t, ok)

	// success - false
	ok = Match("/dire/name.extn", "NOPE")
	assert.False(t, ok)
}

func TestName(t *testing.T) {
	// success
	name := Name("/dire/name.extn")
	assert.Equal(t, "name", name)
}

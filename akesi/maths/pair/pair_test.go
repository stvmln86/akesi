package pair

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	// success
	p := New(1, 2)
	assert.Equal(t, 1, p.X)
	assert.Equal(t, 2, p.Y)
}

func TestAdd(t *testing.T) {
	// success
	p := New(1, 2).Add(New(1, 2))
	assert.Equal(t, 2, p.X)
	assert.Equal(t, 4, p.Y)
}

func TestAddN(t *testing.T) {
	// success
	p := New(1, 2).AddN(1)
	assert.Equal(t, 2, p.X)
	assert.Equal(t, 3, p.Y)
}

func TestEqual(t *testing.T) {
	// success - true
	b := New(1, 2).Equal(New(1, 2))
	assert.True(t, b)

	// success - false x
	b = New(1, 2).Equal(New(0, 2))
	assert.False(t, b)

	// success - false y
	b = New(1, 2).Equal(New(1, 0))
	assert.False(t, b)
}

func TestLesser(t *testing.T) {
	// success - true
	b := New(1, 2).Lesser(New(2, 3))
	assert.True(t, b)

	// success - false x
	b = New(1, 2).Lesser(New(1, 3))
	assert.False(t, b)

	// success - false y
	b = New(1, 2).Lesser(New(2, 2))
	assert.False(t, b)
}

func TestGreater(t *testing.T) {
	// success - true
	b := New(1, 2).Greater(New(0, 1))
	assert.True(t, b)

	// success - false x
	b = New(1, 2).Greater(New(1, 1))
	assert.False(t, b)

	// success - false y
	b = New(1, 2).Greater(New(0, 2))
	assert.False(t, b)
}

func TestSub(t *testing.T) {
	// success
	p := New(1, 2).Sub(New(1, 2))
	assert.Equal(t, 0, p.X)
	assert.Equal(t, 0, p.Y)
}

func TestSubN(t *testing.T) {
	// success
	p := New(1, 2).SubN(1)
	assert.Equal(t, 0, p.X)
	assert.Equal(t, 1, p.Y)
}

// Package test implements unit-testing helper functions.
package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/akesi/akesi/maths/pair"
)

// AssertPair asserts a Pair equals two integers.
func AssertPair(t *testing.T, p *pair.Pair, x, y int) {
	assert.Equal(t, x, p.X)
	assert.Equal(t, y, p.Y)
}

// AssertPairs asserts a Pair slice equals an integer slice.
func AssertPairs(t *testing.T, ps []*pair.Pair, ns [][]int) {
	assert.Equal(t, len(ns), len(ps))
	for i, n := range ns {
		AssertPair(t, ps[i], n[0], n[1])
	}
}

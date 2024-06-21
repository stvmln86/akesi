package test

import (
	"testing"

	"github.com/stvmln86/akesi/akesi/maths/pair"
)

func TestAssertPair(t *testing.T) {
	// success
	AssertPair(t, pair.New(1, 2), 1, 2)
}

func TestAssertPairs(t *testing.T) {
	// success
	AssertPairs(t, []*pair.Pair{pair.New(1, 2)}, [][]int{{1, 2}})
}

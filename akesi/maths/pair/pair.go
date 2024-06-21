// Package pair implements the Pair type and methods.
package pair

// Pair is a single two-dimensional co-ordinate pair.
type Pair struct {
	X int
	Y int
}

// New returns a new Pair.
func New(x, y int) *Pair {
	return &Pair{x, y}
}

// Add returns the sum of the Pair and another Pair.
func (p *Pair) Add(p2 *Pair) *Pair {
	return New(p.X+p2.X, p.Y+p2.Y)
}

// AddN returns the sum of the Pair and an integer.
func (p *Pair) AddN(n int) *Pair {
	return New(p.X+n, p.Y+n)
}

// Equal returns true if the Pair is equal to another Pair.
func (p *Pair) Equal(p2 *Pair) bool {
	return p.X == p2.X && p.Y == p2.Y
}

// Lesser returns true if the Pair is less than another Pair.
func (p *Pair) Lesser(p2 *Pair) bool {
	return p.X < p2.X && p.Y < p2.Y
}

// Greater returns true if the Pair is greater than another Pair.
func (p *Pair) Greater(p2 *Pair) bool {
	return p.X > p2.X && p.Y > p2.Y
}

// Sub returns the difference of the Pair and another Pair.
func (p *Pair) Sub(p2 *Pair) *Pair {
	return New(p.X-p2.X, p.Y-p2.Y)
}

// SubN returns the difference of the Pair and an integer.
func (p *Pair) SubN(n int) *Pair {
	return New(p.X-n, p.Y-n)
}

package ffrepresentation

import "iter"

type Coordinates struct {
	X, Y, Z uint
}

type Point[E FElement] struct {
	Included []*Line[E]
}

type Line[E FElement] struct {
	Includes []*Point[E]
}

type FFProjectivePlane[E FElement] struct {
	Field  Field[E]
	Points map[Coordinates]E
	Lines  map[Coordinates]E
}

func (p *FFProjectivePlane[E]) Canonical(c Coordinates) Coordinates {
	if !p.Field.Equal(p.Field.Ele(c.X), p.Field.Zero()) {
		inv := p.Field.Inv(p.Field.Ele(c.X))
		return Coordinates{X: p.Field.One().Uint(), Y: p.Field.Mult(p.Field.Ele(c.Y), inv).Uint(), Z: p.Field.Mult(p.Field.Ele(c.Z), inv).Uint()}
	}
	if !p.Field.Equal(p.Field.Ele(c.Y), p.Field.Zero()) {
		return Coordinates{X: c.X, Y: p.Field.One().Uint(), Z: p.Field.Mult(p.Field.Ele(c.Z), p.Field.Inv(p.Field.Ele(c.Y))).Uint()}
	}
	return c
}

func CoordinatesFrom(start Coordinates, k uint) iter.Seq[Coordinates] {
	return func(yield func(Coordinates) bool) {
		c := start

		for {
			if !yield(c) {
				return
			}

			// Increment Z.
			if c.Z+1 < k {
				c.Z++
				continue
			}

			// Start next Y.
			c.Z = 0
			if c.Y+1 < k {
				c.Y++
				continue
			}

			// Start next X.
			c.Y = 0
			if c.X+1 < k {
				c.X++
				continue
			}

			return
		}
	}
}

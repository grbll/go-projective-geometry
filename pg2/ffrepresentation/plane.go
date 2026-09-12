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
	Points map[Coordinates]*Point[E]
	Lines  map[Coordinates]*Line[E]
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
			if c.Z+1 < k {
				c.Z++
				continue
			}
			c.Z = 0
			if c.Y+1 < k {
				c.Y++
				continue
			}
			c.Y = 0
			if c.X+1 < k {
				c.X++
				continue
			}

			return
		}
	}
}

func New[E FElement](field Field[E]) *FFProjectivePlane[E] {
	q := field.Card()
	p := &FFProjectivePlane[E]{
		Field:  field,
		Points: make(map[Coordinates]*Point[E], q*q*q-1),
		Lines:  make(map[Coordinates]*Line[E], q*q*q-1),
	}

	card := field.Card()

	for c := range CoordinatesFrom(Coordinates{0, 0, 1}, card) {
		pc := &Point[E]{Included: make([]*Line[E], 0, q+1)}
		lc := &Line[E]{Includes: make([]*Point[E], 0, q+1)}

		for i := uint(1); i < card; i++ {
			m := field.Ele(i)

			cc := Coordinates{
				X: field.Mult(field.Ele(c.X), m).Uint(),
				Y: field.Mult(field.Ele(c.Y), m).Uint(),
				Z: field.Mult(field.Ele(c.Z), m).Uint(),
			}

			p.Points[cc] = pc
			p.Lines[cc] = lc
		}
	}

	return p
}

func (p FFProjectivePlane[E]) generateLineInclusions(c Coordinates) *[]*Point[E] {
	field := p.Field
	line := p.Lines[c]
	if c.Z == 0 {
		line.Includes = append(line.Includes, p.Points[Coordinates{X: 0, Y: 0, Z: 1}])
		if c.Y == 0 {
			for i := uint(0); i < field.Card(); i++ {
				line.Includes = append(line.Includes, p.Points[Coordinates{X: 0, Y: 1, Z: i}])
			}
			return &line.Includes
		} else { //c.Y==0
			Y := field.Mult(field.Neg(field.Ele(c.X)), field.Inv(field.Ele(c.Y))).Uint()
			for i := uint(0); i < p.Field.Card(); i++ {
				line.Includes = append(line.Includes, p.Points[Coordinates{X: 0, Y: Y, Z: i}])
			}
		}
	}
	return &line.Includes
}

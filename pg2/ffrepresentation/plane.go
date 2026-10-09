package ffrepresentation

import "iter"

type Coordinates struct {
	X, Y, Z uint
}

type Point[E FElement] struct {
	Rep      Coordinates
	Included []*Line[E]
}

type Line[E FElement] struct {
	Rep      Coordinates
	Includes []*Point[E]
}

type FFProjectivePlane[E FElement] struct {
	Field     Field[E]
	Points    []*Point[E]
	Lines     []*Line[E]
	Canonical map[Coordinates]uint
}

func (p *FFProjectivePlane[E]) Canonize(c Coordinates) Coordinates {
	if !p.Field.Equal(p.Field.Ele(c.X), p.Field.Zero()) {
		inv := p.Field.Inv(p.Field.Ele(c.X))
		return Coordinates{X: p.Field.One().Uint(), Y: p.Field.Mult(p.Field.Ele(c.Y), inv).Uint(), Z: p.Field.Mult(p.Field.Ele(c.Z), inv).Uint()}
	}
	if !p.Field.Equal(p.Field.Ele(c.Y), p.Field.Zero()) {
		return Coordinates{X: c.X, Y: p.Field.One().Uint(), Z: p.Field.Mult(p.Field.Ele(c.Z), p.Field.Inv(p.Field.Ele(c.Y))).Uint()}
	}
	return Coordinates{X: 0, Y: 0, Z: 1}
}

func CoordinatesFrom(start Coordinates, k uint) iter.Seq[Coordinates] {
	return func(yield func(Coordinates) bool) {
		c := start

		for {
			if !yield(c) {
				return
			}
			if c.X == 0 {
				if c.Y == 0 {
					c.Y = 1
					c.Z = 0
					continue
				} else { //c.Y==1
					if c.Z+1 < k {
						c.Z = c.Z + 1
						continue
					} else { //01k
						c.X = 1
						c.Y = 0
						c.Z = 0
						continue
					}
				}
			} else { //c.X==1
				if c.Z+1 < k {
					c.Z = c.Z + 1
					continue
				} else { //c.Z==k
					if c.Y+1 < k {
						c.Z = 0
						c.Y = c.Y + 1
						continue
					} else { //c.y==k
						return
					}
				}
			}
		}
	}
}

func New[E FElement](field Field[E]) *FFProjectivePlane[E] {
	q := field.Card()
	p := &FFProjectivePlane[E]{
		Field:     field,
		Points:    make([]*Point[E], 0, q*q+q+1),
		Lines:     make([]*Line[E], 0, q*q+q+1),
		Canonical: make(map[Coordinates]uint, q*q*q-1),
	}

	step := uint(0)
	for c := range CoordinatesFrom(Coordinates{0, 0, 1}, q) {
		p.Points = append(p.Points, &Point[E]{Rep: c, Included: make([]*Line[E], 0, q+1)})
		p.Lines = append(p.Lines, &Line[E]{Rep: c, Includes: make([]*Point[E], 0, q+1)})
		for i := uint(1); i < q; i++ {
			m := field.Ele(i)

			cc := Coordinates{
				X: field.Mult(field.Ele(c.X), m).Uint(),
				Y: field.Mult(field.Ele(c.Y), m).Uint(),
				Z: field.Mult(field.Ele(c.Z), m).Uint(),
			}
			p.Canonical[cc] = step
		}
		step++
	}

	p.GenerateInclusions()

	return p
}

func (p FFProjectivePlane[E]) GenerateInclusions() {
	for _, line := range p.Lines {
		f := p.Field
		c := line.Rep
		if c.Z == 0 {
			point := p.Points[p.Canonical[Coordinates{X: 0, Y: 0, Z: 1}]]
			line.Includes = append(line.Includes, point)
			point.Included = append(point.Included, line)
			if c.Y == 0 {
				for i := uint(0); i < f.Card(); i++ {
					point := p.Points[p.Canonical[Coordinates{X: 0, Y: 1, Z: i}]]
					line.Includes = append(line.Includes, point)
					point.Included = append(point.Included, line)
				}
			} else { //c.Y!=0
				Y := f.Neg(f.Mult(f.Ele(1), f.Inv(f.Ele(c.Y)))).Uint()
				for i := uint(0); i < f.Card(); i++ {
					point := p.Points[p.Canonical[Coordinates{X: 1, Y: Y, Z: i}]]
					line.Includes = append(line.Includes, point)
					point.Included = append(point.Included, line)
				}
			}
		} else { // c.Z!=0
			Z := f.Neg(f.Mult(f.Ele(c.Y), f.Inv(f.Ele(c.Z)))).Uint()
			point := p.Points[p.Canonical[Coordinates{X: 0, Y: 1, Z: Z}]]
			line.Includes = append(line.Includes, point)
			point.Included = append(point.Included, line)
			for i := uint(0); i < f.Card(); i++ {
				Z := f.Neg(f.Mult(f.Inv(f.Ele(c.Z)), f.Add(f.Ele(c.X), f.Mult(f.Ele(c.Y), f.Ele(i))))).Uint()
				point := p.Points[p.Canonical[Coordinates{X: 1, Y: i, Z: Z}]]
				line.Includes = append(line.Includes, point)
				point.Included = append(point.Included, line)
			}
		}
	}
}

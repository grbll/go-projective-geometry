package ffrepresentation

type Coordinates[E FElement] struct {
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
	Points map[Coordinates[E]]E
	Lines  map[Coordinates[E]]E
}

func (p *FFProjectivePlane[E]) Canonical(c Coordinates[E]) Coordinates[E] {
	if !p.Field.Equal(p.Field.Ele(c.X), p.Field.Zero()) {
		inv := p.Field.Inv(p.Field.Ele(c.X))
		return Coordinates[E]{X: p.Field.One().Uint(), Y: p.Field.Mult(p.Field.Ele(c.Y), inv).Uint(), Z: p.Field.Mult(p.Field.Ele(c.Z), inv).Uint()}
	}
	if !p.Field.Equal(p.Field.Ele(c.Y), p.Field.Zero()) {
		return Coordinates[E]{X: c.X, Y: p.Field.One().Uint(), Z: p.Field.Mult(p.Field.Ele(c.Z), p.Field.Inv(p.Field.Ele(c.Y))).Uint()}
	}
	return c
}

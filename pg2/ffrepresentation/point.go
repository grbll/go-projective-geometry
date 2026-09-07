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
	field  Field[E]
	Points map[Coordinates[E]]E
	Lines  map[Coordinates[E]]E
}

func (p *FFProjectivePlane[E]) Canonical(c Coordinates[E]) Coordinates[E] {
	if !p.field.Equal(p.field.Ele(c.X), p.field.Zero()) {
		inv := p.field.MInv(p.field.Ele(c.X))
		return Coordinates[E]{X: p.field.One().Uint(), Y: p.field.Mult(p.field.Ele(c.Y), inv).Uint(), Z: p.field.Mult(p.field.Ele(c.Z), inv).Uint()}
	}
	if !p.field.Equal(p.field.Ele(c.Y), p.field.Zero()) {
		return Coordinates[E]{X: c.X, Y: p.field.One().Uint(), Z: p.field.Mult(p.field.Ele(c.Z), p.field.MInv(p.field.Ele(c.Y))).Uint()}
	}
	return c
}

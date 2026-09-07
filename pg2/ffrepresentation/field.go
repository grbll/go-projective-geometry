package ffrepresentation

type FElement interface {
	Uint() uint
}

type Field[E FElement] interface {
	Elements() []E
	Zero() E
	One() E
	Neg(E) E
	Inv(E) E
	Add(E, E) E
	Mult(E, E) E
	Equal(E, E) bool
	Ele(uint) E
}

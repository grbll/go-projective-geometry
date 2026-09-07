package ffrepresentation

type FElement interface {
	UInt() uint
}

type Field[E FElement] interface {
	Elements() []E
	Zero() E
	One() E
	Inv(E) E
	MInv(E) E
	Add(E, E) E
	Mult(E, E) E
	Equal(E, E) bool
	Ele(uint) E
}

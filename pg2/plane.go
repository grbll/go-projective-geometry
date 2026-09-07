package pg2

type Plane[P any, L any] interface {
	Points() []P
	Lines() []L

	Included(P) []L
	Includes(L) []P

	Span(P, P) L
	Intersection(L, L) P
}

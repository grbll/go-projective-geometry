package main

import (
	"github.com/glycerine/algobra/finitefield/ff"
)

type AlgobraElement struct {
	ff.Element
}

func (e AlgobraElement) Uint() uint {
	return e.Element.(interface {
		Uint() uint
	}).Uint()
}

type AlgobraField struct {
	ff.Field
}

func (f AlgobraField) Elements() []AlgobraElement {
	elements := f.Elements()
	result := make([]AlgobraElement, len(elements))

	for i, e := range elements {
		result[i] = AlgobraElement{e}
	}

	return result
}

func (f AlgobraField) Zero() AlgobraElement {
	return AlgobraElement{f.Field.Zero()}
}

func (f AlgobraField) One() AlgobraElement {
	return AlgobraElement{f.Field.One()}
}

func (f AlgobraField) Neg(a AlgobraElement) AlgobraElement {
	return AlgobraElement{a.Neg()}
}

func (f AlgobraField) Inv(a AlgobraElement) AlgobraElement {
	return AlgobraElement{a.Inv()}
}

func (f AlgobraField) Add(a, b AlgobraElement) AlgobraElement {
	return AlgobraElement{a.Plus(b.Element)}
}

func (f AlgobraField) Mult(a, b AlgobraElement) AlgobraElement {
	return AlgobraElement{a.Times(b.Element)}
}

func (f AlgobraField) Equal(a, b AlgobraElement) bool {
	return a.Equal(b.Element)
}

func (f AlgobraField) Ele(n uint) AlgobraElement {
	return AlgobraElement{f.Field.ElementFromUnsigned(n)}
}

func New(f ff.Field) AlgobraField {
	return AlgobraField{Field: f}
}

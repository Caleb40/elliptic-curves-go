package elliptic_curve

import (
	"fmt"
)

type FieldElement struct {
	// class FieldElement (equivalent in python)
	order uint64 // field order
	num   uint64 // value of the given element in the field
}

func NewFieldElement(order uint64, num uint64) *FieldElement {
	/*
		init function for FieldElement, __init__ from python
		(initialization)
	*/
	if num >= order {
		err := fmt.Sprintf("Num not in the range of 0 to %d", order-1)
		panic(err)
	}
	return &FieldElement{
		order: order,
		num:   num,
	}
}

func (f *FieldElement) String() string {
	// equivalent of "__repr__" in python
	return fmt.Sprintf("FieldElement{order:%d, num:%d}\n", f.order, f.num)
}

func (f *FieldElement) EqualTo(other *FieldElement) bool {
	return f.order == other.order && f.num == other.num
}

func (f *FieldElement) Add(other *FieldElement) *FieldElement {
	if f.order != other.order {
		panic("Add operation can only be performed on elements with the same order.")
	}
	// remember the modulo
	// operator overloading for +, __add__ python
	return NewFieldElement(f.order, (f.num+other.num)%f.order)
}

func (f *FieldElement) Negate() *FieldElement {
	// (a + b) % order === order - a
	//return NewFieldElement(f.order, (f.order-f.num)%f.order)
	return NewFieldElement(f.order, (-f.num)%f.order)
}

func (f *FieldElement) Subtract(other *FieldElement) *FieldElement {
	if f.order != other.order {
		panic("Subtract operation can only be performed on elements with the same order.")
	}
	return f.Add(other.Negate())
}

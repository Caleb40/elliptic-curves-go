package elliptic_curve

import (
	"fmt"
	"math"
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

func (f *FieldElement) checkOrder(other *FieldElement) {
	if f.order != other.order {
		panic(fmt.Sprintf(
			"order mismatch: source element has order %d, target element has order %d",
			f.order, other.order,
		))
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
	f.checkOrder(other)
	// remember the modulo
	// operator overloading for +, __add__ python
	return NewFieldElement(f.order, (f.num+other.num)%f.order)
}

func (f *FieldElement) Negate() *FieldElement {
	// (a + b) % order === order - a
	//return NewFieldElement(f.order, (-f.num)%f.order) equivalent
	return NewFieldElement(f.order, (f.order-f.num)%f.order)
}

func (f *FieldElement) Subtract(other *FieldElement) *FieldElement {
	// Simply the sum of the element and the negation of the other
	f.checkOrder(other)
	return f.Add(other.Negate())
}

func (f *FieldElement) Multiply(other *FieldElement) *FieldElement {
	// Arithmetic multiplication over the modulo of the order
	f.checkOrder(other)
	return NewFieldElement(f.order, (f.num*other.num)%f.order)
}

func (f *FieldElement) Power(power int64) *FieldElement {
	// Arithmetic power over the modulo of the order
	return NewFieldElement(f.order,
		uint64(math.Pow(float64(f.num), float64(power)))%f.order)
}

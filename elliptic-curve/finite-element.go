package elliptic_curve

import (
	"fmt"
	"math/big"
)

type FieldElement struct {
	// class FieldElement (equivalent in python)
	order *big.Int // field order
	num   *big.Int // value of the given element in the field
}

func NewFieldElement(order *big.Int, num *big.Int) *FieldElement {
	/*
		init function for FieldElement, __init__ from python
		(initialization)
	*/
	if order.Cmp(num) == -1 {
		err := fmt.Sprintf("Num not in the range of 0 to %s", order.Sub(order, big.NewInt(1)))
		panic(err)
	}
	return &FieldElement{
		order: order,
		num:   num,
	}
}

func (f *FieldElement) checkOrder(other *FieldElement) {
	if f.order.Cmp(other.order) != 0 {
		panic(fmt.Sprintf(
			"order mismatch: source element has order %d, target element has order %d",
			f.order, other.order,
		))
	}
}

func (f *FieldElement) String() string {
	// equivalent of "__repr__" in python
	return fmt.Sprintf("FieldElement {order:%s, num:%s}\n", f.order.String(), f.num.String())
}

func (f *FieldElement) EqualTo(other *FieldElement) bool {
	return f.order.Cmp(other.order) == 0 && f.num.Cmp(other.num) == 0
}

func (f *FieldElement) Add(other *FieldElement) *FieldElement {
	f.checkOrder(other)
	// remember the modulo
	// operator overloading for +, __add__ python
	var op big.Int
	return NewFieldElement(f.order, op.Mod(op.Add(f.num, other.num), f.order))
}

func (f *FieldElement) Negate() *FieldElement {
	// (a + b) % order === order - a
	//return NewFieldElement(f.order, (-f.num)%f.order) equivalent
	var op = big.Int{}
	return NewFieldElement(f.order, op.Sub(f.order, f.num))
}

func (f *FieldElement) Subtract(other *FieldElement) *FieldElement {
	// Simply the sum of the element and the negation of the other
	f.checkOrder(other)
	return f.Add(other.Negate())
}

func (f *FieldElement) Multiply(other *FieldElement) *FieldElement {
	// Arithmetic multiplication over the modulo of the order
	f.checkOrder(other)
	var op big.Int
	mul := op.Mul(f.num, other.num)
	return NewFieldElement(f.order, op.Mod(mul, f.order))
}

func (f *FieldElement) Power(power *big.Int) *FieldElement {
	// Arithmetic power over the modulo of the order
	// k ^ (p-1) % p = 1, power > p-1 => power %p(p-1)

	var op big.Int
	t := op.Mod(power, op.Sub(f.order, big.NewInt(int64(1))))
	powerRes := op.Exp(f.num, t, nil)
	return NewFieldElement(f.order, op.Mod(powerRes, f.order))
}

func (f *FieldElement) ScalarMul(val *big.Int) *FieldElement {
	var op big.Int
	mul := op.Mul(f.num, val)
	return NewFieldElement(f.order, op.Mod(mul, f.order))
}

func (f *FieldElement) Divide(other *FieldElement) *FieldElement {
	f.checkOrder(other)
	// a / b => a * b^(p-2)
	var op = big.Int{}
	otherReverse := other.Power(op.Sub(f.order, big.NewInt(2)))
	return f.Multiply(otherReverse)
}

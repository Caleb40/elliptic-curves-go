package main

import (
	ecc "elliptic_curve"
	"fmt"
	"math/big"
	"math/rand"
)

func SolveField19MultiplySet() {
	// randomly select a num from 1 - 18
	minimum := 1
	maximum := 18
	k := rand.Intn(maximum-minimum) + minimum
	fmt.Printf("Randomly selected value, k is %d\n", k)
	element := ecc.NewFieldElement(big.NewInt(19), big.NewInt(int64(k)))

	for i := 0; i < 19; i++ {
		fmt.Printf("Element: %d, multiplied with %d is: %v", k, i,
			element.ScalarMul(big.NewInt(int64(i))))
	}
}

/*
If p is a field of order, p = 7, 11, 17, 19, 31...
Select any element, k in the field with order p, compute k ^ (p-1). What is the result for this?
=> { 1^(p-1), 2^(p-1), ..., (p-1)^(p-1) % p }
∴ for any element, k in the field with order p => k ^ (p-1) % p == 1
*/

func ComputeFieldOrderPower() {
	orders := []int{7, 11, 17, 19, 31}
	for _, p := range orders {
		fmt.Printf("The value of p is %d\n", p)
		for i := 1; i < p; i++ {
			elem := ecc.NewFieldElement(big.NewInt(int64(p)), big.NewInt(int64(i)))
			fmt.Printf("For element: %v, it's value to the power of p-1 is %v\n", elem, elem.Power(big.NewInt(int64(p-1))))
		}
	}
}

func main() {
	f44 := ecc.NewFieldElement(big.NewInt(57), big.NewInt(44))
	f33 := ecc.NewFieldElement(big.NewInt(57), big.NewInt(33))
	res := f44.Add(f33)
	fmt.Printf("Field element 44 add to field element 33 is %v", res)
	// [-44], the negated value of 44 is 57-44 = 13 [mod 57]
	fmt.Printf("Negated value of FieldElement 44 is %v\n", res.Negate())

	fmt.Printf("Field Element 44 - 33 is: %v", f44.Subtract(f33))
	fmt.Printf("Field element 33 - 44 is %v", f33.Subtract(f44))
	// f33 - f44 == f46 => f46 + f44 == f33 [mod 57]
	f46 := ecc.NewFieldElement(big.NewInt(57), big.NewInt(46))
	fmt.Printf("Field Element 46 + 44 is: %v\n", f46.Add(f44))

	fmt.Printf("Product of Element 46 with itself is: %v", f46.Multiply(f46))
	fmt.Printf("Exponent of Element 46 with 2 is: %v", f46.Power(big.NewInt(2)))
	// Check:
	fmt.Printf("CHECK: 46^2 and 46 * 46 are equivalent: %v\n\n", f46.Multiply(f46).EqualTo(f46.Power(big.NewInt(2))))

	SolveField19MultiplySet()

	ComputeFieldOrderPower()

	// Division:
	f2 := ecc.NewFieldElement(big.NewInt(int64(19)), big.NewInt(int64(2)))
	f7 := ecc.NewFieldElement(big.NewInt(int64(19)), big.NewInt(int64(7)))
	fmt.Printf("Field element 2/7 with order 19 is %v\n", f2.Divide(f7))

	f46 = ecc.NewFieldElement(big.NewInt(int64(57)), big.NewInt(int64(46)))

	fmt.Printf("Field element 46 * 46 with order 57 is %v&n", f46.Multiply(f46))
	fmt.Printf("Field element 46 to the power of 58 is %v\n", f46.Power(big.NewInt(int64(58))))

	a := ecc.NewFieldElement(big.NewInt(13), big.NewInt(7))
	b := ecc.NewFieldElement(big.NewInt(13), big.NewInt(8))
	fmt.Printf("Field element 7^-3 with order 13 == field element 8: %v\n", a.Power(big.NewInt(-3)).EqualTo(b))
}

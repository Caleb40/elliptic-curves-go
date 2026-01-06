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
}

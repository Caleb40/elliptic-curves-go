package main

import (
	ecc "elliptic_curve"
	"fmt"
)

func main() {
	f44 := ecc.NewFieldElement(57, 44)
	f33 := ecc.NewFieldElement(57, 33)
	res := f44.Add(f33)
	fmt.Printf("Field element 44 add to field element 33 is %v\n\n", res)
	// -44 negate of 44 is 57-44 = 13
	fmt.Printf("Negated value of FieldElement 44 is %v\n\n", res.Negate())

	fmt.Printf("Field Element 44 - 33 is: %v\n", f44.Subtract(f33))
	fmt.Printf("Field element 33 - 44 is %v\n", f33.Subtract(f44))
	// f33 - f44 == f46 => f46 + f44 == f33 [mod 57]
	f46 := ecc.NewFieldElement(57, 46).Add(f44)
	fmt.Printf("Field Element 46 + 44 is: %v\n", f46)
}

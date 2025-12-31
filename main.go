package main

import (
	ecc "elliptic_curve"
	"fmt"
)

func main() {
	f44 := ecc.NewFieldElement(57, 44)
	f33 := ecc.NewFieldElement(57, 33)
	res := f44.Add(f33)
	fmt.Println("Field element 44 add to field element 33 is %v\n", res)
	// -44 negate of 44 is 57-44 = 13
	fmt.Println("Negated value of FieldElement 44 is %v\n", res.Negate())
}

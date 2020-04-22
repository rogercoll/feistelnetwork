// Package feistel provides the basic network to create chipers with provided functions
package feistel

import (
	"fmt"
)

//Xor Bittwise operator between two int32 arrays
func xor(a,b *[]rune) (*[]rune) {
	c := make([]rune, len(*a))
	for i := 0; i < len(*a); i++ {
		c[i] = (*a)[i]^(*b)[i]
	}
	return &c
}

//Splits a string into two equal parts and retuns it's ASCII values
func splitString(a string) (*[]rune, *[]rune, error) {
	//each character in rune corresponts to a int32 ASCII value, no need for a uint32
	if len(a)%2 != 0 {
		return nil, nil, fmt.Errorf("Not able(for now) to split an odd lenght string: %v", a)
	}
	l := []rune(a[:len(a)/2])
	r := []rune(a[len(a)/2:])
	return &l, &r, nil
}

// Main function which will allow a string and a crypthographic function as parameters
func Run() error {
	return nil
}
// Package feistel provides the basic network to create chipers with provided functions.

// In cryptography, a Feistel cipher is a symmetric structure used in the construction of block ciphers, 
// named after the German-born physicist and cryptographer Horst Feistel who did pioneering research while working for IBM (USA); // part. To reconstruct the original secret, a minimum number of parts is required. In the threshold
// it is also commonly known as a Feistel network. A large proportion of block ciphers use the scheme, including the Data Encryption Standard (DES).
// The Feistel structure has the advantage that encryption and decryption operations are very similar, even identical in some cases, 
// requiring only a reversal of the key schedule. 
// Therefore, the size of the code or circuitry required to implement such a cipher is nearly halved.
// See (wiki page) https://en.wikipedia.org/wiki/Feistel_cipher.
package feistel

import (
	"fmt"
)

type Parts struct {
	l,r *[]rune
	rounds int
}

// cipher algorithm that takes an string and returns the encoded string
type cipher func([]rune) []rune

func New(message string, r int) (*Parts, error) {
	left,right, err := splitString(message)
	if err != nil {
		return nil, err
	}
	n := Parts{rounds: r, l: left, r: right}
    return &n, nil
}

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

func joinRune(a *[]rune, b *[]rune) []byte {
	s := make([]byte, len(*a) + len(*b))
	for i := 0; i < len(*a); i++ {
		s[i] = byte((*a)[i])
	}
	for i := len(*a); i < len(*a) + len(*b); i++ {
		s[i] = byte((*b)[i-len(*a)])
	}
	return s
}

func toString(a *[]rune, b *[]rune) string {
	s := ""
	for i := 0; i < len(*a); i++ {
		s += string((*a)[i])
	}
	for i := len(*a); i < len(*a) + len(*b); i++ {
		s += string((*b)[i-len(*a)])
	}
	return s
}

// Main function which will allow a string and a crypthographic function as parameters
// n is the number of rounds to run the feisel network 
func (p *Parts)Run(fn cipher) ([]byte, error) {
	for i := 0; i < p.rounds; i++ {
		c := *p.r
		e := fn(*p.r)
		p.r = xor(p.l,&e)
		p.l = &c
	}
	//last round swaps l and r ?? yees is okey, decryption is exactly the same proces
	d := *p.r
	p.r = p.l
	p.l = &d
	return joinRune(p.l,p.r), nil
}


func (p *Parts)Reverse(fn cipher) (string,error) {
	for i := 0; i < p.rounds; i++ {
		c := *p.r
		e := fn(*p.r)
		p.r = xor(p.l,&e)
		p.l = &c
	}
	d := *p.r
	p.r = p.l
	p.l = &d
	return toString(p.l,p.r),nil
}
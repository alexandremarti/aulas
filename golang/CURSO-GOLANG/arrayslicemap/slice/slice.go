package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} // array
	s1 := []int{1, 2, 3}  // slice
	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}

	// Slice não é um array! Slide define um pedaço de um array.
	s2 := a2[1:3]
	fmt.Println(a2, s2)

	s3 := a2[:2] // novo slice, mas aponta para o mesmo array
	fmt.Println(a2, s3)

	// vc pode imaginar um slice como: tamanho e um ponteiro para um elemento de um array
	s4 := s2[:1]
	fmt.Println(s2, s4)

	// como  slice é um ponteiro para um pedaço do array, quando eu altero o array reflete em todos os slices
	a2[1] = 9
	fmt.Println(s2, s4, a2)

	// da mesma forma, se eu altero o valor em um slice na verdade altero no array e em todos os slices
	// que referenciam aquele pedaço
	s2[0] = 7
	fmt.Println(s2[0], s2, s4, a2)

}

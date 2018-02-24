package main

import "fmt"

func main() {
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))

	// ao adicionar um novo elemente além do tamanho, o array interno é dobrado
	// se o meu array estiver com 20 passa para 40
	// se o meu array estiver com 40 passa para 80
	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))
}

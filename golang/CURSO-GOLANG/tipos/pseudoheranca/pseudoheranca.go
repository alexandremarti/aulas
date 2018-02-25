package main

import "fmt"

type carro struct {
	nome            string
	velocidadeAtual int
}

// adicioneio o método
func (c carro) getVelocidade() int {
	return c.velocidadeAtual
}

type ferrari struct {
	carro       // campos anonimos
	turboLigado bool
}

// adicionei o método
func (f ferrari) getNome() string {
	return f.nome
}

func main() {
	f := ferrari{}
	f.nome = "F40"
	f.velocidadeAtual = 80
	f.turboLigado = true

	// exemplos que eu adicionei
	fmt.Printf("A ferrari %s está com turbo ligado? %v\n", f.nome, f.turboLigado)
	fmt.Println(f)

	fmt.Printf("Carro %s\n", f.getNome())
	fmt.Printf("Carro %s tem velocidade de %d\n", f.getNome(), f.getVelocidade())
}

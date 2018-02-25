package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
	// pode adicionar mais métodos
}

type bwm7 struct{}

func (b bwm7) ligarTurbo() {
	fmt.Println("Turbo...")
}

func (b bwm7) fazerBaliza() {
	fmt.Println("Baliza...")
}

func main() {
	var b esportivoLuxuoso = bwm7{}
	b.ligarTurbo()
	b.fazerBaliza()

	// demonstra que pode ser usado qualquer das interfaces
	var c luxuoso = bwm7{}
	c.fazerBaliza()
	// nesse caso não posso acessar o método ligarTurbo pois a interface luxuso não implementa ele!
	//c.ligarTurbo()

	var d esportivo = bwm7{}
	d.ligarTurbo()

}

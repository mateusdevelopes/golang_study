package main

import (
	"fmt"
	// Renomeação de import
	m "math"
)

func main() {
	const PI float64 = 3.14
	var raio = 3.2 // tipo implicito de float64

	// Forma reduzida de criar uma var ":=" - atribui e inicializa.
	area := PI * m.Pow(raio, 2)
	fmt.Println("A área é =", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	//Multiplas atribuições e inicializações
	fmt.Println(a, b, c, d)
	var e, f bool = true, false
	fmt.Println(e, f)
	g, h, i := 2, false, "Olhá só!"
	fmt.Println(g, h, i)
}

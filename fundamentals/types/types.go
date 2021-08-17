package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(1, 2, 1000)
	fmt.Println("\nLiteral integer é", reflect.TypeOf(32000))

	var by byte = 255
	fmt.Println("\nO byte é", reflect.TypeOf(by))

	//sem sinal (só positivo) .... uint8 uint16 uint32 uint64
	i1 := math.MaxInt64
	fmt.Println("\nO valor máximo int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabela Unicode (int32)
	fmt.Println("\nO valor de i2 é", i2)
	fmt.Println("O tipo de i2 é", reflect.TypeOf(i1))

	var x float32 = 49.99
	// ou var x float32(49.99)
	fmt.Println("\nO tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99", reflect.TypeOf(49.99))

	bo := true
	fmt.Println("\nO tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	s1 := "Olá mano!"
	fmt.Println("\n" + s1 + ":D")
	fmt.Println("O tamanho da string é", len(s1))

	s2 := `Uma
	frase
	nova`
	fmt.Println("\nTexto:", s2)
	fmt.Println("Tam do texto:", len(s2))

	char := 'a'
	fmt.Println("\nO tipo é", reflect.TypeOf(char))

	// PATICULARIDADES DE VALORES DO TIPO ZERO!
	var a int
	var b float64
	var c bool
	var d string
	var e *int

	fmt.Printf("%v %v %v %v %q %v", a, b, c, d, d, e)
}

package main

import "fmt"

func main() {
	fmt.Print("Em uma mesma")
	fmt.Print("linha!")

	fmt.Println("Em uma nova")
	fmt.Println("linha!")

	x := 3.149

	xs := fmt.Sprint(x)

	fmt.Println("O valor de xs é" + xs)
	fmt.Println("O valor de xs é", xs)

	fmt.Printf("O valor de x é %.2f", x)

	a := 1
	b := 1.9999
	c := false
	d := "Opa!"

	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}

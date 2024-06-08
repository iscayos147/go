package main

import (
	"fmt"
	"math"
)

var (
	cateto_a   float64
	cateto_b   float64
	hipotenusa float64
	area       float64
	perimetro  float64
)

const precision int = 2

func main() {

	fmt.Print("Digite el valor del cateto a: ")
	fmt.Scanln(&cateto_a)

	fmt.Print("Digite el valor del cateto b: ")
	fmt.Scanln(&cateto_b)

	hipotenusa = math.Sqrt(math.Pow(cateto_a, 2) + math.Pow(cateto_b, 2))

	area = (cateto_b * cateto_a) / 2

	perimetro = cateto_a + cateto_b + hipotenusa

	fmt.Printf("area del triangulo: %.*f y el perimetro es %.*f\n", precision, area, precision, perimetro)

}

package main

import "fmt"

func main() {

	hello()
	hello2("ismael")
	fmt.Println(hello3("castillo"))

	//llamar funcion que devuelve mas de un valor
	suma, multi := calculadora(8, 9)

	fmt.Println(suma, multi)
}

//declarar funciones

// funcion sin parametros
func hello() {
	fmt.Println("Hola")
}

// funcion con parametros
func hello2(nombre string) {
	fmt.Println(nombre)
}

// funcion que retorna datos
func hello3(apellido string) string {
	return apellido
}

// funcion que retorna multiples valores
func calculadora(num1, num2 int) (int, int) { //aqui se se esta indicandoque devolvera 2 enteros

	sum := num1 + num2
	mul := num1 * num2
	return sum, mul
}

// otra manera de hacerlo
func calculadora2(num1, num2 int) (sum, mul int) { //aqui se se esta indicandoque devolvera 2 enteros

	sum = num1 + num2
	mul = num1 * num2
	return
}

//al tener un for dentro de una funcion y colocar return dentro de ese for, esto hace que se termine el ciclo y la funcion ya que se le indica que la funcion termino

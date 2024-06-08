package main

import (
	"fmt"
	"strconv"
)

func main() {

	//VARIABLES
	/*
		info:
			- Declaracion: para declarar variables se utiliza la palabra reservada va
						   manera convencional: var <nombre variable> <tipo dato> = <asignacion>
			- Advertencias:
				* si en go declaramos una variable y no la utilizamos nos marcara error
	*/
	var firstName string

	//Declarar multiples variables del mismo tipo
	var secondName, lastName string

	//Declarar multiles variables de diferentes tipos
	var (
		cadena1, cadena2 string
		age              int
		estatura         float32
	)

	//Declarar de manera mas simple
	var variable1, variable2, numero = "var 1", "var 2", 96

	/*
		DECLARAR VARIABLES SIN USAR LA PALABRA RESERVADA var
			- Extructura: <nombre variable> := <asignacion>
			- Advertencia: Este tipo de variables solo se pueden declarar dentro de las funciones
	*/
	prueba := 985

	// asignacion de variables, se hace con =
	firstName = "ismael"
	secondName = "de jesus"
	lastName = "castillo hoyos"
	cadena1 = "prueba"
	cadena2 = "variables"
	estatura = 3.6 // float es con .
	age = 27

	//CONSTANTES
	/*
		info:
			- Declaracion: se usa la palabra reservada const
			- Datos generales:
				* si se define la constante y no se utilizan entonces no hay problemas
				* se recomienda iniciar nombres de las constantes con MAYUSCULAS
	*/
	const Pi = 3.14

	//Declarar varias constantes al tiempo de diferente tipo de dato
	const (
		X = 100
		Y = 0b1010 //tipo de dato binario
		Z = 0o12   //OCTAL
		W = 0xFF   //hexadecimal
	)

	//TIPOS DE DATOS
	/*
		info:
			- numericos:
				* int: enteros de todo tipo
				* uint: solo almacena enteros positivos
				* float32 y float64: punto decimal o flotante
				* byte: un alias de uint8, solo almacena datos enteros positivos de 0 a 255, se utilzia para representar caracteres assi
				* rune: almacena numeros int32 se usa para representar caracteres unicode en go (emogis)

			- logicos:
				* bool: booleanos true o false

			- string: cadena de caracteres

		tips:
			- conocer los valores maximos y minimos de un tipo de dato numerico: usar la clase math.maxint64 ..
			- defaul<tipodato>: permite obtener el valor por defecto de un tipo de dato


	*/
	var integer int32 = 123
	var integerpos uint16 = 123
	var flotantes float32 = 23.6
	var boleanos bool = true
	var a byte = 'a' //imprimiria el valor de codigo assi
	var caden3 = "hola"
	var r rune = '❤'
	fmt.Println(caden3[0]) //esto imprime el valor en byte de el primer caracter

	//CONVERSIONES
	/*
		info:
			-numericos:
				* int32() e int64(): funciones nativas, se introduce el valor a convertir dentro e igual apra float

			- string a numero: se usa el pakete strconv

		obserbaciones: no se peude hacer operaciones entre datos por ejemplo int32 e int64, es necesario convertir
	*/

	//de string a entero
	str := "100"
	strconvertido, _ := strconv.Atoi(str) //devuelve un entero y un error sino se pudo hacer la conversion, sino queremos capturar el error se usa _

	//de entero a string
	n := 52
	stt := strconv.Itoa(n)

	//PAQUETE FMT
	/*
		info:
			- uso: imprimir salida en consola
			- documentacion: https://pkg.go.dev/fmt@go1.22.3
			- investigar sobre sprintf y printf
			- Scan recibe la erferencia de la memoria, Scanln(&nombre de la variable), solo permite escanear un valor

	*/

	//EL PAQUETE MATH SE USA PARA OPERACIONES ARITMETICAS:

	//IMPRIMIENDO VALORES
	fmt.Println("nombre completo: ", firstName, secondName, lastName) //se puede concatenar con ,
	fmt.Println(cadena1, cadena2, estatura, age)
	fmt.Println(variable1, variable2, numero)
	fmt.Println(prueba)
	fmt.Print("constante: ", Pi)
	fmt.Println(X, Y, Z, W)
	fmt.Println(integer, integerpos, flotantes, boleanos, a)
	fmt.Println(r)
	fmt.Println(strconvertido + strconvertido)
	fmt.Println(stt)

}

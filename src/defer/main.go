package main

import (
	"fmt"
	"os"
)

func main() {

	//la palabra clave defer se utiliza para posponer la ejecucion de una funcion hasta que la funcion que la tiene halla finalizado

	defer fmt.Println(3) // esta se va a ejecutar al final, defer lo que hace es posponer la ejecucion de esta funcion hasta que termine la ejecucion del main y ahi si la ejecuta
	fmt.Println(1)
	fmt.Println(2)

	//CUANDO TODAS LAS FUNCIONES TIENEN DEFER LAS FUNCIONES ESPECIFICADAS SE AGREGAN A UNA PILA DE EJECUCION DE FUNCIONES DEFINIDAS
	fmt.Println(6)
	fmt.Println(5)
	fmt.Println(4)

	//CREANDO UN ARCHIVO EJEMPPLO DE DEFER
	file, error2 := os.Create("hola.txt") // retorna un puntero de tipo file y un error
	if error2 != nil {
		fmt.Println(error2)
		return
	}
	defer file.Close() // como se debe usar el close en el segundo if entonces se pospone hasta elfinal de la ejecucion del main

	_, error3 := file.Write([]byte("hola ismael"))
	if error3 != nil {
		fmt.Println(error3)

		return
	}
	//cerrando el flujo de entrada y salida de dato
	//file.Close()
}

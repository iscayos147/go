package main

import "fmt"

func main() {

	/*
		REBANADAS
			Info: representa una seccion contigua de un areglo, permiten areglar y eliminar elementos de manera dinamica

	*/

	//declarar rebanadas
	//no se necesita tener un areglo declarado, pero si hay un areglo declarado igual podemos crear una rebanada de este areglo
	//se usa la palabra reserbada var asi: var nombre rebanada []tipo dato

	var rebanada []int

	//declarar e inicializar una rebanada
	diasSemana := []string{"domingo", "lunes", "martes", "miercoles", "jueves"}

	//crear una rebanada a partir de otra rebanada
	diarebanada := diasSemana[0:4]

	//dinamismo de una rebanada
	//agregar elemtos: se usa la funcion append(rebanada, cantidad de datos a agregar)
	diarebanada = append(diarebanada, "viernes") //esta funcion devuielve una rebanada nueva

	//quitar elementos de una rebanada, se hace creando una nueva rebanada quitando una posicion
	diarebanada = append(diarebanada[:1], diarebanada[2:]...) //aqui se elimina la posicion 1, desde el indice 0 hasta el 1 y luego le mandamos otra rebanada desde donde queremos que inicie
	//en pocas palabras se deja de tener una posicion y se concatena la otra

	//longitud de una rebanada, se hace con len, pero la maxima cantidad de optiene con cap, esto para el caso de ser una rebanada de un tipo de objeto o de otra rebanada
	fmt.Println(len(diarebanada))
	fmt.Println(cap(diarebanada)) // la capacidad: cantidad de elementos que puede almacenar

	//funcion make (crea rebanadas)
	nombres := make([]string, 5, 10) //los numeros significan 5 longitud inicial, 10 capacidad maxima inicial
	nombres[1] = "ismael"            // asi se peuden agregar hasta los primeros 5 elementos, pero despue sd elos 5 ya toca usar la funcion append

	// funcion copy, se usa para copiar rebanadas
	nombres2 := make([]string, 5)
	copy(nombres2, nombres) //estamos copiando a nombres en nombres2

	fmt.Println(rebanada, diasSemana, diarebanada)
	fmt.Println(nombres)
	fmt.Println(nombres2)
}

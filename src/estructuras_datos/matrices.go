package main

import "fmt"

func main() {

	//MATRICES
	/*
		Info:
			- se definen con un tamaño fijo y pueden ser de un tipo de dato creado por el usuario

	*/

	//declaracion de una matriz
	var matriz [2]int

	//modificar elemntos de una matriz
	matriz[1] = 15

	//delarar una matriz inicializada
	var mtx = [3]int{74, 85, 96}

	//declarar e inicializar una matriz dinamica(sin tamaño definido)
	var mtx2 = [...]int{74, 85, 96}

	//recorrer una matriz con un for
	for i := 0; i < len(mtx2); i++ {
		fmt.Println(mtx2[i])
	}

	//iterar la matriz con range
	for index, value := range mtx2 {
		fmt.Printf("indice: %d, valor %d\n", index, value)
	}

	//iterar la matriz con range para solo imprimir el valor
	for _, value := range mtx2 {
		fmt.Printf("valor %d\n", value)
	}

	//definior una matriz multi demensional var nombre matriz [3][3]int, la primera llave representa el numero de filas y la segunda el numero de columnas
	var mtbd [3][3]int

	var mtbd2 = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	fmt.Println(matriz, mtx, mtx2[2], mtbd, mtbd2)

}

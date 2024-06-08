package main

import "fmt"

func main() {

	//MAPAS: estructura de datos dinamica que permiten almacenar y recuperar datos por pares clave valor

	// definicion de un mapa e inicializacion
	colors := map[string]string{ //dentro de los corchetes va el tipo de dato de la clave y afuera va el tipo de dato del valor de la clave
		"rojo":  "#FF0000",
		"verde": "#00FF00",
		"azul":  "#0000FF", //LLEVA COMA AL FINAL
	}

	//imprimir el valor no la clave de un elemento
	fmt.Println(colors["rojo"])

	//agregar nueva clave
	colors["negro"] = "#000000"

	//los mapas no respeta el orden en que almacenamos los datos
	fmt.Println(colors)

	//se le pueden asignar los valores a variables comunes
	valor := colors["negro"]
	fmt.Println(valor)

	//se pueden hacer validaciones e imprimirlas
	valor2, ok := colors["negro"] //si existe la clave negro se guarda en el ok
	fmt.Println(valor2, ok)

	//usando if
	if ok {
		fmt.Println("si existe la clave")
	} else {
		fmt.Println("no existe")
	}

	//usando if con asignacion
	if valor3, ok := colors["amarillo"]; ok {
		fmt.Println("si existe la clave", valor3)
	} else {
		fmt.Println("no existe")
	}

	//eliminar un elemento, se usa delete
	delete(colors, "azul")

	//como iterar un mapa
	for clave, valor := range colors {
		fmt.Printf("la clave es %s y el valor es %s\n", clave, valor)
	}
}

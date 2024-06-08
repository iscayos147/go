package main

import "fmt"

// ESTRUCTURAS, struc
// un tipo de dato personalizado compuesto por diferentes elementos con nombres y tipos de dato especificos
// sintaxis: type nombre_dela_estructura struc {}

type persona struct {
	nombre string
	edad   int
	correo string
}

// declarar un metodo, se debe poner un receptor y es colocando un parentecis antes del nombre de la funcion
func (p *persona) diHola() {
	fmt.Println("hola mi nombre es", p.nombre)
}

func main() {

	//instanciar y usar punteros
	//los metodos pertenencen a una estructura colocando los receptores
	p2 := persona{"alex", 25, "correo electronico"}
	p2.diHola()

	fmt.Println("")

	//instanciar una estructura, es crear datos de tipo persona
	var juan persona
	juan.nombre = "juan andres"
	juan.correo = "juan34@gmail.com"
	juan.edad = 24

	fmt.Println(juan)

	//asiganr e instanciar en una sola linea las estructuras
	pedro := persona{"pedro", 12, "coreo"}
	fmt.Println(pedro)

	//en go se peuden definir metodos para realizar operaciones a las variables de las estruturas
	//los metodos se definen como funciones que tiene un receptor que es un puntero

	//PUNTERO: VARIABLE QUE ALMACENA LA DIRECCION DE LA MEMORIA DE OTRA VARIABLE
	//se utilizan para referenciar y acceder a la variable original atraves de su direccion de memoria

	var x int = 10
	var p *int = &x //con el & mas el nombre de la variable le asignamos la direccion de memoria al puntero
	//implementacion de funcion con puntero
	editar(&x)
	fmt.Println(x, p)
}

// LA VENTAJA DE UTILIZAR PUNTEROS ES QUE AL ACCEDER A LA MEMORIA SE OPTIMIZA ESPACIO EN MEMORIA
// ADEMAS CUANDO SE LLAMA A UNA VARIABLE PARA IMPRIMIRLA O REASIGANARLA ESTAMOS CREANDO UNA COPIA DE LA VARIABLE EN MEMORIA
// LOS PUNTEROS PERMITEN EDITAR DATOS DE UNA VARIABLE A TRAVES DE UNA FUNCION SIN HACER COPIAS DE LA MEMORIA Y SIN RETORNAR DATOS DESDE LA FUNCION
func editar(x *int) {
	*x = 20
}

package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	t := time.Now()
	hora := t.Hour()
	//sintaxis del if: if condicion (la condicion va sin parentesis){}
	if hora < 12 {
		fmt.Println("mañana")
	} else { //se peude poner else if{}
		fmt.Println("tarde")
	}

	//go permite declarar e inicializar variables en la condicion y asi mismo evual la condicion
	if t2 := time.Now(); t.Hour() < 12 {
		fmt.Println(t2)
	}

	//SWITCH CASE

	os := runtime.GOOS // devuelve el sitema operativo

	//se evaluara en que sistema operatovo esta corriendo go
	//no es necesario colocar la instruccion break
	switch os {
	case "windows":
		fmt.Println("go run in windows")
	case "linux":
		fmt.Println("go run in linux")
	default: //es lo que sucederia sino se cumpli ninguno de los casos
		fmt.Println("otro os")
	}

	//con switch tambien se peude declara y usar la variable al tiempo como en el condicional

	//CICLO FOR
	//ES LA UNICA ESTRUCTURA EN BUCLE QUE EXISTE

	//CICLO INFINITO EN FOR
	for {
		break
	}

	//CICLO CON CONDICION, ESTE REMPLAZARIA AL WHILE
	var i int
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	//CICLO FOR CONVENCIONAL
	for x := 1; x < 10; x++ {
		fmt.Println(x)
	}

	//BREAK Y CONTINUE
	//break se utilzia para salir de un buble antes de que la condicion sea alcanzada
	for y := 1; y < 10; y++ {
		fmt.Println(y)
		if y == 2 {
			break
		}
	}

	//continue se utiliza para saltar a la siguiente iteracion de un ciclo
	for z := 1; z < 10; z++ {

		if z == 2 {
			continue //cuando z sea 2 entonces se salta inmediatamente a la 3cera iteracion sin importar lo que este despues de la palabra continue
		}
		fmt.Println(z)
	}

}

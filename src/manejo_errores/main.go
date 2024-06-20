package main

import (
	"errors"
	"fmt"
	"strconv"
)

func divide(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		//para crear nuestras propias funciones se utiliza el paquete errors
		//otra opcion es utilizar el paquete fmt con la funcion fmt.Errorf() para asi darle formato al error
		return 0, errors.New("No se peude dividir por cero") //se usa apra crear un nuevo mensaje de error
	}
	return dividendo / divisor, nil
}

func main() {

	//el control de errores es fundamental para garantizar la robustes de la aplicacion

	//EN GO EL MANEJO DE ERRORES SE HACE USANDO UN IF

	srt := "123"

	num, err := strconv.Atoi(srt) // la funcion Atoi devuelve un entero y un error

	//se maneja usando un if
	//sino hay error la variable err llega nula osea nil y si hay un error va a devolver un mensaje de error

	if err != nil {
		fmt.Println("Error:", err) //los espacios despues de la coma tambien se le suman al mensaje
		return
	}

	fmt.Println("numero:", num)

	//ERRORES PERSONALIDOS
	result, error := divide(10, 0)

	if error != nil {
		fmt.Println("Error:", error)
		return
	}

	fmt.Println("resultado:", result)

}

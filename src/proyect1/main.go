package main

import (
	"bufio"
	"fmt"
	"os"
)

type Tarea struct {
	nombre     string
	desc       string
	completado bool
}

type ListaTareas struct {
	tareas []Tarea
}

// agregar tareas
func (l *ListaTareas) agregarTarea(t Tarea) {
	l.tareas = append(l.tareas, t)
}

// marcar como completada una tarea
func (l *ListaTareas) marcarCompletado(index int) {
	l.tareas[index].completado = true
}

// editar tarea
func (l *ListaTareas) editarTarea(index int, t Tarea) {
	l.tareas[index] = t
}

// eliminar tarea
func (l *ListaTareas) eliminarTarea(index int) {
	l.tareas = append(l.tareas[:index], l.tareas[index+1:]...)
}

func main() {

	//INSTANCIANDO UNA LISTA DE TAREAS
	lista := ListaTareas{}

	//INSTANCIA DE BUFIO PARA ENTRADA DE DATOS
	leer := bufio.NewReader(os.Stdin)

	for {
		var opcion int
		fmt.Println("seleccione una opcion: \n",
			"1. agregar tarea\n",
			"2. marcar tarea como completada\n",
			"3. editar tarea\n",
			"4. aliminar tarea\n",
			"5. salir\n")

		fmt.Print("ingrese la opcion:")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			var t Tarea
			fmt.Print("ingrese el nombre e la tarea: ")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Print("ingrese lad escripcion de la tarea: ")
			t.desc, _ = leer.ReadString('\n')
			lista.agregarTarea(t)
			fmt.Println("tarea agregada correctamente")
		case 2:
			var index int
			fmt.Print("ingrese el indice de la tarea que desea marcar como completada: ")
			fmt.Scanln(&index)
			lista.marcarCompletado(index)
			fmt.Println("tarea marcada completada correctamente")
		case 3:
			var index int
			var t Tarea
			fmt.Print("ingrese el indice de la tarea que desea actualziar: ")
			fmt.Scanln(&index)
			fmt.Print("ingrese el nombre e la tarea: ")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Print("ingrese lad escripcion de la tarea: ")
			t.desc, _ = leer.ReadString('\n')
			lista.editarTarea(index, t)
		case 4:
			var index int
			fmt.Print("ingrese el indice de la tarea que desea actualziar: ")
			fmt.Scanln(&index)
			lista.eliminarTarea(index)
			fmt.Print("tarea eliminada correctamente")
		case 5:
			fmt.Print("sañiendo del programa...")
			return
		default:
			fmt.Println("opcion invalida")
		}

		//LISTAR TAREAS
		fmt.Println("lista de tareas")
		fmt.Println("===================================================")
		for i, t := range lista.tareas {
			fmt.Printf("%d. %s - %s - Completado: %t\n", i, t.nombre, t.desc, t.completado)
		}
		fmt.Println("===================================================")
	}
}

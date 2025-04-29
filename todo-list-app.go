package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)


	type Todo struct {
		id string
		description string
		status string
		createdAt string
}

func createTodo(todos []Todo, description string, status string) Todo {
	return Todo{
		id: strconv.Itoa(len(todos) + 1),
		description: description,
		status: status,
		createdAt: time.Now().UTC().String(),
	}
}

func getTodos(todos []Todo) {
	for _, todo := range todos {
		fmt.Println("ID: %s, Descripción: %s, Estado: %s, Fecha de creación: %s\n", todo.id, todo.description, todo.status, todo.createdAt)
	}
}

func deleteTodo(todos []Todo, id string) []Todo {
	for i, todo := range todos {
		if todo.id == id {
			fmt.Printf("Eliminando la tarea con ID: %s\n", id)
			// :i es un slice que contiene todos los elementos antes del índice i
			// :i+1 es un slice que contiene todos los elementos después del índice i
			//cuando los agregamos juntos, con append, eliminamos el elemento en la posición i
			// y mantenemos el resto de los elementos en la lista
			return append(todos[:i], todos[i+1:]...)
		}
	}
	return todos
}


func updateTodo(todos []Todo, id string, description string, status string) []Todo {
	for i, todo := range todos {
		if todo.id == id {
			todos[i].description = description
			todos[i].status = status
			todos[i].createdAt = strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
			return todos
		}
	}
	return todos
}


func main() {

	todos := []Todo{}
	 scanner := bufio.NewScanner(os.Stdin)
	 fmt.Println("---------------------------------")
	 fmt.Println("Bienvenido a to lista de Tareas!!:")
	 fmt.Println("---------------------------------")
	 fmt.Println("Qué deseas hacer?")
	 fmt.Println("1. Crear una tarea")
	 fmt.Println("2. Ver tareas")
	 fmt.Println("3. Actualizar una tarea")
	 fmt.Println("4. Eliminar una tarea")
	 fmt.Println("5. Salir")



	for {
	 fmt.Printf("Selecciona una opción (1-5): ")
 	 scanner.Scan()
	 text  := scanner.Text()
		switch text {
		case "1":
			fmt.Println("Creando una tarea...")
			fmt.Printf("Ingresa la descripción de la tarea: ")
			scanner.Scan()
			description := scanner.Text()
			fmt.Printf("Ingresa el estado de la tarea: ")
			scanner.Scan()
			estado:= scanner.Text()
			newTodo := createTodo(todos, description, estado)
			todos = append(todos, newTodo)
			fmt.Println("Tarea creada con éxito!")
			break
		case "2":
			fmt.Println("Lista de tareas: ")
			if len(todos) == 0 {
				fmt.Println("No hay tareas disponibles." )
			} else {
				getTodos(todos)
			}
		case "3":
			fmt.Println("Actualizando una tarea...")
			fmt.Println("Ingresa el ID de la tarea a actualizar: ")
			getTodos(todos)
			scanner.Scan()
			id:= scanner.Text()

			if len(todos) == 0 {
				fmt.Println("No hay tareas disponibles." )
				fmt.Println("---------------------------------")

			} else {
				fmt.Printf("Ingresa la nueva descripción de la tarea: ")
				scanner.Scan()
				description:= scanner.Text()
				fmt.Printf("Ingresa el nuevo estado de la tarea: ")
				scanner.Scan()
				estado:= scanner.Text()
				todos = updateTodo(todos, id, description, estado)
					fmt.Println("Tarea actualizada con éxito!" )
	 				fmt.Println("---------------------------------")
			}
		case "4":

			if len(todos) == 0 {
				fmt.Println("No hay tareas disponibles.")
			    fmt.Println("---------------------------------")

			} else {
				fmt.Printf("Ingresa el ID de la tarea a eliminar: ")
				getTodos(todos)
				scanner.Scan()
				id:= scanner.Text()
				todos = deleteTodo(todos, id)
				fmt.Println("Eliminando una tarea...")
				fmt.Println("Tarea eliminada con éxito!")
	      		fmt.Println("---------------------------------")
				getTodos(todos)


			}
		case "5":
				fmt.Println("Gracias por usar la app")
				return

		}
	}
}












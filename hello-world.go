package main

import (
	"fmt"
	"slices"
)




func main() {

	fmt.Println("Hello, World!")

	//tipos de datos:

	var nombre string = "Juan";
	var edad int = 25;
	var altura float32 = 1.75;
	var esEstudiante bool = true;


	//imprimir variables:
	fmt.Println("Nombre:", nombre)
	fmt.Println("Edad:", edad)
	fmt.Println("Altura:", altura)
	fmt.Println("Es estudiante:", esEstudiante)
	fmt.Println("Paso un año... :")

	edad = 28

	fmt.Println("Edad:", edad)

	experiencia :=  5 //este no es necesario tiparlo, ya que lo infiere el compilador
	fmt.Println("Experiencia adquirió: ", experiencia)



	fmt.Println("Altura:", altura)
	fmt.Println("Es estudiante:", esEstudiante)

	//arreglos:
	var numeros [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Arreglo de numeros:", numeros)

	var stringArray [3]string = [3]string{"uno", "dos", "tres"}
	fmt.Println("Arreglo de strings:", stringArray)


	var objetos [3]map[string]int = [3]map[string]int{
		{"uno": 1, "dos": 2},
		{"tres": 3, "cuatro": 4},
		{"cinco": 5, "seis": 6}}
	fmt.Println("Arreglo de objetos:", objetos)

	//mapas:
	var estudiantes map[string]int = map[string]int{"Juan": 25, "Maria": 30}
	fmt.Println("Mapa de estudiantes:", estudiantes)

	//map con slices indefinido  que contiene tiene un valor de tipo struct
	type Estudiante struct {
		nombre string
		edad int
	}



	//slices:
	var frutas  []string
	frutas = []string{"manzana", "banana", "naranja"}
	fmt.Println("Slice de frutas:", frutas)

	// dynamic slices
	frutas = append(frutas, "kiwi")
	fmt.Println("Añadiendo un elemento al final del Slice de frutas:", frutas)

	//Delete  elements from slice
	index := 1
	frutas = slices.Delete(frutas, index, index+1)
	fmt.Println("Eliminando el segundo elemento en Slice de frutas:", frutas)



	//for in go

	for i := 0; i < len(frutas); i++ {
		fmt.Println("Elemento en la posición", i, ":", frutas[i])
	}

	// for each
	for i, fruta := range frutas {
		fmt.Println("Elemento en la posición", i, ":", fruta)
	}

	// for each with condition
	for i, fruta := range frutas {
		if fruta == "banana" {
			fmt.Println("Elemento en la posición", i, "es banana")
		}
	}

	//while loop
	i := 0
	for i < len(frutas) {
		fmt.Println("Elemento en la posición", i, ":", frutas[i])
		i++
	}

	//do while loop
	i = 0
	for {
		fmt.Println("Elemento en la posición", i, ":", frutas[i])
		i++
		if i == len(frutas) {
			break
		}
	}






}


package main

func saludar(nombre string) string {
	return "Hola " + nombre
}

// funcion con retorno:
func sumar(a int, b int) int {
	return a + b
}
// funcion con retorno y sin retorno:
func sumarYMostrar(a int, b int) {
	resultado := a + b
	println("La suma es:", resultado)
}

//funcion que recibe un una funcion como parametro:

func operar(a int, b int, operacion func(int, int) int) int {
	return operacion(a, b)
}

//funcion que recibe un puntero

func modificarValor(valor *int) {
	*valor = *valor + 1
}



func createFunction() {

	//Definicion de  funciones en Go:

	//funcion sin retorno:
	println(saludar("Juan"))








}
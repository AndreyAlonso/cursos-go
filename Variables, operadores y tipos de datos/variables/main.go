package main

// Siempre se debe utilizar
import "fmt"

func main() {
	// forma 1
	// var dog string
	// dog = "perro"

	// forma 2
	// var dog, string = "perro"

	//forma 3
	var dog, cat string = "perro", "gato"

	//forma 4
	var perro, gato = "soy un perro", "soy un gato"

	// variable de asignacion
	nombre := "Andrey"

	// Multiples variables de asignacion
	// Como la variable face es nueva, no sucede un error al volver a asignar la variable cat
	cat, face := ":gato", ":face"

	// Si una variable  no se usa, entonces no se puede compilar
	fmt.Println(dog, cat)
	fmt.Println(perro, gato)
	fmt.Println("Hola soy ", nombre)
	fmt.Println(cat, face)
}

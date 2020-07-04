package main

import "fmt"

func main() {
	// Declaracion de constantes
	// Si no se declara como const, el default es var
	const pi = 3.14

	// pi = 4.15  -- NO PERMITE VOLVER A ASIGNAR UN CAMPO PORQUE ES CONSTANTE
	fmt.Println(pi)
}

package main

import "fmt"

func main() {
	// Operadores Aritméticos (), *, /, %, +, -
	var a = 4 + 2
	fmt.Println(a)

	// Operadores de asignación: =, +=, -=, *=, /=, %=
	var b = 10
	b += 2
	fmt.Println(b)

	// Declaración post-incremento y post-decremento: ++, --
	// (no son una expresión si no una declaración)
	var c = 3
	//c = c++ + 2 ERROR
	c++
	//fmt.Println(c++) ERROR
	fmt.Println(c)

	// Operadores de comparación: >, <, ==, !=, >=, <=
	fmt.Println(4 > 5) // true/false
	fmt.Println(4 == 4)

	// Operadores Lógicos &&, ||
	var edad = 30
	fmt.Println("edad:", edad >= 18 && edad <= 60)

	// Unario: !
	fmt.Println(!(4 == 4))
}

package main

import "fmt"

func main() {
	//bool, string, numeric
	var a bool = true
	//GO utiliza verbose
	//Para acceder al tipo utilizar %T	RESPETAR MAYUSCULA
	//Para acceder al valor utilizar %v RESPETAR MINUSCULA
	//Para agregar comillas al valor utilizar %q

	fmt.Printf("Tipo: %T, Valor: %v", a, a)

	var b string = "EDteam"
	fmt.Printf("\nTipo: %T, Valor: %v", b, b)

	//byte es igual a uint8
	//rune es igual a uint32
	var c uint8 = 100
	fmt.Printf("\nTipo: %T, Valor: %v", c, c)

	var d float64 = 23.56
	fmt.Printf("\nTipo: %T, Valor: %v", d, d)

	// Go. Al ser fuertemente tipado no permite hacer operaciones de diferentes tipos
	// Se requiere hacer el cast
	var e uint8 = 100
	var f uint16 = 23000
	g := uint16(e) + f // Sin el cast, sucede un error
	fmt.Printf("\nTipo: %T, Valor: %v", g, g)

	// Todos los tipos de datos que se declaran,GO les asigna un valor

	// Default de string, ""
	var cadena string
	fmt.Printf("\nDefault: %q", cadena)

	//Default de uint, 0
	var numero uint
	fmt.Printf("\nDefault: %v", numero)

	//Default de bool, false
	var bandera bool
	fmt.Printf("\nDefault: %v", bandera)
}

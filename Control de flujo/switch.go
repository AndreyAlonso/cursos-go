package main

import "fmt"

func main(){
	animal := "gato"
	// Forma 1
	switch animal {
	case "gato":
		fmt.Println("Es un gato")
	case "perro":
		fmt.Println("Es un perro")
	case "ave":
		fmt.Println("Es un ave")
	default:
		fmt.Println("Es un: ", animal)
	}

	vehiculo := "tren"

	// Forma 2
	switch vehiculo {
	case "auto","tren":
		fmt.Println(vehiculo, " es un vehiculo terrestre")
	case "avion", "helicoptero":
		fmt.Println(vehiculo, " es un vehiculo aereo")
	}

	// Forma 3
	switch {
	case vehiculo == "auto" || vehiculo == "tren":
		fmt.Println("Vehiculo terrestre")
	case vehiculo == "avion" || vehiculo == "helicoptero":
		fmt.Println("Vehiculo aereo")

	}
}

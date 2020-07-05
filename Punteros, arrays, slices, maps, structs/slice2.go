package main

import "fmt"

func main() {
	// len(): Numero de elementos en el slice
	// cap(): Numero de elementos del array donde apunta el slice, a partir del indice
	//de donde se creo el slice

	food := [5]string{"hot_dog", "fresa", "limon", "hamburguesa", "pizza"}
	fruits := food[1:3]

	// Se reemplaza el valor de la posición 3 de food
	fruits = append(fruits, "piña")

	// en food, solo cambia hasta donde fruits tiene acceso
	fruits = append(fruits, "manzana", "pera", "uva")

	fmt.Println("Food: ", food)
	fmt.Println("Fruits: ", fruits)

	fmt.Println("Fruits tamaño: ", len(fruits))
	fmt.Println("Food capacidad: ", cap(fruits)) //toma el cap del arreglo al que hace referencia

	// Otra forma de crear slice
	frutas := make([]string,0,3) //(tipo, inicio,fin)
	fmt.Println("\n\nFrutas: ", frutas)
	fmt.Println("Frutas tamaño: ", len(frutas))
	fmt.Println("Frutas capacidad: ", cap(frutas))
	fmt.Println(frutas == nil)
}
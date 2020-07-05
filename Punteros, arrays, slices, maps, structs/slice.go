package main

import "fmt"

func main(){
	// Slice no poseen datos, son apuntadores a un Array.
	set := [7]string{"leon","caballo", "perro", "gato", "ave", "avion", "barco"}

	animales := set[0:4] // (leon ... gato) no incluye el ave
	vehiculos := set[5:7] // (avion .. barco)
	allAnimals := set[:5] // si no se indica el inicio, default es cero [:4]

	//Se modifica el valor en el array "set"
	vehiculos[0] = "Tren"

	fmt.Println("Array completo:",set)
	fmt.Println("Animales:",animales)
	fmt.Println("Todos los animales: ", allAnimals)
	fmt.Println("Vehiculos: ",vehiculos)
}
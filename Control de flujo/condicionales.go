package main

import "fmt"

func main(){
	animal := "perro"

	if(animal == "gato") {
		fmt.Println("Es un gato")
	} else if(animal == "perro"){
		fmt.Println("Es un perro")
	}else{
		fmt.Printf("El animal es: %s", animal)
	}
	if vehiculo := "auto"; vehiculo == "tren"{
		fmt.Println("Es un tren")
	}else if vehiculo == "auto"{
		fmt.Println("Es un auto")
	}else{
		fmt.Printf("El vehiculo es: %s",vehiculo)
	}


}
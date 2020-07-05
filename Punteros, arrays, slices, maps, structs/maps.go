package main

import "fmt"

func main(){
	animales := make(map[string]string)//tipo de dato llave, tipo de dato a almacenar
	animales["animal01"] = "gato"
	animales["animal02"] = "perro"

	fmt.Println(animales)

	frutas := map[string]string {//tipo de dato llave, tipo de dato a almacenar
		"fruta01": "manzana",
		"fruta02": "banana", //Al final siempre una coma
	}
	fmt.Println(frutas)

	//Eliminar un elemento de un map
	delete(frutas,"fruta02")
	fmt.Println(frutas)

	//buscar un elemento de un mal
	fmt.Println(animales["animal02"])

	//Si no existe el key, devuelve string vacio
	fmt.Printf("%q",animales["no-existe"])

	//identificado blank, no se usa el primer campo
	if _, ok := animales["nuevoAnimal"]; !ok{
		animales["nuevoAnimal"] = "gorila"
	}
	fmt.Println("\n",animales)

}

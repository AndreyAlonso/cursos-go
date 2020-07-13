package main

import (
	"fmt"
	"os"
)

func main() {
	// Se guardan en una pila al utilizar defer
	defer fmt.Println(3)
	defer fmt.Println(1)
	defer fmt.Println(2)

	a := 5
	defer fmt.Println(" Defer: ", a)

	a = 10
	fmt.Println(a)
	//Utilizar defer despues de crear el archivo

	file, err := os.Create("hello.txt")
	if err != nil {
		fmt.Printf("Ocurrió un error al crear el archivo: %v", err)
		return
	}
	defer file.Close() //se ejecuta hasta que suceda un return

	_, err = file.Write([]byte("Hello Gophers"))
	if err != nil {
		fmt.Printf("Ocurrió un error al escribir el archivo: %v", err)
		return
	}

}

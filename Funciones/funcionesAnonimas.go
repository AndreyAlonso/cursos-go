package main

import "fmt"

func main(){
	x := func(){ // Las funciones anonimas no tienen nombre
		fmt.Println("Hola a todos")
	}
	x() // se manda llamar a la funcion anonima

	y := func(text string){
		fmt.Println("Hola",text)
	}
	y("Andrey")
}
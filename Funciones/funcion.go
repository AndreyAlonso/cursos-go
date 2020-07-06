package main

import (
	"fmt"
	"strings"
)

func main(){
	hello()
	nombre := "Andrey"

	saluda(nombre, "Hernandez")
	change(nombre)// No realiza el cambio porque es parametro por valor
	saluda(nombre, "Hernandez")
	changeReferencia(&nombre) // Se realiza el cabio, parametro por referencia
	saluda(nombre, "Hernandez")

	fmt.Println("\n",sum(2,4)) // retorno de un valor

	texto := "mEnSaJe"
	minusc, mayus := convert(texto)
	fmt.Println(minusc,mayus)

}

func hello(){
	fmt.Println("Hola EDteam")
}

func saluda(firsName string, lastName string){
	fmt.Printf("\nHola %s %s", firsName, lastName)
}

func change(value string) {
	value = "Hector"
}

func changeReferencia(value *string){
	*value = "Hector"
}

// se especifica que retorna un valor entero
func sum(num1 int, num2 int) int{
	return num1 + num2
}
// cuando la funcion retorna mas de 1 valor
// se especifia entre parentesis lo que regresa
func convert(text string) (string,string) {
	minuscula := strings.ToLower(text)
	mayuscula := strings.ToUpper(text)

	return minuscula, mayuscula
}
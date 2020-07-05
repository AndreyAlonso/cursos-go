package main

import "fmt"

func main(){
	fruit := "manzana"
	//Declaración de un puntero, se utiliza *
	var p *string
	p = &fruit //se asigna a p la direccíón de fruit

	// con & se obtiene la dirección de memoria de la variable
	fmt.Printf("Tipo: %T, Valor: %s, Drección: %v\n", fruit,fruit, &fruit)
	fmt.Printf("Tipo: %T, Valor: %v, Desreferenciación: %s\n", p,p,*p)

	//A donde apunta p, se cambia su valor
	*p = "piña"
	fmt.Println("Cambio de valor desde apuntador\n")
	fmt.Printf("Tipo: %T, Valor: %s, Drección: %v\n", fruit,fruit, &fruit)
	fmt.Printf("Tipo: %T, Valor: %v, Desreferenciación: %s\n", p,p,*p)



}
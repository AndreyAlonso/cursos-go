package main

import "fmt"

func main(){
	numeros := []int{1,4,70,5,67,90,2} //slice
	result := filter(numeros,func(num int) bool{
		if num <= 10 {
			return  true
		}
		return false
	})
	fmt.Println(result)
	x := hello("Andrey ")("Hernadez") // Se ejecutan las 2 funciones
	fmt.Println(x)
}
// Ejemplo funcion recibiendo una funcion
func filter(nums []int, callback func(int) bool ) []int{
	result := []int{}
	for _, v := range nums{
		if callback(v) { //true
			result = append(result,v)
		}
	}
	return result
}

// Ejemplo funcion retornando una funcion
func hello(name string) func(string) string{
	return func(lastname string) string{
		return "Hola " + name + lastname
	}
}
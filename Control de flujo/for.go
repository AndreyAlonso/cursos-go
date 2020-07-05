package main

import "fmt"

func main(){

	// foma 1, for clasico
	for i := 1; i <= 10; i++{
		fmt.Printf("[%v]",i)
	}
	fmt.Println("\n")
	// forma 2, for continuo
	i := 1
	for i <= 10{
		fmt.Printf("[%v]",i)
		i++
	}

	// forma 3, for forever
	//i = 1
	//for {
	//	fmt.Println(i)
	//	i++
	//}
	fmt.Println("\n")

	// forma 4, for range slice
	nums := []uint8{2,4,6,8}
	//indice, valor
	for i, v := range nums{
		//v *= 2//esto cambia el valor de la variable v
		nums[i] *= 2 //esto cambia el valor dentro del arreglo
		fmt.Println("Indice:",i,"valor:",v)
	}
	fmt.Println(nums)

	//forma 5, for range maps
	sports := map[string]string{"01": "basketball", "02": "soccer"}

	for key, v := range sports{
		fmt.Println("key:",key,"valor:",v)
	}

	//forma 6, for range string
	hello := "hello"
	for _,v := range hello{
		fmt.Println(string(v)) // si no se usa el cast, muestra los bytes
	}


}

package main

import "fmt"

func main() {
	var food [3]string
	food[0] = "hamburguesa"
	food[1] = "pizza"
	food[2] = "hot dog"
	fmt.Println(food) // [hamburguesa pizza hot dog]


	comida := [3]string{"hamburguesa","pizza", "hot dog"}
	fmt.Println(comida)

}


package main

import "fmt"

func main(){
	var res = sum(1,2,3,4)
	fmt.Println(res)
	res = sum(4,4)
	fmt.Println(res)
}

// los tres puntos indican una cantidad indefiniada de variables
func sum(nums ...int) int{
	total := 0
	for _,v := range nums{
		total += v
	}
	return total
}

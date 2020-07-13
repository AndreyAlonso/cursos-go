package main

import "fmt"

func main() {
	divide(10, 2)
	divide(40, 3)
	divide(12, 0)
	divide(20, 4)
}

func divide(dividendo, divisor int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recuperandome del panic", r)
		}
	}()
	validarDivisor(divisor)
	fmt.Println(dividendo / divisor)
}

func validarDivisor(divisor int) {
	if divisor == 0 {
		panic("PANIC!!!")
	}
}

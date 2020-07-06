package main

import (
	"errors"
	"fmt"
)

func main(){
	// Contenido del archivo, error al abrir el archivo
	//content, err := ioutil.ReadFile("./archivo.txt")

	//if err != nil { // error diferente de su valor cero
	//	fmt.Printf("Ocurrio un error: %v",err)
	//	return
	//}
	//fmt.Println(string(content))

	result, err := division(10,2)
	if err != nil{
		fmt.Printf("Ocurrio un error: %v", err)
		return
	}
	fmt.Println(result)

}

// parametros nombrados
func division(dividendo, divisor int) (result int, err error) {
	if divisor == 0 {
		err = errors.New("No se puede dividir entre 0")
		return
	}
	result = dividendo/ divisor
	return //no es necesario especificar lo que retorna
	// si no se uso la variable, toma su valor cero
	//return dividendo/divisor,nil
}
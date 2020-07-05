package main

import "fmt"

func main(){
	//Declaracion de la estructura
	type course struct {
		Name string
		Professor string
		Country string
	}

	db := course{
		Name: "Bases de datos",
		Professor: "Alexys",
		Country: "Colombia",
	}

	// {Bases de datos Alexys Colombia}
	fmt.Print(db,"\n")
	//{Name:Bases de datos Professor:Alexys Country:Colombia}
	fmt.Printf("%+v\n",db)

	git := course{"Git", "Beto", "Bolivia"}
	css := course{Professor: "Alvaro"}//Cuando no se asigna un valor, se queda vacío o 0
	fmt.Printf("%+v\n",git)
	fmt.Printf("%+v\n",css)

	//Acceder a los campos de la estructura
	fmt.Println("Profesor de GIT: ", git.Professor)

	p := &db
	//Modificar un campo de la estructura
	//Forma 1
	(*p).Professor = "Andrey"
	fmt.Printf("%+v\n",p)
	//Forma 2
	p.Professor = "Hector"
	fmt.Printf("%+v\n",p)


}

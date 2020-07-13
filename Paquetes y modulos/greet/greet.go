package greet //siempre indicar al paquete que pertenece

var nombre = "Andrey" //variable a nivel de paquete

// GreetEnglish retorna saludo en inglés
func English() string {
	return "Hi" + nombre
}

// Italian retorna saludo en italiano
func Italian() string {
	return "Ciao " + nombre
}

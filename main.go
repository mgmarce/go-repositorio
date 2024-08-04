package main

import (
	"fmt"
	"github.com/mgmarce/go-repositorio/ejercicios"
	//"runtime"
	//"github.com/mgmarce/go-repositorio/variables"
)

func main(){
	/*estado, texto := variables.ConviertoaTexto(1101)
	fmt.Println(estado)
	fmt.Println(texto)*/

	/*if os := runtime.GOOS; os == "Linux." {
		fmt.Println("Este no es windows")
	}else{
		fmt.Println("Este es windows")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "darwin":
		fmt.Println("Esto es Darwin")
	default:
		fmt.Printf("%s \n", os)
	}*/

	numero, texto := ejercicios.ConvNumerico("1")
	fmt.Println(numero)
	fmt.Println(texto)
}


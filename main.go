package main

import (
	"fmt"

	"github.com/mgmarce/go-repositorio/variables"
)

func main(){
	estado, texto := variables.ConviertoaTexto(1101)
	fmt.Println(estado)
	fmt.Println(texto)
}


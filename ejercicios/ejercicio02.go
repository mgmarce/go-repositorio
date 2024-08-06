package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*var num int
var err error

func TablaMultiplicacion() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Ingrese numero 1: ")
	if scanner.Scan() {
		num, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingresado es incorrecto " + err.Error())
		}else{
			for i := 0; i <= 10; i++ {
				fmt.Println(i," x ",num," = ",(num*i))
			}
		}
	}
}*/

var numero int
var problema error
var texto string

func TabladeMultiplicar() string{
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Ingrese un numero: ")
		if scanner.Scan() {
			numero, problema = strconv.Atoi(scanner.Text())
			if problema != nil {
				continue
			} else {
				break
			}
		}
	}

	for i := 1; i <= 10; i++ {
		texto += fmt.Sprintf("%d x %d = %d \n", numero, i, i*numero)
	}
	
	return texto
}

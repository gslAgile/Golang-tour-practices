package main

import (
	"fmt"
)

type ErrNegativeSqrt float64


/* Variables de tipo funcion funcion*/
var goutln = fmt.Println
var gout = fmt.Printf
var soutln = fmt.Sprintf

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number:%v", float64(e))
}


func Sqrt(x float64) (z float64, err error) {
	
	/* Inicializacion de variables:*/
	z = 1		/* Raiz de inicio 1*/
	err = nil	/* Error inicia como nulo*/

	if x<0 {
		err = ErrNegativeSqrt(x)
	} else {
		/* Iteracion metodo raiz de Newton */
		for i:=0; i < 10; i++ {
			z = z - ((z*z - x)/(2*z))
		}
	}

	return z, err
}

func main() {
	goutln(Sqrt(2))
	goutln(Sqrt(-2))
}
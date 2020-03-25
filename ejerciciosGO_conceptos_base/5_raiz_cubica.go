package main

import "fmt"
import cpx "math/cmplx"

/* Variables de tipo funcion funcion*/
var goutln = fmt.Println
var gout = fmt.Printf

var alCubo = func(x complex128) complex128 {
		return cpx.Pow(x, complex128(3))
	}

var alCuadrado = func(x complex128) complex128 {
		return cpx.Pow(x, complex128(2))
	}


func Cbrt(x complex128) (z complex128) {

	/* Inicializacion de raiz */
	z = 1.0

	/* Algoritmo de raiz cubica de Newton*/
	for i:=0; i<10; i++ {
		//z = z - (((z*z*z)-x)/(3*z*z))
		z = z - ((alCubo(z)-x)/(3*alCuadrado(z)))	
	}
	
	return
}


func main() {	

	goutln(Cbrt(8 + 1i))

}
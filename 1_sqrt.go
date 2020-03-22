package main

import (
		"fmt"
		"math"
)

func Sqrt_newton(x float64, delta float64) (z float64, iteraciones int, err float64) {
	
	var z_prev float64 = 0
	var merror float64 = 1
	z = 1
	iteraciones = 0
	
	/* Iteracion metodo raiz de Newton */
	//for i:=1; i<10; i++ {
	for ; delta < merror; iteraciones++ {
		
		z_prev = z
		z = z - ((z*z - x)/(2*z))
		
		/* Calculo de delta */
		if iteraciones>0 {
			merror = math.Abs((z - z_prev))/math.Abs(z)
		}
		
	}

	return z, iteraciones, merror
}

func main() {
	
	/* Configuracion de error */
	delta := 0.00000000001
	
	/* Calculo de raices de 1 a 10*/
	for i:=1; i<10; i++ {

		/* Calculo de valores de retorno asociados a funcion raiz newton*/
		raiz, it, err := Sqrt_newton(float64(i),delta)
		fmt.Printf("Raiz(%d) \tCalculo raiz = %g \tRaiz real = %g \tIteraciones = %d \tError = %g\n", i, raiz, math.Sqrt(float64(i)), it, err)
	}

	//fmt.Println(Sqrt_newton(2, delta))
	//fmt.Println(math.Sqrt(2))
}
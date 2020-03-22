/*---------------------------------------------------------------------------------------------------------------*
 * Ejercicio: Bucles y funciones
 * Para jugar un poco con funciones y bucles, implementa la raíz cuadrada usando el método de Newton.
 *
 * En este caso, el método de Newton consiste en aproximar Sqrt(x), escogiendo un punto inicial z y luego repitiendo:
 *						z = z - [(z^2 - x)/(2*z)]
 *
 * Newton's method
 * Para empezar, repite el cálculo 10 veces y mira cuánto te aproximas al valor real para diferentes valores
 * (1, 2, 3, ...).
 *
 * Luego, cambia la condición del bucle para que se detenga cuando el valor no cambie (o cambie menos que un valor 
 * delta muy pequeño).
 * Comprueba si eso da más o menos iteraciones. Cuánto te aproximas a la función math.Sqrt?
 * 
 * Pista: para declarar e inicializar un valor real, utiliza un literal con sintaxis de número real o utiliza una
 * conversión:
 *
 *	z := float64(1)
 *	z := 1.0
 *---------------------------------------------------------------------------------------------------------------*/
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
/*---------------------------------------------------------------------------------------------------------------*
 * Ejercicio: Slice
 * Implementa Pic. Debería devolver un slice de longitud dy, donde cada elemento es a su vez un slice de dx elementos
 * de tipo entero sin signo de 8 bits. Cuando ejecutes el programa, mostrará tu imagen, interpretando los enteros
 * como un nivel de gris (bueno, de azul).
 *
 * La imagen se deja a tu elección. Hay funciones interesantes como x^y, (x+y)/2, y x*y.
 *
 * (Necesitarás un bucle para reservar memoria para cada []uint8 dentro de [][]uint8.)
 *
 * (Utiliza uint8(valorEntero) para hacer la conversión a uint8.)
 *---------------------------------------------------------------------------------------------------------------*/
package main

import (
	//"math"
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	
	/*Creacion de slice pic */
	pic := make([][]uint8, dy) /* asignacion de mem para el tipo de dato nivel 1. */
	
	/* Recorro slice pic, asignando mem a cada elemento.*/
	for i := range pic {
		pic[i] = make([]uint8, dx) /* asignacion de mem para el tipo de dato nivel 2. */
		    for j := range pic[i] {
			    pic[i][j] = uint8(i*j)
			    //pic[i][j] = uint8((i+j)/2)
			    //pic[i][j] = uint8(math.Pow(float64(i),float64(j)))
		    }
    }
    return pic	
}

func main() {
	pic.Show(Pic)
}
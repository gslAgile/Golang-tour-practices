/*---------------------------------------------------------------------------------------------------------------*
 * Implementa la función fibonacci que devuelve otra función (una clausura) que devuelva los números sucesivos 
 * de fibonacci.
 *---------------------------------------------------------------------------------------------------------------*/
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	/**/
	x, y := 0, 1
    	return func() (r int) {
        	r = x
        	x, y = y, x + y
        return 
    }
}

/* Variables funcion*/
var goutln = fmt.Println
var gout = fmt.Printf

func main() {
	f := fibonacci()
	
	for i := 0; i < 3; i++ {
		goutln(f())
	}
}

/*---------------------------------------------------------------------------------------------------------------*
 * Ejercicio: Maps
 * Implementa WordCount. Debería devolver un map de contadores para cada “palabra” en la cadena de caracteres s.
 * La función wc.Text correrá una batería de tests sobre la función proporcionada y dará un veredicto.
 *
 * Quizás encuentres útil la función strings.Fields.
 *---------------------------------------------------------------------------------------------------------------*/
package main

import (
	"fmt"
	st "strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	/* Extraccion de palabras*/
	palabras := st.Fields(s)
	
	//fmt.Println(palabras)
	//fmt.Println(palabras[0])
	
	/* Creacion de maps*/ 
	var my_maps = make (map[string]int)
	
	for i:=0; i<len(palabras); i++ {
		my_maps[palabras[i]] += 1
	}
	
	return my_maps//map[string]int{"x": 1}
}

func main() {
	wc.Test(WordCount)
	my_maps := make(map[string]int)
	my_maps = WordCount("Hola mundo Go !!! Bienvenidos al mundo de Golang de Go :B")
	
	fmt.Println(my_maps)
	
	fmt.Println("\nLa frase 'Go' se repitio: ", my_maps["Go"], "veces.")
}
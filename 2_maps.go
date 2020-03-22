package main

import (
	"fmt"
	"strings"
	//"tour/wc"
)

func WordCount(s string) map[string]int {
	/* Extraccion de palabras*/
	palabras := strings.Fields(s)
	
	//fmt.Println(palabras)
	//fmt.Println(palabras[0])
	
	/* Creacion de maps*/ 
	var my_maps = make (map[string]int)
	
	for i:=0; i<len(palabras); i++ {
		my_maps[palabras[i]]= i+1;
	}
	
	return my_maps//map[string]int{"x": 1}
}

func main() {
	//wc.Test(WordCount)
	my_maps := make(map[string]int)
	my_maps = WordCount("Hola mundo! Bienvenido al mundo Golang :B")
	
	fmt.Println(my_maps)
	
	fmt.Println("La palabra Golang es la palabra nº:", my_maps["Golang"])
}
package main

import (
	"golang.org/x/tour/tree"
	"fmt"
	//"runtime"
)

/* Variables de tipo funcion funcion*/
var goutln = fmt.Println
var gout = fmt.Printf
var newln = func () {
	fmt.Println("")
	}

// Walk recorre la hojas del árbol t 
// enviando los valores por el canal ch.
func Walk(t *tree.Tree, ch chan int) {
	
	var walker func(t *tree.Tree)
	walker = func (t *tree.Tree) {
        if (t == nil) {
            return
        }
        
        walker(t.Left)
        ch <- t.Value
        walker(t.Right)
    }

    walker(t)
    close(ch)
}

// Same determina si los árboles
// t1 y t2 contienen los mismos valores.
func Same(t1, t2 *tree.Tree) bool {
	/* Creo 2 canales de datos enteros*/
	ch1, ch2 := make(chan int), make(chan int)

	/* Ejecuto dos gorutinas para obtener los datos de cada btree, desde ch1 y ch2.*/
	go Walk(t1, ch1)
    go Walk(t2, ch2)

    /**/
    for {
        v1,ok1 := <- ch1
        v2,ok2 := <- ch2

        if v1 != v2 || ok1 != ok2 {
            return false
        }

        goutln(v1, "=", v2)

        if !ok1 {
            break
        }
    }

    return true
}

func main() {
	
	goutln("CASO 1: tree.New(1) con tree.New(1)")
	goutln("Resultado:", Same(tree.New(1), tree.New(1))) // debería devolver true
	newln()
	goutln("CASO 2: tree.New(1) con tree.New(2)")
	goutln("Resultado:", Same(tree.New(1), tree.New(2))) // debería devolver false.
}
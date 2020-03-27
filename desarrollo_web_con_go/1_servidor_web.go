package main

import (
	h "net/http"
	"fmt"
)

/* Variables de tipo funcion funcion*/
var Print_h1 = func(p_w h.ResponseWriter, x string) {
		fmt.Fprint(p_w, "<h1>", x, "</h1>")
	}

func main() {
	var st = "hola mundoo..."
	/* Crear una ruta (PATH) mediante la funcion HandleFunc del pck http. */
	h.HandleFunc("/",	/* Raiz de nuestro sitio. */
				func(w h.ResponseWriter,r *h.Request) { /* Manejador: funcion que se llama cuando entramos a nuestro sitio web.*/
					Print_h1(w,st)
					})

	/* Creacion de servidor web. Disponible en modo de Listen (escuchando)*/
	h.ListenAndServe(":8080", nil)	/*ip:localhost puerto 8080. */
}
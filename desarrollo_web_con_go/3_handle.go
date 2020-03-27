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
	var st = "ServeMux: hola mundo!"
	
	/* ServeMux: Enrutador de peticiones http. Funciona como Handle.*/
	/* Creaacion de serveMux*/
	mux := h.NewServeMux()

	/* Creacion de ruta (PATH) mediante la funcion HandleFunc del pck http. */
	mux.HandleFunc("/",	/* Raiz de nuestro sitio. */
				func(w h.ResponseWriter,r *h.Request) { /* Manejador: funcion que se llama cuando entramos a nuestro sitio web.*/
					Print_h1(w,st)
					})

	/* Creacion de mas rutas: */
	mux.HandleFunc("/prueba",	/* Raiz de nuestro sitio. */
				func(w h.ResponseWriter,r *h.Request) { /* Manejador: funcion que se llama cuando entramos a nuestro sitio web.*/
					Print_h1(w,"ServeMux desde /prueba: Hola mundo!")
					})

	mux.HandleFunc("/usuario",	/* Raiz de nuestro sitio. */
				func(w h.ResponseWriter,r *h.Request) { /* Manejador: funcion que se llama cuando entramos a nuestro sitio web.*/
					Print_h1(w,"Hola usuario.")
					})

	/* Creacion de servidor web. Disponible en modo de Listen (escuchando)*/
	h.ListenAndServe(":8080", mux)	/*ip:localhost puerto 8080. */
}
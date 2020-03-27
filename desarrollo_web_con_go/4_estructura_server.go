package main

import (
	h "net/http"
	"fmt"
	"time"
	"log"
)


/* Tipo de datos propios*/
type mensaje struct {
	c_msg string
}

/**/


/* Variables de tipo funcion funcion*/
var Print_h1 = func(p_w h.ResponseWriter, x string) {
		fmt.Fprint(p_w, "<h1>", x, "</h1>")
	}


/* Funciones*/
func(m mensaje) ServeHTTP (w h.ResponseWriter,r *h.Request) {
	Print_h1(w, m.c_msg)
}


func hello_raiz (w h.ResponseWriter,r *h.Request) { /* Manejador: funcion que se llama cuando entramos a nuestro sitio web.*/
	Print_h1(w,"ServeMux: hola mundo!")
}

func hello_prueba (w h.ResponseWriter,r *h.Request) { /* Manejador: funcion que se llama cuando entramos a nuestro sitio web.*/
	Print_h1(w,"ServeMux desde /prueba: Hola mundo!")
}

func hello_usuario (w h.ResponseWriter,r *h.Request) { /* Manejador: funcion que se llama cuando entramos a nuestro sitio web.*/
	Print_h1(w,"Hola usuario.")
}




func main() {
	
	msg := mensaje{
		c_msg: "Handle: Hola mundo!",
	}

	/* ServeMux: Enrutador de peticiones http. Funciona como Handle.*/
	/* Creaacion de serveMux*/
	mux := h.NewServeMux()

	/* Creacion de ruta (PATH) mediante la funcion HandleFunc del pck http. */
	mux.HandleFunc("/", hello_raiz)

	/* Creacion de mas rutas: */
	mux.HandleFunc("/prueba", hello_prueba)

	mux.HandleFunc("/usuario", hello_usuario)

	/* Creacion de ruta utilizando Handle*/
	mux.Handle("/prueba2", msg)

	/* Creacion de servidor web.*/
	server := &h.Server{
		Addr: 				"localhost:8080",
		Handler: 			mux,
		ReadTimeout: 		10* time.Second,
		WriteTimeout: 		10* time.Second,
		MaxHeaderBytes: 	1 << 20,
	}
	log.Println("Listening ...")
	log.Fatal(server.ListenAndServe())
	
}
package main 

import (
	"github.com/gorilla/mux"
	h "net/http"
	"fmt"
	"time"
	"log"
)

/* Variables de tipo funcion funcion*/
var Print_h1 = func(p_w h.ResponseWriter, x string) {
		fmt.Fprint(p_w, "<h1>", x, "</h1>")
	}


/* Funciones */
func GetUsers(w h.ResponseWriter,r *h.Request) {
	Print_h1(w, "Mensaje desde el metodo GET")
}


func PostUsers(w h.ResponseWriter,r *h.Request) {
	Print_h1(w, "Mensaje desde el metodo POST")
}


func PutUsers(w h.ResponseWriter,r *h.Request) {
	Print_h1(w, "Mensaje desde el metodo PUT")
}


func DeleteUsers(w h.ResponseWriter,r *h.Request) {
	Print_h1(w, "Mensaje desde el metodo DELETE")
}



func main() {
	
	/* Creacion de enrutador de paquetaes http*/
	/* StricSlash(false): solo acepta path sin / al final. Pe: user/go pero no usr/go/ .*/
	r := mux.NewRouter().StrictSlash(false) 

	r.HandleFunc("/api/users", GetUsers).Methods("GET")
	r.HandleFunc("/api/users", PostUsers).Methods("POST")
	r.HandleFunc("/api/users", PutUsers).Methods("PUT")
	r.HandleFunc("/api/users", DeleteUsers).Methods("DELETE")

	/* Creacion de servidor web.*/
	server := &h.Server{
		Addr: 				"localhost:8080",
		Handler: 			r,
		ReadTimeout: 		10* time.Second,
		WriteTimeout: 		10* time.Second,
		MaxHeaderBytes: 	1 << 20,
	}
	log.Println("Listening ...")
	log.Fatal(server.ListenAndServe())

}
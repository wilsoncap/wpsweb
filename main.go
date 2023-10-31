package main

import (
	"log"
	"net/http"
	"rpsweb/handlers"
)

func main() {

	//CREAR ENRUTADOR
	router := http.NewServeMux()

	//MANEJADOR PARA SERVIR ARCHIVOS ESTATICOS
	fs := http.FileServer(http.Dir("static"))

	//RUTA PARA ACCEDER A LOS ARCHIVOS ESTATICOS
	router.Handle("/static/", http.StripPrefix("/static/", fs))

	//CONFIGURAR RUTAS
	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/new", handlers.NewGame)
	router.HandleFunc("/game", handlers.Game)
	router.HandleFunc("/play", handlers.Play)
	router.HandleFunc("/about", handlers.About)

	port := ":8082"
	log.Printf("servidor escuchando en http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))

	//ROUTER:: dirige las solicitudes http a las funciones del handler correspondiente ej: "/", => func(w http.ResponseWriter, r *http.Request)

	//HANDLER:: funcion que se encarga de manejar una solicitud http
}

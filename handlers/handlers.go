package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"rpsweb/rps"
	"strconv"
)

type Player struct {
	Name string
}

var player Player

const (
	templateDir  = "templates/"
	templateBase = templateDir + "base.html"
)

func Index(w http.ResponseWriter, r *http.Request) {
	restarValue()
	renderTemplate(w, "index.html", nil)
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	restarValue()
	renderTemplate(w, "new-game.html", nil)
}

func Game(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}
		player.Name = r.Form.Get("name")
	}
	//Si el campo del nomre esta vacio redireccione nuevamente a ese formulario
	if player.Name == "" {
		http.Redirect(w, r, "/new", http.StatusFound)
	}
	renderTemplate(w, "game.html", player)
}

func Play(w http.ResponseWriter, r *http.Request) {
	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("c"))
	result := rps.PlayRound(playerChoice)

	out, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(out)

}

func About(w http.ResponseWriter, r *http.Request) {
	restarValue()
	renderTemplate(w, "about.html", nil)
}

func renderTemplate(w http.ResponseWriter, page string, data any) {
	// devuelve un template y si al momento de cargar Must detecta un error va entrar en un panico nustra aplicacion
	tpl := template.Must(template.ParseFiles(templateBase, templateDir+page))
	err := tpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, "error al renderizar la plantilla", http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

func restarValue() {
	player.Name = ""
	rps.ComputerScore = 0
	rps.PlayerScore = 0
}

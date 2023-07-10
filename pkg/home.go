package pkg

import (
	"fmt"
	"html/template"
	"net/http"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {

	err := template.Must(template.ParseFiles("templates/filter.html")).Execute(w, nil)
	if err != nil {
		fmt.Println("Erreur de l'execution du template", err)
	}
}

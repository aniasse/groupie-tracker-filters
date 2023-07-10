package pkg

import (
	"html/template"
	"net/http"
)

func GetStatus(w http.ResponseWriter, errpage string) {

	temp := template.Must(template.ParseFiles("./templates/" + errpage + ".html"))

	if errpage == "error400" {
		w.WriteHeader(http.StatusBadRequest)
	} else if errpage == "error404" {
		w.WriteHeader(http.StatusNotFound)
	} else if errpage == "error500" {
		w.WriteHeader(http.StatusInternalServerError)
	}

	temp.Execute(w, nil)
}

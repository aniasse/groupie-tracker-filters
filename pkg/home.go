package pkg

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {

	err := template.Must(template.ParseFiles("templates/home.html")).Execute(w, nil)
	if err != nil {
		fmt.Println("Erreur de l'execution du template", err)
	}
}

func HandleFilter(w http.ResponseWriter, r *http.Request) {

	Art := GetArtistData()

	creation_date := []string{}
	tabId := []int{}
	for _, v := range Art {
		if !NoRepeatLoc(creation_date, strconv.Itoa(v.Acread)) {
			creation_date = append(creation_date, strconv.Itoa(v.Acread))
			tabId = append(tabId, v.Aid)
		}
	}
	// fmt.Println(creation_date, tabId)

	location := GetLocationData()
	tabloc := []string{}

	for _, v := range location.Loc {
		for i := 0; i < len(v.Loca); i++ {
			if !NoRepeatLoc(tabloc, v.Loca[i]) {
				tabloc = append(tabloc, v.Loca[i])
			}
		}
	}
	for i := 0; i < len(tabloc); i++ {
		for j := i + 1; j < len(tabloc); j++ {
			if tabloc[i] > tabloc[j] {
				swap := tabloc[i]
				tabloc[i] = tabloc[j]
				tabloc[j] = swap
			}
		}
	}
	NewFilter := Filter{
		// Artists:   Art,
		Locations: tabloc,
	}

	// dat_debut := r.FormValue("datdebut")
	// dat_fin := r.FormValue("datfin")
	// debut_album := r.FormValue("debutalbum")
	// final_album := r.FormValue("finalalbum")
	// loc := r.FormValue("location")
	// member := r.FormValue("members")

	//fmt.Println(dat_debut, dat_fin, debut_album, final_album, member, loc)

	err := template.Must(template.ParseFiles("templates/home.html")).Execute(w, NewFilter)
	if err != nil {
		fmt.Println("Erreur de l'execution du template", err)
	}
}

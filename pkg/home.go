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
	album := []string{}
	//tabId := []int{}
	for _, v := range Art {
		album = append(album, v.Afalbum)
		//tabId = append(tabId, v.Aid)
		if !NoRepeatLoc(creation_date, strconv.Itoa(v.Acread)) {
			creation_date = append(creation_date, strconv.Itoa(v.Acread))
		}
	}

	//fmt.Println(creation_date, album, tabId)
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

	//Les champs coches par l'utiisateur
	// creat_date := r.FormValue("creationdate")
	// first_album := r.FormValue("firstalbum")
	// checkb_members := r.FormValue("members")
	// checkb_location := r.FormValue("location")

	// dat_debut := r.FormValue("datdebut")
	// dat_fin := r.FormValue("datfin")
	// debut_album := r.FormValue("debutalbum")
	// final_album := r.FormValue("finalalbum")
	// loc := r.FormValue("loc")
	// member := r.FormValue("member")
	// filters := make(map[string]string)

	// if Active(creat_date) {
	// 	filters["Acread"] = append(filters["Acread"], dat_debut)
	// }
	// if Active(first_album) {
	// 	filters["album1"] = debut_album
	// 	filters["album2"] = final_album
	// }
	// if Active(checkb_members) {
	// 	filters["member"] = member
	// }
	// if Active(checkb_location) {
	// 	filters["location"] = loc
	// }

	// for keyfilt, filt := range filters {
	// 	for _, art := range Art {

	// 	}
	// }

	// if Active(creat_date) && Active(first_album) && Active(checkb_members) && Active(checkb_location) {

	// 	//Les donnees choisit par l'utilisateurs
	// 	dat_debut := r.FormValue("datdebut")
	// 	dat_fin := r.FormValue("datfin")

	// 	debut_album := r.FormValue("debutalbum")
	// 	final_album := r.FormValue("finalalbum")
	// 	loc := r.FormValue("loc")
	// 	member := r.FormValue("member")
	// }

	// fmt.Println(creat_date, first_album, checkb_members, checkb_location)

	// fmt.Println(dat_debut, dat_fin, debut_album, final_album, member, loc)

	err := template.Must(template.ParseFiles("templates/home.html")).Execute(w, NewFilter)
	if err != nil {
		fmt.Println("Erreur de l'execution du template", err)
	}
}

func FilterDetail(Art []Artist, tab map[string]string) {

}

func Active(str string) bool {
	if str == "on" {
		return true
	}
	return false
}

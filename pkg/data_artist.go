package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type LocArtist struct {
	LAid  int      `json:"id"`
	LAloc []string `json:"locations"`
	LAdat string   `json:"dates"`
}
type DatArtist struct {
	DAid  int      `json:"id"`
	DAdat []string `json:"dates"`
}
type RelArtist struct {
	RAid     int                 `json:"id"`
	RAdatloc map[string][]string `json:"datesLocations"`
}

func Loc_Artist(api string) LocArtist {

	content_api, err := http.Get(api)
	if err != nil {
		fmt.Println("Erreur de recuperation des donnees", err)
	}
	defer content_api.Body.Close()
	scanner, er := ioutil.ReadAll(content_api.Body)
	if er != nil {
		fmt.Println("Erreur de lecture", er)
	}
	var ArtistLoc LocArtist
	erreur := json.Unmarshal([]byte(scanner), &ArtistLoc)
	if erreur != nil {
		fmt.Println("Erreur lors du decodage", erreur)
	}
	return ArtistLoc
}

func Dat_Artist(api string) DatArtist {

	content_api, err := http.Get(api)
	if err != nil {
		fmt.Println("Erreur de recuperation des donnees", err)
	}
	defer content_api.Body.Close()
	scanner, er := ioutil.ReadAll(content_api.Body)
	if er != nil {
		fmt.Println("Erreur de lecture", er)
	}
	var ArtistDat DatArtist
	erreur := json.Unmarshal([]byte(scanner), &ArtistDat)
	if erreur != nil {
		fmt.Println("Erreur lors du decodage", erreur)
	}
	return ArtistDat
}

func Rel_Artist(api string) RelArtist {

	content_api, err := http.Get(api)
	if err != nil {
		fmt.Println("Erreur de recuperation des donnees", err)
	}
	defer content_api.Body.Close()
	scanner, er := ioutil.ReadAll(content_api.Body)
	if er != nil {
		fmt.Println("Erreur de lecture", er)
	}
	var ArtistRel RelArtist
	erreur := json.Unmarshal([]byte(scanner), &ArtistRel)
	if erreur != nil {
		fmt.Println("Erreur lors du decodage", erreur)
	}
	return ArtistRel
}

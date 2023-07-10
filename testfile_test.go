package main

import (
	"fmt"
	"groupie-tracker/pkg"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleArtist(t *testing.T) {

	//Creation d'une requete factice(requete simule pour le test)
	req, err := http.NewRequest("GET", "/artists", nil)
	if err != nil {
		t.Fatal(err)
	}

	//Creation d'un enregistreur de reponse factice
	record := httptest.NewRecorder()

	//Appel de la fonction HandleArtist
	pkg.HandleArtist(record, req)

	//Verification du status de la reponse
	status := record.Code
	if status != http.StatusOK {
		t.Errorf("❌ Test Failed the good status is 200 not %d", status)
	} else {
		fmt.Println("✅ Test Succeeded")
	}
}

func TestHandleArtistDetail(t *testing.T, r *http.Request) {
	artistID := (r.URL.Query().Get("Id"))

	req, err := http.NewRequest("GET", "/artist-details?Id="+artistID, nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(req.Body)
}

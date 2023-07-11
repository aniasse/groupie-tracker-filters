package main

import (
	"fmt"
	"groupie-tracker-filters/pkg"
	"html/template"
	"net/http"
)

func main() {
	//Managing CSS and image files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/", errorHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			//pkg.HandleHome(w, r)
			pkg.HandleFilter(w, r)
		case "/artists":
			pkg.HandleArtist(w, r)
		case "/artist-details":
			pkg.HandleArtistDeatail(w, r)
		case "/locations":
			pkg.HandleLocation(w, r)
		case "/location-detail":
			pkg.HandleLocationDetail(w, r)
		case "/dates":
			pkg.HandleDAte(w, r)
		case "/date-infos":
			pkg.HandleDateInfo(w, r)
		default:
			error404Handler(w, r)
		}

	})))

	fmt.Println("The program is running on http://localhost:4040")
	http.ListenAndServe("localhost:4040", nil)

}

// error500Handler handles the HTTP response with a 500 error (Internal Server Error).
func error500Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	tmpl := template.Must(template.ParseFiles("templates/error500.html"))
	tmpl.Execute(w, nil)
}

// error404Handler handle the HTTP response with a 404 error (Not Found).
func error404Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	tmpl := template.Must(template.ParseFiles("templates/error404.html"))
	tmpl.Execute(w, nil)
}

// error400Handler handle the HTTP response with a 400 error (Bad Request).
func error400Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	tmpl := template.Must(template.ParseFiles("templates/error400.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		handleError(w, r, http.StatusInternalServerError)
		return
	}
}

// handleError handles HTTP errors and redirects to the appropriate error handlers.
func handleError(w http.ResponseWriter, r *http.Request, statusCode int) {
	switch statusCode {
	case http.StatusNotFound:
		error404Handler(w, r)
	case http.StatusInternalServerError:
		// Recover the server's internal error code
		errorCode := http.StatusInternalServerError

		// Your treatment according to the internal error code
		switch errorCode {
		case http.StatusInternalServerError:
			// Handles error 500 (Internal Server Error)
			error500Handler(w, r)
		default:
			// Handles other internal errors
			http.Error(w, "Autre erreur interne", errorCode)
		}
	case http.StatusBadRequest:
		error400Handler(w, r)
	default:
		// Handles the unexpected errors
		error500Handler(w, r)
	}
}

// errorHandler is a middleware application for handling errors when processing HTTP requests.
func errorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Erreur interne du serveur:", r)
				error500Handler(w, nil)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

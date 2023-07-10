package pkg

import (
	"fmt"
	"html/template"
	"net/http"
)

// error500Handler manage the HTTP Response with an error 500 (Internal Server Error).
func error500Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	tmpl := template.Must(template.ParseFiles("templates/error500.html"))
	tmpl.Execute(w, nil)
}

// error404Handler manage the HTTP Response with an error 404 (Not Found).
func error404Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	tmpl := template.Must(template.ParseFiles("templates/error404.html"))
	tmpl.Execute(w, nil)
}

// error400Handler manage the HTTP Response with an error 400 (Bad Request).
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
		// Retrieve the server's internal error code
		errorCode := http.StatusInternalServerError

		// Your treatment according to the internal error code
		switch errorCode {
		case http.StatusInternalServerError:
			// Manage the 500 error (Internal Server Error)
			error500Handler(w, r)
		default:
			// Manage the other errors internes
			http.Error(w, "Autre erreur interne", errorCode)
		}
	case http.StatusBadRequest:
		error400Handler(w, r)
	default:
		// Manage the unexpected error
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

package main

import (
	"fmt"
	"net/http"
)

// TODO
func (app *application) viewUserContentHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show the details of content %d\n", id)
}

// TODO
func (app *application) createUserContentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create user content")
}

// TODO
func (app *application) editUserContentHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "edit the details of user content %d\n", id)
}

// TODO
func (app *application) deleteUserContentHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "delete user content %d\n", id)
}

// TODO
func (app *application) listUserContentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "list user content")
}

package api

import (
	"fmt"
	"net/http"
)

func (app *application) viewUserContentHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show the details of content %d\n", id)
}

func (app *application) createUserContentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create user content")
}

func (app *application) editUserContentHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "edit the details of user content %d\n", id)
}

func (app *application) deleteUserContentHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "delete user content %d\n", id)
}

func (app *application) listUserContentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "list user content")
}

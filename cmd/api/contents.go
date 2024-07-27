package main

import (
	"fmt"
	"net/http"
)

// TODO
func (app *application) viewContentHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show the details of content %d\n", id)
}

// TODO
func (app *application) createContentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create content")
}

// TODO
func (app *application) editContentHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "edit the details of content %d\n", id)
}

// TODO
func (app *application) deleteContentHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "delete content %d\n", id)
}

// TODO
func (app *application) listContentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "list content")
}

// TODO
func (app *application) listGenresHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "list genres")
}

// TODO
func (app *application) listContentTypesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "list types")
}

// TODO
func (app *application) listContentStatusesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "list statuses")
}

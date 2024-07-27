package api

import (
	"fmt"
	"net/http"
)

func (app *application) viewContentHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show the details of content %d\n", id)
}

func (app *application) createContentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create content")
}

func (app *application) editContentHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "edit the details of content %d\n", id)
}

func (app *application) deleteContentHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "delete content %d\n", id)
}

func (app *application) listContentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "list content")
}

func (app *application) listGenresHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "list genres")
}

func (app *application) listContentTypesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "list types")
}

func (app *application) listContentStatusesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "list statuses")
}

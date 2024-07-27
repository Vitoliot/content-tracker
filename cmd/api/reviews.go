package main

import (
	"fmt"
	"net/http"
)

func (app *application) viewReviewHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show the details of review %d\n", id)
}

func (app *application) createReviewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create review")
}

func (app *application) editReviewHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "edit the details of review %d\n", id)
}

func (app *application) deleteReviewHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "delete review %d\n", id)
}

func (app *application) listReviewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "list review")
}

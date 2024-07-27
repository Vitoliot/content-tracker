package api

import (
	"fmt"
	"net/http"
)

func (app *application) viewImageHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show the details of image %d\n", id)
}

func (app *application) createImageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create image")
}

func (app *application) deleteImageHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "delete image %d\n", id)
}

func (app *application) listImageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "list image")
}

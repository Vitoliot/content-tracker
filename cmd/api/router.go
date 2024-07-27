package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)

	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/v1/contents", app.listContentHandler)
	router.HandlerFunc(http.MethodPost, "/v1/contents", app.createContentHandler)
	router.HandlerFunc(http.MethodGet, "/v1/contents/:id", app.viewContentHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/contents/:id", app.editContentHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/contents/:id", app.deleteContentHandler)

	router.HandlerFunc(http.MethodPost, "/v1/userContents", app.createUserContentHandler)
	router.HandlerFunc(http.MethodGet, "/v1/userContents/:id", app.viewUserContentHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/userContents/:id", app.editUserContentHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/userContents/:id", app.deleteUserContentHandler)

	router.HandlerFunc(http.MethodPost, "/v1/reviews", app.createReviewHandler)
	router.HandlerFunc(http.MethodGet, "/v1/reviews/:id", app.viewReviewHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/reviews/:id", app.editReviewHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/reviews/:id", app.deleteReviewHandler)

	router.HandlerFunc(http.MethodPost, "/v1/images", app.createImageHandler)
	router.HandlerFunc(http.MethodGet, "/v1/images/:id", app.viewImageHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/images/:id", app.deleteImageHandler)

	router.HandlerFunc(http.MethodGet, "/v1/genres", app.listGenresHandler)
	router.HandlerFunc(http.MethodGet, "/v1/types", app.listContentTypesHandler)
	router.HandlerFunc(http.MethodGet, "/v1/statuses", app.listContentStatusesHandler)

	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/activated", app.activateUserHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/password", app.updateUserPasswordHandler)

	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)
	router.HandlerFunc(http.MethodPost, "/v1/tokens/activation", app.createActivationTokenHandler)
	router.HandlerFunc(http.MethodPost, "/v1/tokens/password-reset", app.createPasswordResetTokenHandler)

	return router
}

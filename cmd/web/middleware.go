package main

import (
	"github.com/justinas/nosurf"
	"net/http"
)


// NoSurf adds CSRF protection to all post request
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// SessionLoad provides middleware which automatically loads and saves session data for the current request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
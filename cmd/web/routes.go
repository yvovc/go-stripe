package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Get("/virtual-terminal", app.VirtualTerminal)
	mux.Post("/payment-succeeded", app.PaymentSucceded)
	mux.Get("/charge-once", app.ChargeOnce)

	fileServer := http.FileServer(http.Dir("./internal/static"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}

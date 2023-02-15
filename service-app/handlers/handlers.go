package handlers

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"service-app/middlewares"
	"service-app/web"
)

func Api(log *log.Logger) http.Handler {

	//r := chi.NewRouter()
	//r.MethodFunc(http.MethodGet, "/check", check)
	app := web.App{
		Mux: chi.NewRouter(),
	}

	m := middlewares.NewMid(log)
	app.HandleFunc(http.MethodGet, "/check", m.Logger(m.Error(m.Panic(check))))

	return app
}

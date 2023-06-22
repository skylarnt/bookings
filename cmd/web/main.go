package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/skylarnt/bookings/pkg/config"
	"github.com/skylarnt/bookings/pkg/handlers"
	"github.com/skylarnt/bookings/pkg/render"
)

var sessionManager *scs.SessionManager

const portNumber = ":8080"

var App config.AppConfig

func main() {
	App.Production = false

	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	sessionManager.Cookie.Secure = App.Production

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	App.Session = sessionManager

	App.TemplateCache = tc
	App.UseCache = false

	repo := handlers.NewRepo(&App)
	handlers.NewHandlers(repo)
	render.NewTemplates(&App)

	fmt.Println(fmt.Sprintf("Starting the development server on port %s", portNumber))
	// _ = http.ListenAndServe(portNumber, nil)
	server := &http.Server{
		Addr:    portNumber,
		Handler: Routes(&App),
	}

	err = server.ListenAndServe()
	log.Fatal(err)
}

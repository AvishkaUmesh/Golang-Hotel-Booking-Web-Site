package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/AvishkaUmesh/Golang-Hotel-Booking-Web-Site/pkg/config"
	"github.com/AvishkaUmesh/Golang-Hotel-Booking-Web-Site/pkg/handlers"
	"github.com/AvishkaUmesh/Golang-Hotel-Booking-Web-Site/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const SERVER_ADDRESS = "localhost"
const PORT = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	// change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = templateCache
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	fmt.Println("Server is running on ", SERVER_ADDRESS+PORT)

	serv := &http.Server{
		Addr:    SERVER_ADDRESS + PORT,
		Handler: routes(&app),
	}

	err = serv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

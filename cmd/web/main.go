package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/AvishkaUmesh/Golang-Hotel-Booking-Web-Site/internal/config"
	"github.com/AvishkaUmesh/Golang-Hotel-Booking-Web-Site/internal/handlers"
	"github.com/AvishkaUmesh/Golang-Hotel-Booking-Web-Site/internal/helpers"
	"github.com/AvishkaUmesh/Golang-Hotel-Booking-Web-Site/internal/models"
	"github.com/AvishkaUmesh/Golang-Hotel-Booking-Web-Site/internal/render"
	"github.com/alexedwards/scs/v2"
)

const SERVER_ADDRESS = "localhost"
const PORT = ":8080"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {

	err := run()

	if err != nil {
		log.Fatal(err)
	}

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

func run() error {
	gob.Register(models.Reservation{})
	// change this to true when in production
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
		return err
	}

	app.TemplateCache = templateCache
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)
	helpers.NewHelpers(&app)

	return nil

}

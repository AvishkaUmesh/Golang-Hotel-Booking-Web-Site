package render

import (
	"encoding/gob"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/AvishkaUmesh/Golang-Hotel-Booking-Web-Site/internal/config"
	"github.com/AvishkaUmesh/Golang-Hotel-Booking-Web-Site/internal/models"
	"github.com/alexedwards/scs/v2"
)

var testApp config.AppConfig
var session *scs.SessionManager

func TestMain(m *testing.M) {

	gob.Register(models.Reservation{})
	// change this to true when in production
	testApp.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	testApp.Session = session

	app = &testApp

	os.Exit(m.Run())

}

type myWriter struct {
}

func (tw *myWriter) Header() http.Header {
	var h http.Header
	return h
}

func (tw *myWriter) WriteHeader(i int) {
}

func (tw *myWriter) Write(b []byte) (int, error) {
	length := len(b)
	return length, nil
}

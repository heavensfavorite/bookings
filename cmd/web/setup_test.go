package main

import (
	"encoding/gob"
	"net/http"
	"testing"
	"text/template"

	"github.com/alexedwards/scs/v2"
	"github.com/heavensfavorite/bookings/internal/models"
)

// use the app variable declared in main.go; do not redeclare it here
var testSession *scs.SessionManager
var pathToTemplates = "./../../templates"

var functions = template.FuncMap{
	"humanDate":  render.humanDate,
	"formatDate": render.formatDate,
	"iterate":    render.iterate,
	"add":        render.add,
}

func TestMain(m *testing.M) {
	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})
	gob.Register(map[string]int{})
}

type myHandler struct{}

func (mh *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

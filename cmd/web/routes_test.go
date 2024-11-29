package main

import (
	"fmt"
	"testing"

	"github.com/go-chi/chi"
	"github.com/heavensfavorite/bookings/internal/config"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	var _ *config.AppConfig = &app
	mux := routes()

	switch v := mux.(type) {
	case *chi.Mux:
		// do nothing
	default:
		t.Error(fmt.Sprintf("type is not *chi.Mux, but is %T", v))
	}
}

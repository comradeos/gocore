package main

import (
	"ApplicationABC/controllers"
	"testing"
)

func TestGeneralController(t *testing.T) {
	got := controllers.GeneralController()
	want := "GeneralController"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
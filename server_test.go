package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAvailableBikesNumber(t *testing.T) {
	t.Run("Returns 9 as available bikes number around Splio's HQ", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/fetch", nil)
		response := httptest.NewRecorder()

		VelibServer(response, request)

		got := response.Body.String()
		want := "9"

		if got != want {
			t.Error("got", got, "want", want)
		}
	})
}

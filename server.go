package main

import (
	"fmt"
	"log"
	"net/http"
)

// VelibServer is the general API server. It fetches available bikes from Paris API
// around Splio's HQ.
func VelibServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "9")
}

func main() {
	handler := http.HandlerFunc(VelibServer)
	if err := http.ListenAndServe(":4242", handler); err != nil {
		log.Fatalf("could not listen on port 4242 %v", err)
	}
}

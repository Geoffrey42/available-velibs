package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

var (
	mutex          sync.RWMutex
	availableBikes int32
)

type geofilter struct {
	lat      string
	lng      string
	distance string
}

func fetchAvailableVelibsEndlessly(geofilter geofilter) {

	geofilterValue := url.QueryEscape(geofilter.lat + ", " + geofilter.lng + ", " + geofilter.distance)
	openDataURL := "https://opendata.paris.fr/api/records/1.0/search/?dataset=velib-disponibilite-en-temps-reel&q=&geofilter.distance=" + geofilterValue

	for {
		mutex.Lock()
		response, err := http.Get(openDataURL)
		if err != nil {
			log.Fatalf("could not fetch data from opendata.paris: %v", err)
		}
		defer response.Body.Close()

		log.Println("Response status:", response.Status)

		buf := new(strings.Builder)
		_, err = io.Copy(buf, response.Body)
		if err != nil {
			log.Fatalf("could not convert response body to string: %v", err)
		}

		fmt.Println(buf.String())

		mutex.Unlock()

		time.Sleep(1 * time.Minute)
	}
}

func main() {

	splioHQ := geofilter{lat: "48.8819732984", lng: "2.30113215744", distance: "500"}

	go fetchAvailableVelibsEndlessly(splioHQ)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mutex.RLock()
		defer mutex.RUnlock()

		fmt.Fprint(w, availableBikes)
	})

	if err := http.ListenAndServe(":4242", handler); err != nil {
		log.Fatalf("could not listen on port 4242 %v", err)
	}
}

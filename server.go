package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"
)

var (
	mutex   sync.RWMutex
	results GlobalResponse
)

type geofilter struct {
	lat      string
	lng      string
	distance string
}

// Fields gathers data for a single station
// NumBikesAvailable: number of available bikes in this station
// CoordonneesGeo: station coordinates
type Fields struct {
	NumBikesAvailable int       `json:"numbikesavailable"`
	CoordonneesGeo    []float64 `json:"coordonnees_geo"`
}

// Records is a placeholder for station data
type Records struct {
	Fields Fields `json:"fields"`
}

// GlobalResponse gathers relevant data from opendata.paris
// Total : total number of available bikes
// NHits : number of stations in the area
// Records: results for each station
type GlobalResponse struct {
	Total   int
	NHits   int       `json:"nhits"`
	Records []Records `json:"records"`
}

// Sum get a sum of every NumBikesAvailable
// in a Records slice. See GlobalResponse type.
func (g *GlobalResponse) Sum() {
	sum := 0

	for _, v := range g.Records {
		sum += v.Fields.NumBikesAvailable
	}

	g.Total = sum
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

		err = json.NewDecoder(response.Body).Decode(&results)
		if err != nil {
			log.Fatalf("could not encode opendata response to json: %v", err)
		}

		results.Sum()

		mutex.Unlock()

		time.Sleep(1 * time.Minute)
	}
}

func main() {

	splioHQ := geofilter{lat: "48.8819732984", lng: "2.30113215744", distance: "500"}

	go fetchAvailableVelibsEndlessly(splioHQ)

	http.HandleFunc("/api/fetch", func(w http.ResponseWriter, r *http.Request) {
		mutex.RLock()
		defer mutex.RUnlock()

		var buffer bytes.Buffer
		json.NewEncoder(&buffer).Encode(&results)

		fmt.Fprint(w, buffer.String())
	})

	log.Fatal(http.ListenAndServe(":4242", nil))
}

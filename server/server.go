package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/rs/cors"
)

type geofilter struct {
	lat      string
	lng      string
	distance string
}

// Fields gathers data for a single station
// Name: Velib's station name
// NumBikesAvailable: number of available bikes in this station
// CoordonneesGeo: station coordinates
type Fields struct {
	Name              string    `json:"name"`
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
	Distance int       `json:distance`
	Total    int       `json:"total"`
	NHits    int       `json:"nhits"`
	Records  []Records `json:"records"`
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

func fetchAvailableVelibsEndlessly(c chan GlobalResponse) {

	searchArea := geofilter{lat: "48.8819732984", lng: "2.30113215744", distance: "500"}
	geofilterValue := url.QueryEscape(searchArea.lat + ", " + searchArea.lng + ", " + searchArea.distance)
	openDataURL := "https://opendata.paris.fr/api/records/1.0/search/?dataset=velib-disponibilite-en-temps-reel&q=&geofilter.distance=" + geofilterValue

	for {
		response, err := http.Get(openDataURL)
		if err != nil {
			log.Fatalf("could not fetch data from opendata.paris: %v", err)
		}
		defer response.Body.Close()

		log.Println("Response status:", response.Status)

		var results GlobalResponse
		err = json.NewDecoder(response.Body).Decode(&results)
		if err != nil {
			log.Fatalf("could not encode opendata response to json: %v", err)
		}

		results.Distance, err = strconv.Atoi(searchArea.distance)
		if err != nil {
			log.Fatalf("could not convert distance to int: %v", err)
		}

		results.Sum()

		c <- results

		time.Sleep(1 * time.Minute)
	}
}

func main() {

	c := make(chan GlobalResponse)

	go fetchAvailableVelibsEndlessly(c)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/fetch", func(w http.ResponseWriter, r *http.Request) {

		var buffer bytes.Buffer
		results := <-c
		json.NewEncoder(&buffer).Encode(&results)

		fmt.Fprint(w, buffer.String())
	})

	corsHandler := cors.Default().Handler(mux)
	log.Fatal(http.ListenAndServe(":4242", corsHandler))
}

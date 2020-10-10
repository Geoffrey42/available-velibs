package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var (
	mutex          sync.RWMutex
	availableBikes int32
)

func fetchAvailableVelibsEndlessly() {
	for {
		mutex.Lock()
		response, err := http.Get("https://opendata.paris.fr/api/records/1.0/search/?dataset=velib-disponibilite-en-temps-reel&q=&geofilter.distance=48.8819732984%2C+2.30113215744%2C+500")
		if err != nil {
			log.Fatalf("could not fetch data from opendata.paris: %v", err)
		}
		defer response.Body.Close()

		log.Println("Response status:", response.Status)

		scanner := bufio.NewScanner(response.Body)
		for i := 0; scanner.Scan() && i < 5; i++ {
			fmt.Println(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			panic(err)
		}

		mutex.Unlock()

		time.Sleep(1 * time.Minute)
	}
}

func main() {

	go fetchAvailableVelibsEndlessly()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mutex.RLock()
		defer mutex.RUnlock()

		fmt.Fprint(w, availableBikes)
	})

	if err := http.ListenAndServe(":4242", handler); err != nil {
		log.Fatalf("could not listen on port 4242 %v", err)
	}
}

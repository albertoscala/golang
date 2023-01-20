package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

func trainCompanyWeb(w http.ResponseWriter, r *http.Request) {
	file, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	line := TrainLine{}

	err = json.Unmarshal(file, &line)
	if err != nil {
		log.Fatal(err)
	}

	var totalDistance int
	var travelTime int

	for i := 0; i < len(line.Stations); i++ {
		totalDistance += line.Stations[i].Distance
		travelTime += line.Stations[i].ETA
	}

	stats := LineStats{TotalDistance: totalDistance, TravelTime: travelTime}

	stats.ArrivalTime = line.Departure.Add(time.Minute * time.Duration(travelTime))

	infos, err := json.MarshalIndent(stats, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/text")
	_, err = w.Write(infos)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/trainCompanyWeb", trainCompanyWeb)

	log.Fatal(http.ListenAndServe(":8090", nil))
}

package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type Station struct {
	Name     string `json:"name"`
	Distance int    `json:"distance"`
	ETA      int    `json:"ETA-from-previous-station"`
	Platform string `json:"platform"`
}

type TrainLine struct {
	Line      string    `json:"line"`
	Stations  []Station `json:"stations"`
	Departure time.Time `json:"departingTime"`
}

type LineStats struct {
	TotalDistance int       `json:"totalDistance"`
	TravelTime    int       `json:"travelTime"`
	ArrivalTime   time.Time `json:"arrivalTime"`
}

func trainCompanyWeb(w http.ResponseWriter, r *http.Request) {
	file, _ := io.ReadAll(r.Body)

	line := TrainLine{}

	_ = json.Unmarshal(file, &line)

	var totalDistance int
	var travelTime int

	for i := 0; i < len(line.Stations); i++ {
		totalDistance += line.Stations[i].Distance
		travelTime += line.Stations[i].ETA
	}

	stats := LineStats{TotalDistance: totalDistance, TravelTime: travelTime}

	stats.ArrivalTime = line.Departure.Add(time.Minute * time.Duration(travelTime))

	infos, _ := json.MarshalIndent(stats, "", " ")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/text")
	w.Write(infos)

	return
}

func main() {
	http.HandleFunc("/trainCompanyWeb", trainCompanyWeb)

	log.Fatal(http.ListenAndServe(":8090", nil))
}

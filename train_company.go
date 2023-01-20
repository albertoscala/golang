package main

import (
	"encoding/json"
	"log"
	"os"
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

func main() {
	file, err := os.ReadFile("trains.json")
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

	info, err := json.MarshalIndent(stats, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("infos.json", info, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

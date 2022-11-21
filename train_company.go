package main

import (
	"encoding/json"
	"io/ioutil"
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
	file, _ := ioutil.ReadFile("trains.json")

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

	info, _ := json.MarshalIndent(stats, "", " ")

	_ = ioutil.WriteFile("infos.json", info, 0644)
}

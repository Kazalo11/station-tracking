package database

import (
	"encoding/json"
	"log"
	"station-tracking/internal/models"
)

func GetAllStations() []models.Station {
	var stations []models.Station

	if err := json.Unmarshal([]byte(stations_data), &stations); err != nil {
		log.Fatal("Error unmarshalling JSON: ", err)
	}

	return stations
}

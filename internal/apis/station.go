package apis

import (
	"encoding/json"
	"net/http"

	"station-tracking/internal/database"
)

func GetAllStations(w http.ResponseWriter, r *http.Request) {
	stations := database.GetAllStations()

	json.NewEncoder(w).Encode(stations)
	(w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

}

package handlers

import (
    // "fmt"
    // "reflect"
    "net/http"
    // "database/sql"
    "encoding/json"
    "github.com/NirShihor/spotlas_api_golang4/db"
    "github.com/NirShihor/spotlas_api_golang4/models"
)

func SpotsInArea(w http.ResponseWriter, r *http.Request){


    db := db.SetupDB()

    latitude := r.URL.Query().Get("lat")
    longitude := r.URL.Query().Get("lon")
    dims := r.URL.Query().Get("dims")
    shape := r.URL.Query().Get("shape")

    if shape == "circle" {
        rows, err := db.Query(`SELECT "name", "website", "coordinates", "description", "rating"
        FROM geospots
        WHERE ST_DWithin("coordinates", ST_MakePoint( $1, $2)::geography, $3)
        ORDER BY
        CASE
        WHEN ST_DWithin("coordinates", ST_MakePoint($1, $2)::geography, 5000)
        THEN "coordinates"::geography <-> ST_MakePoint(0.0,0.0)::geography end,
        CASE
        WHEN ST_DWithin("coordinates", ST_MakePoint($1, $2)::geography, 50)
        THEN rating end;`, latitude, longitude, dims)
    
        var geospots []models.Geospot
    
        for rows.Next() {
            var ID string
            var Name string
            var Website string
            var Coordinates string
            var Description string
            var Rating float64
    
            err = rows.Scan(&Name, &Website, &Coordinates, &Description, &Rating)
    
            checkErr(err)
    
            geospots = append(geospots, models.Geospot{ID, Name, Website, Coordinates, Description, Rating})
            }
    
            var response = models.JsonResponse{Type: "success", Data: geospots}
    
            json.NewEncoder(w).Encode(response)
        } 
}
            

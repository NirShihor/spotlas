package handlers 

import (
    "fmt"
    "encoding/json"
    "net/http"
    
    _ "github.com/lib/pq"

    "github.com/NirShihor/spotlas_api_golang4/models"
    "github.com/NirShihor/spotlas_api_golang4/db"
)

func GetAllGeospots(w http.ResponseWriter, r *http.Request) {
    db := db.SetupDB()

    printMessage("Getting geospots...")

    // Get all movies from movies table that don't have movieID = "1"
    rows, err := db.Query(`SELECT "id", "name", COALESCE("website", ' '), "coordinates", COALESCE("description", ' '), "rating" FROM geospots`)

    // check errors
    checkErr(err)

    // var response []JsonResponse
    var geospots []models.Geospot
    // Foreach geospot
    for rows.Next() {
        var ID string
        var Name string
        var Website string
        var Coordinates string
        var Description string
        var Rating float64

        err = rows.Scan(&ID, &Name, &Website, &Coordinates, &Description, &Rating)

        // check errors
        checkErr(err)

        geospots = append(geospots, models.Geospot{ID, Name, Website, Coordinates, Description, Rating})
    }

    var response = models.JsonResponse{Type: "success", Data: geospots}

    json.NewEncoder(w).Encode(response)
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

// Function for handling messages
func printMessage(message string) {
    fmt.Println("")
    fmt.Println(message)
    fmt.Println("")
}
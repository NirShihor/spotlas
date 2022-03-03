package handlers

import (
    "fmt"
    "log"
    "database/sql"
    "net/http"
    "encoding/json"
    "net/url"
    "github.com/NirShihor/spotlas_api_golang4/models"
    "github.com/NirShihor/spotlas_api_golang4/db"
)

func WebsiteChange(w http.ResponseWriter, r *http.Request) {

    db := db.SetupDB()
    
    id := r.URL.Query().Get("id")

    row := db.QueryRow(`SELECT "id", "name", COALESCE("website", ' '), "coordinates", COALESCE("description", ' '), "rating" FROM geospots WHERE id = $1`, id)

    var gs models.Geospot



    switch err := row.Scan(&gs.ID, &gs.Name, &gs.Website, &gs.Coordinates, &gs.Description, &gs.Rating) ; err{
        case sql.ErrNoRows:
            fmt.Println("No rows were returned!")
        case nil:
            fmt.Println(gs.Website)
            str := gs.Website
            u, err := url.Parse(str)
            if err != nil {
                log.Fatal(err)
            }
            fmt.Println(u.Hostname())
            json.NewEncoder(w).Encode(u.Hostname())
        default:
            panic(err)
        }
    }

    





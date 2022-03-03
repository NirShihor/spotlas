package handlers

import (
    "fmt"
    "net/http"
    "log"
    "encoding/json"
    "strconv"

    "github.com/NirShihor/spotlas_api_golang4/db"

)

func CountDomain(w http.ResponseWriter, r *http.Request) {
    db := db.SetupDB()

    var count int
    
    row := db.QueryRow(`SELECT COUNT(*) FROM geospots GROUP BY website HAVING COUNT(*)> 1`)

    err := row.Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(count)
    json.NewEncoder(w).Encode("Spots that contain the same domain: " + strconv.Itoa(count))

}

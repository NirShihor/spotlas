package handlers

import (
    "fmt"
    "net/http"
    "log"
    "encoding/json"
    "strconv"

    "github.com/NirShihor/spotlas_api_golang4/db"

)

func DomainGreaterThan(w http.ResponseWriter, r *http.Request) {
    db := db.SetupDB()

    var count int
    var website string
    
    row := db.QueryRow(`SELECT COALESCE("website", ' '), COUNT(*) FROM geospots GROUP BY website HAVING COUNT(*)> 1`)

    err := row.Scan(&website, &count)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(count)
    json.NewEncoder(w).Encode("Spots that have a domain count greater than 1: " + strconv.Itoa(count))

}

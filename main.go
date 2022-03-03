package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/NirShihor/spotlas_api_golang4/db"
    "github.com/NirShihor/spotlas_api_golang4/handlers"
)

func printMessage(message string) {
    fmt.Println("")
    fmt.Println(message)
    fmt.Println("")
}

func main(){


    db.SetupDB()

    router := mux.NewRouter()

    router.HandleFunc("/geospots", handlers.GetAllGeospots).Methods("GET")
    router.HandleFunc("/geospot", handlers.WebsiteChange).Methods("GET")
    router.HandleFunc("/domains", handlers.CountDomain).Methods("GET")
    router.HandleFunc("/domain-count", handlers.DomainGreaterThan).Methods("GET")
    router.HandleFunc("/spots-area", handlers.SpotsInArea).Methods("GET")

    fmt.Println("Server at 8000")
    log.Fatal(http.ListenAndServe(":8000", router))
}
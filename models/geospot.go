package models

type Geospot struct {
    ID string `json:"id"`
    Name  string `json:"name"`
    Website string `json:"website"`
    Coordinates string `json:"coordinates"`
    Description string `json:"description"`
    Rating float64 `json:"rating"`
}

type JsonResponse struct {
    Type    string `json:"type"`
    Data    []Geospot `json:"data"`
    Message string `json:"message"`
}
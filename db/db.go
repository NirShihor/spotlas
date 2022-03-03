package db

import(
    "fmt"
    "database/sql"
    "github.com/joho/godotenv"
    "log"
    "os"
    _ "github.com/lib/pq"

)

const (
    DB_USER     = "postgres"
    DB_NAME     = "postgis_test"
)

// DB set up
func SetupDB() *sql.DB {

    password := os.Getenv("DB_PASSWORD")

    err := godotenv.Load()
    if err != nil {
      log.Fatal("Error loading .env file")
    }
    
    dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, password, DB_NAME)
    db, err := sql.Open("postgres", dbinfo)

    checkErr(err)

    fmt.Println("Connected to database")

    return db
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
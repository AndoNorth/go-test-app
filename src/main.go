package main

import (
  "log"
  "database/sql"
  _"github.com/lib/pq"
  "net/http"
)

var db *sql.DB

func main() {

  var err error
  db, err = sql.Open("postgres", "host= port=5432 user=user password=password dbname=mydb sslmode=disable")
  if err != nil {
    log.Fatalf("unable to connect to db\n")
  }
  defer db.Close()

  http.HandleFunc("/health", HealthCheckEndpoint)
  http.HandleFunc("/insert", InsertSampleRecordEndpoint)

  log.Println("starting server on :8080\n")
  http.ListenAndServe(":8080", nil)
}

func HealthCheckEndpoint(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusOK)
  w.Write([]byte("OK"))
}

func InsertSampleRecordEndpoint(w http.ResponseWriter, r *http.Request) {
  _, err := db.Exec("INSERT INTO sample_table (name) VALUES ($1)", "Sample Name")
  if err != nil {
    log.Println("failed to insert record\n")
  }
  w.WriteHeader(http.StatusOK)
  w.Write([]byte("inserted successfully"))
}
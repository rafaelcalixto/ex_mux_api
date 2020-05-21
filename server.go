package main

import(
  // Core libraries
  "fmt"
  "net/http"
  "log"
  // Framework
  "github.com/gorilla/mux"
  // Proprietary libraries
  api "endpoint"
)

func main() {
    // Init Router
    router := mux.NewRouter()

    // Route Handlers / Endpoints
    router.HandleFunc("/", api.IndexHandler)
    router.HandleFunc("/send", api.ReceiveFiles).Methods("POST")

    fmt.Println("Starting API")
    log.Fatal(http.ListenAndServe(":8080", router))
}

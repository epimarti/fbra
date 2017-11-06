package server

import (
  "log"
  "fmt"
  "encoding/json"
  "net/http"
  "github.com/gorilla/mux"

  "github.com/epimarti/fbra/fizzbuzz"
)

type Params struct {
  Int1 int
  Int2 int
  Limit int
  String1 string
  String2 string
}

func fizzBuzzHandler(w http.ResponseWriter, r *http.Request) {
  dec := json.NewDecoder(r.Body)
  var p Params

  if err := dec.Decode(&p); err != nil {
    log.Fatal(err)
  }
  fmt.Println(p.Int1, p.Int2, p.Limit, p.String1, p.String2)

  json.NewEncoder(w).Encode(fizzbuzz.Generate(p.Int1, p.Int2, p.Limit, p.String1, p.String2))
}

func Start(port int) {
  router := mux.NewRouter()
  router.HandleFunc("/", fizzBuzzHandler).Methods("PUT")
  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}

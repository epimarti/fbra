package server

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"

	"github.com/epimarti/fbra/fizzbuzz"
)

type Params struct {
	Int1    int
	Int2    int
	Limit   int
	String1 string
	String2 string
}

func getParams(vars map[string]string) (Params, error) {
	var p Params

	int1, err := strconv.Atoi(vars["int1"])
	if err != nil {
		return p, err
	}

	int2, err := strconv.Atoi(vars["int2"])
	if err != nil {
		return p, err
	}

	limit, err := strconv.Atoi(vars["limit"])
	if err != nil {
		return p, err
	}

	p.Int1 = int1
	p.Int2 = int2
	p.Limit = limit
	p.String1 = vars["string1"]
	p.String2 = vars["string2"]

	return p, nil
}

func fizzBuzzHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	p, err := getParams(vars)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(fizzbuzz.Generate(p.Int1, p.Int2, p.Limit, p.String1, p.String2))
}

func Start() {
	router := mux.NewRouter()
	router.HandleFunc("/{int1:[0-9]+}/{int2:[0-9]+}/{limit:[0-9]+}/{string1}/{string2}", fizzBuzzHandler).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe("", router))
}

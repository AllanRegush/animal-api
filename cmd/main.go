package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/allanregush/animal-api/lib"
)

var data [256]Animal

type Animal struct {
	Id   int    `json: "id"`
	Kind string `json: "type"`
	Name string `json: "name"`
	Age  int    `json: "age"`
}

func home(w http.ResponseWriter, r *http.Request) {
	value := r.URL.Query().Get("q")
	result := make([]Animal, 0, 5)
	for _, v := range data {
		if strings.Contains(strings.ToLower(v.Kind), strings.ToLower(value)) {
			result = append(result, v)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func initData() {
	chance := lib.Chance{}
	for i := 0; i < len(data); i++ {
		data[i] = Animal{
			Id:   i,
			Kind: chance.Animal(),
			Name: chance.Name(),
			Age:  chance.Age(),
		}
	}
}

func main() {
	initData()
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	log.Println("Serving on 8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err.Error())
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// type seting struct{}

// Settings ::
type Settings struct {
	Route    string `json:"route"`
	Method   string `json:"method"`
	Request  string `json:"request"`
	Response string `json:"response"`
	Param    string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func shopHandler(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// method, _ := r.Method
	fmt.Println(r.RequestURI, r.Method, mux.Vars(r))
	fmt.Fprint(w, `[{"shname":"BTC","loname":"Bitcoin"},{"shname":"ETH","loname":"Ethereum"},{"shname":"BCH","loname":"Bitcoin Cash"},{"shname":"XRP","loname":"Ripple"}]`)
}

var data []Settings

func init() {
	file, err := ioutil.ReadFile("test.json")
	if err != nil {
		fmt.Println()
	}

	json.Unmarshal([]byte(file), &data)
}

// ServeHTTP ::
func (h *Settings) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(h.Param))
}

func main() {
	router := mux.NewRouter()

	for _, s := range data {
		router.Handle(s.Route, &Settings{Param: s.Response}).Methods(s.Method)
	}

	router.HandleFunc("/", indexHandler)
	handler := cors.Default().Handler(router)
	http.Handle("/", handler)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}

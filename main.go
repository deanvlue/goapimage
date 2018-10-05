package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// restarted

func init() {
}

var router = mux.NewRouter()

func main() {
	router.Path("/api/namedgoldcard/").Queries("name", "{name}").HandlerFunc(NamedGoldCardHandler).Name("NamedGoldCardHandler")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// NamedGoldCardHandler handles the input request
func NamedGoldCardHandler(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")

	_, err := router.Get("NamedGoldCardHandler").URL("name", name)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintln(w, "Hello,", name)
}

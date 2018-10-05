package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// restarted

func init() {
}

var router = mux.NewRouter()

func main() {
	router.Path("/api/namedgoldcard/").Queries("name", "{name}", "code", "{code}").HandlerFunc(NamedGoldCardHandler).Name("NamedGoldCardHandler")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// NamedGoldCardHandler handles the input request
func NamedGoldCardHandler(w http.ResponseWriter, r *http.Request) {

	authCode := os.Getenv("authCode")
	name := r.URL.Query().Get("name")
	code := r.URL.Query().Get("code")

	if code == authCode {
		log.Println("name:", name)
		log.Println("code:", code)
		fmt.Fprintln(w, "Name:", name)
		fmt.Fprintln(w, "Code:", code)
	} else {
		http.Error(w, "Unauthorized", 401)
	}

}

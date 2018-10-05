package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// restarted

type helloWorldResponse struct {
	Message string `json:"message"`
	Author  string `json:"-"`
	// do not output the fuel if the value is empty
	Date string `json:",omitempty"`
	//convert output to a string and rename to id
	Id int `json:"id, string"`
}

func init() {
}

func main() {
	// implementar llamada al api
	port := 8080

	http.HandleFunc("/holaMundo", helloWorldHandler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))

	// cargar imagen y fuentes de resources
	// desplegar imagen en la llamada al api en base64
	// tomar una variable del query que se manda al API
	// separar esa cadena por espacios y poner los índices 0 y 1
	// escribir la cadena resultante en la imagen
	//

}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "Hola Mundo\n")
	response := helloWorldResponse{Message: "Hello World"}
	data, err := json.Marshal(response)
	if err != nil {
		panic("Oppppsspososos")
	}

	fmt.Fprint(w, string(data))
}

/*func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
*/

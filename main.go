package main

import (
	"bufio"
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
	imgFile, err := os.Open("./resources/goldcard.jpg") // gold card

	if err != nil {
		fmt.Println(err)
		log.Println(err)
		os.Exit(1)
	}

	defer imgFile.Close()

	//creamos un buffer del tamaño de la imagen
	fInfo, _ := imgFile.Stat()
	var size = fInfo.Size()
	buf := make([]byte, size)

	// vamos a leer el contenido de la imagen en el buffer
	fReader := bufio.NewReader(imgFile)
	fReader.Read(buf)

	//  si creas una imagen nueva en lugar de cargarla desde archivo como lo vamos hacer en un momento, codifica el buffer entoences con png.Encode()
	// png.Encode(&buf, image)

	//convierte el buffer de byes a base64 string - use buf.Bytes() para nueva imagen
	//imgBase64Str := base64.StdEncoding.EncodeToString(buf)

	//embebemos la imagen en un html para darle salida
	//	img2html := "<html><body> <img src=\"data:image/jpeg;base64," + imgBase64Str + "\" /></body></html>"
	//imgBase64Display := "data:image/jpeg;base64," + imgBase64Str

	//w.Write([]byte(fmt.Sprintf(imgBase64Display)))
	w.Write(buf)
}

/*
func NamedGoldCardHandler(w http.ResponseWriter, r *http.Request) {

	authCode := os.Getenv("authCode")
	//log.Println(authCode)
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
*/

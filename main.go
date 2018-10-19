package main

import (
	"bytes"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/golang/freetype"
	"golang.org/x/image/font"

	"github.com/gorilla/mux"
)

// restarted

func init() {
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
}

var router = mux.NewRouter()

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	router.Path("/api/namedgoldcard/").Queries("name", "{name}", "code", "{code}", "type", "{type}").HandlerFunc(NamedGoldCardHandler).Name("NamedGoldCardHandler")
	router.Path("/api/namedgoldcard/").Queries("name", "{name}", "code", "{code}").HandlerFunc(NamedGoldCardHandler)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

// NamedGoldCardHandler handles the input request
func NamedGoldCardHandler(w http.ResponseWriter, r *http.Request) {
	//imagen, err := genGoldCard(r.URL.Query().Get("name")) // Original
	nombre := r.FormValue("name")
	//codigo := r.FormValue("code")
	tipoImagen := r.FormValue("type")

	imagen, err := genGoldCard(nombre, tipoImagen)

	if err != nil {
		fmt.Println(err)
		log.Println(err)
		//os.Exit(1)
	}

	w.Write(imagen)
}

func genGoldCard(nombre string, tipo string) (goldCard []byte, err error) {

	var fontSize float64 = 86
	shortName := "Usuario Gold"

	if nombre != "" {
		shortName = cleanName(nombre)
	}

	/// Carga el archido de imagen
	//imgFile, err := os.Open("./resources/goldcard.jpg") // gold card
	imgFile, err := Asset("resources/goldcard.jpg")

	if err != nil {
		fmt.Println(err)
		log.Println(err)
	}

	//defer imgFile.Close()

	// Carga la fuente
	fontFile, err := Asset("resources/avenir_font.ttf")
	if err != nil {
		log.Println("Font not found")
		os.Exit(1)
	}
	//fontFile := "./resources/avenir_font.ttf"
	//fontBytes, err := ioutil.ReadFile(fontFile)
	/*if err != nil {
		log.Println(err)
	}*/

	//f, err := freetype.ParseFont(fontBytes) // this is with the file
	f, err := freetype.ParseFont(fontFile) //this is with bindata
	if err != nil {
		log.Println(err)
	}

	// Comenzamos a generar la imagen
	//img, _, err := image.Decode(imgFile)
	img, _, err := image.Decode(bytes.NewReader(imgFile)) // abre la imagen embebida
	rgba := image.NewRGBA(image.Rect(0, 0, 948, 597))
	draw.Draw(rgba, rgba.Bounds(), img, image.ZP, draw.Src)
	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(f)
	c.SetFontSize(fontSize)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	fg := image.White
	c.SetSrc(fg)
	c.SetHinting(font.HintingFull)

	//****** dibuja el texto *****//

	pt := freetype.Pt(42, 128+int(c.PointToFixed(fontSize)>>6))
	_, err = c.DrawString(shortName, pt)
	if err != nil {
		log.Println(err)
	}
	//creamos un buffer del tamaño de la imagen
	buf := new(bytes.Buffer)

	//  si creas una imagen nueva en lugar de cargarla desde archivo como lo vamos hacer en un momento, codifica el buffer entoences con png.Encode()
	//png.Encode(&buf, rgba)

	switch tipo {
	case "png":
		// no se necesita comprimir por que la compresión es loseless en PNG
		err := png.Encode(buf, rgba)
		if err != nil {
			log.Fatal(err)
		}
	default: // Encode to JPEG
		var opt jpeg.Options
		opt.Quality = 50
		if err := jpeg.Encode(buf, rgba, &opt); err != nil {
			log.Fatal(err)
		}
	}

	//convierte el buffer de byes a base64 string - use buf.Bytes() para nueva imagen
	//imgBase64Str := base64.StdEncoding.EncodeToString(buf)

	//embebemos la imagen en un html para darle salida
	//	img2html := "<html><body> <img src=\"data:image/jpeg;base64," + imgBase64Str + "\" /></body></html>"
	//imgBase64Display := "data:image/jpeg;base64," + imgBase64Str

	//w.Write([]byte(fmt.Sprintf(imgBase64Display)))
	goldImage := buf.Bytes()
	return goldImage, nil
}

func cleanName(nombre string) (name string) {
	splitName := strings.Split(nombre, " ")
	shortName := splitName[0] + " " + splitName[1]
	return shortName
}

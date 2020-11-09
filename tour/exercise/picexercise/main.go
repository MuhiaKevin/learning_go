package main

import (
	"html/template"
	"net/http"
	"picexercise/pic"
	// "golang.org/x/tour/pic"
)

type ImageExercise struct {
	ImageBase64 string
}

// https://stackoverflow.com/questions/39804861/what-is-a-concise-way-to-create-a-2d-slice-in-go
// https://golangbot.com/go-packages/
// https://stackoverflow.com/questions/25459474/go-tour-slices-exercise-logic
// https://base64.guru/converter/decode/image
// https://stackoverflow.com/questions/2659312/how-do-i-convert-a-numpy-array-to-and-display-an-image

func Pic(dx, dy int) [][]uint8 {
	// create the intial slice of type uint8 and  length dy
	pic := make([][]uint8, dy)

	// loop the slice
	for y := range pic {
		// add dx elements tp each element of the intial slice
		for x := 0; x < dx; x++ {
			// each element should have a value of type unint8 of size (x+y)/2, x*y, and x^y (x^y)*(x^y)
			pic[y] = append(pic[y], uint8((x ^ y)))
		}
	}

	return pic
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	p := ImageExercise{ImageBase64: pic.IMAGESTRING}
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, p)
}

func main() {
	// when run in the terminal base64 encoding string is returned
	pic.Show(Pic)
	// TODO: show a template with base64 image from pic
	// pic.IMAGESTRING

	http.HandleFunc("/", newsAggHandler)
	http.ListenAndServe(":8000", nil)
}

package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var maps1 map[string]Vertex

func main() {
	// maps1 = make(map[string]Vertex)
	// maps1["Bell Labs"] = Vertex{
	// 	40.68433, -74.39967,
	// }
	// // prints the value of key Bell Labs
	// fmt.Println(maps1["Bell Labs"])

	// // Map literals => keys are required
	// var maps2 = map[string]Vertex{
	// 	"Bell Labs": Vertex{
	// 		40.68433, -74.39967,
	// 	},
	// 	"Google": Vertex{
	// 		37.42202, -122.08408,
	// 	},
	// }

	// fmt.Println(maps2)

	// mutating or changing maps
	m := make(map[string][]string)
	m["Boys"] = append(m["Boys"], "Kefin", "George", "Chris")
	m["Girls"] = append(m["Girls"], "Mary", "Stephanie", "Alexi")
	delete(m, "Boys") // deletes the key and its values
	fmt.Println(m["Girls"], m["Boys"])

}

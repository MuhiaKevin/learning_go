package main

import(
	"sort"
	"fmt"
)


type human struct{
	fullname string
	height int
	weight int
}


func main(){
	slice := make([]human, 0)
	humanobj := human{"Kevin Muhia",176, 79}
	slice = append(slice, humanobj)
	humanobj = human{"Geoffrey Kahiga Muhia",170, 70}
	slice = append(slice, humanobj)
	humanobj = human{"Hanslit",176, 79}
	slice = append(slice, humanobj)
	humanobj = human{"Allan Mbugua",154, 40}
	slice = append(slice, humanobj)
	humanobj = human{"Allan Mbugua",154, 40}
	slice = append(slice, humanobj)
	// fmt.Println("0:", slice)
	sort.Slice(slice, func(i, j int) bool{
		return slice[i].weight < slice[j].weight
	})
	fmt.Println("<:", slice)
	sort.Slice(slice, func(i, j int) bool {
		return slice[i].weight >slice[j].weight
	})
	fmt.Println(">:", slice)
	
}
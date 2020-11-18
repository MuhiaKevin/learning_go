package main

import (
	"fmt"
	"sort"
)

type person struct {
	name   string
	height float64
	weight int
}

func main() {
	people := make([]person, 0)

	a := person{"Kefin Muhia", 1.76, 84}
	people = append(people, a)

	a = person{"Osu Papa", 1.82, 66}
	people = append(people, a)

	a = person{"Papu Gomez", 1.60, 60}
	people = append(people, a)

	a = person{"james paul", 1.91, 93}
	people = append(people, a)

	a = person{"Reece Francise", 1.85, 84}
	people = append(people)

	fmt.Println("People : ", people)
	fmt.Println("Sorted people by weight : ", sortSlice(people))
}

func sortSlice(people []person) []person {
	sort.Slice(people, func(i, j int) bool { return people[i].weight < people[j].weight })
	return people
}

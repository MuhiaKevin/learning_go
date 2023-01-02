package main

import (
	"fmt"
	// 	"github.com/AlecAivazis/survey/v2"
)

type Activity struct {
	name        string
	description string
}
type Todo struct {
	list []Activity
}

func (todo *Todo) updateTodo(activity Activity) {
	todo.list = append(todo.list, activity)
}

func main() {
	var todo Todo
	// activity := Activity{name: "Code in  Golang", description: "I want to code in golang"}
	todo.updateTodo(Activity{name: "Feed Chickens", description: "I want to feed the chickents"})
	todo.updateTodo(Activity{name: "Code in  Golang", description: "I want to code in golang"})
	todo.updateTodo(Activity{name: "Do some home work", description: "I Want to do some homework"})
	todo.updateTodo(Activity{name: "Code in  Rust", description: "I want to code in Rust"})
	fmt.Println(todo.list)

}

// package main

// import (
// 	"fmt"

// 	"github.com/AlecAivazis/survey/v2"
// )

// // the questions to ask
// var qs = []*survey.Question{
// 	{
// 		Name:      "name",
// 		Prompt:    &survey.Input{Message: "What is your name?"},
// 		Validate:  survey.Required,
// 		Transform: survey.Title,
// 	},
// 	{
// 		Name: "color",
// 		Prompt: &survey.Select{
// 			Message: "Choose a color:",
// 			Options: []string{"red", "blue", "green"},
// 			Default: "red",
// 		},
// 	},
// 	{
// 		Name:   "age",
// 		Prompt: &survey.Input{Message: "How old are you?"},
// 	},
// }

// func main() {
// 	// the answers will be written to this struct
// 	answers := struct {
// 		Name          string // survey will match the question and field names
// 		FavoriteColor string `survey:"color"` // or you can tag fields to match a specific name
// 		Age           int    // if the types don't match, survey will convert it
// 	}{}

// 	// perform the questions
// 	err := survey.Ask(qs, &answers)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	fmt.Printf("%s chose %s.", answers.Name, answers.FavoriteColor)
// }

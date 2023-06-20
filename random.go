package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"

// 	"github.com/AlecAivazis/survey/v2"
// 	"github.com/tjarratt/babble"
// )

// var qs = []*survey.Question{
// 	{
// 		Name: "generate",
// 		Prompt: &survey.Select{
// 			Message: "Choose:",
// 			Options: []string{"words", "numbers"},
// 		},
// 	},
// }

// func main() {
// 	answers := struct {
// 		Choose string `survey:"generate"` // or you can tag fields to match a specific name
// 	}{}
// 	rand.Seed(time.Now().UnixNano())
// 	babbler := babble.NewBabbler()
// 	fmt.Println("welcome to random generator ")
// 	time.Sleep(500 * time.Millisecond)
// 	fmt.Println("we can generate words and even numbers")
// 	time.Sleep(500 * time.Millisecond)
// 	err := survey.Ask(qs, &answers)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
// 	switch answers.Choose {
// 	case "words":
// 		fmt.Println("here is your word")
// 		babbler.Separator = " "
// 		println(babbler.Babble())
// 	case "numbers":
// 		fmt.Println("here are your numbers")
// 		fmt.Println(rand.Int())
// 	default:
// 		fmt.Println("i didn't understand here are both ")
// 		time.Sleep(2 * time.Second)
// 		fmt.Println("here are your numbers")
// 		fmt.Println(rand.Int())
// 		fmt.Println("here is your word")
// 		babbler.Separator = " "
// 		println(babbler.Babble())
// 	}
// }

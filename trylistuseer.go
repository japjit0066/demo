package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"

// 	"github.com/AlecAivazis/survey/v2"
// )

// func Checkboxes(label string, opts []string) []string {
// 	res := []string{}
// 	prompt := &survey.MultiSelect{
// 		Message: label,
// 		Options: opts,
// 	}
// 	survey.AskOne(prompt, &res)

// 	return res
// }

// func main() {
// 	answers := Checkboxes(
// 		"Which proggramming languages do you want to learn ?",
// 		[]string{
// 			"Html",
// 			"Css",
// 			"Typescript",
// 			"Astrojs",
// 			"Nextjs",
// 			"Solidjs",
// 			"JavaScript",
// 			"Go",
// 		},
// 	)
// 	s := strings.Join(answers, ", ")
// 	fmt.Println("Oh, I see! You want to learn", s)
// 	for {

// 		switch s {
// 		case "Html":
// 			fmt.Println("Html nice")
// 			fmt.Println("hers all the learning sites")
// 			fmt.Println("www.w3schools.com")
// 			fmt.Println("www.codecademy.com/learn/learn-html")
// 			fmt.Println("www.learn-html.org")
// 			fmt.Println("web.dev/learn/html/lists/")
// 			fmt.Println("www.simplilearn.com/free-html-course-for-beginners-skillup")
// 			fmt.Println("www.mygreatlearning.com/academy/learn-for-free/courses/html-tutorial")
// 		case "Css":
// 			fmt.Println("Css nice")
// 		case "Typescript":
// 			fmt.Println("Typescript nice")
// 		case "Astrojs":
// 			fmt.Println("Astrojs nice")
// 		case "Nextjs":
// 			fmt.Println("Nextjs nice")
// 		case "Solidjs":
// 			fmt.Println("Solidjs nice")
// 		case "JavaScript":
// 			fmt.Println("JavaScript nice")
// 		case "Go":
// 			fmt.Println("Go nice")
// 		default:
// 			fmt.Println("choose again ")
// 			answers := Checkboxes(
// 				"choose only one language",
// 				[]string{
// 					"Html",
// 					"Css",
// 					"Typescript",
// 					"Astrojs",
// 					"Nextjs",
// 					"Solidjs",
// 					"JavaScript",
// 					"Go",
// 				},
// 			)
// 			s := strings.Join(answers, ", ")
// 			fmt.Println("Oh, I see!", s)
// 			switch s {
// 			case "Html":
// 				fmt.Println("Html nice")
// 				fmt.Println("hers all the learning sites")
// 				fmt.Println("www.w3schools.com")
// 				fmt.Println("www.codecademy.com/learn/learn-html")
// 				fmt.Println("www.learn-html.org")
// 				fmt.Println("web.dev/learn/html/lists/")
// 				fmt.Println("www.simplilearn.com/free-html-course-for-beginners-skillup")
// 				fmt.Println("www.mygreatlearning.com/academy/learn-for-free/courses/html-tutorial")
// 				os.Exit(0)
// 			case "Css":
// 				fmt.Println("Css nice")
// 			case "Typescript":
// 				fmt.Println("Typescript nice")
// 			case "Astrojs":
// 				fmt.Println("Astrojs nice")
// 			case "Nextjs":
// 				fmt.Println("Nextjs nice")
// 			case "Solidjs":
// 				fmt.Println("Solidjs nice")
// 			case "JavaScript":
// 				fmt.Println("JavaScript nice")
// 			case "Go":
// 				fmt.Println("Go nice")
// 			}
// 		}
// 	}

// }

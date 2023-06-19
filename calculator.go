package main

import (
	"fmt"
	"strings"
	"time"
)

var num1 int64
var num2 int64
var charcter string

func main() {
	for {
		welcome()
		Answer()
	}
}
func welcome() {
	fmt.Println("welcome to calcnum ™")
	time.Sleep(400 * time.Millisecond)
	fmt.Println("lets start ")
	time.Sleep(400 * time.Millisecond)
	fmt.Println("pls enter a number  : ")
	fmt.Scan(&num1)
	time.Sleep(400 * time.Millisecond)
	fmt.Println("do you want to multiple , add, divide or subtract (choose `*` = multiple, `+` = add , `-` = subtract, `:` = divide)")
	fmt.Scan(&charcter)
	time.Sleep(400 * time.Millisecond)
	fmt.Println("pls enter the other number  : ")
	fmt.Scan(&num2)
	time.Sleep(400 * time.Millisecond)
	fmt.Println("ok wait two sec... ")
	for i := 0; i < 30; i++ {
		time.Sleep(500 * time.Microsecond) // mimics work
		printProgressBar(i+1, 30, "Progress", "", 25, "#")
	}
}
func Answer() {
	if charcter == "*" {
		fmt.Println(num1 * num2)
	} else if charcter == "+" {
		fmt.Println(num1 + num2)
	} else if charcter == "-" {
		fmt.Println(num1 - num2)
	} else if charcter == ":" {
		fmt.Println(num1 / num2)
	}
}

func printProgressBar(iteration, total int, prefix, suffix string, length int, fill string) {

	percent := uint(iteration) / uint(total)

	filledLength := int(length * iteration / total)
	iteration = 0
	total = 100
	end := ">>"

	if iteration == total {
		end = "#"
	}
	if iteration == total {
		suffix = "done"
	}
	bar := strings.Repeat(fill, filledLength) + end + strings.Repeat("-", (length-filledLength))
	fmt.Printf("\r%s [%s]  %d%% %s", prefix, bar, percent, suffix)
	if iteration == total {
		fmt.Println()
	}
}

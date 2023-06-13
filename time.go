package main

// import (
// 	"fmt"
// 	"os"
// 	"time"
// )

// var name = "time teller"
// var Input string
// var yesNo string
// var pl = fmt.Println
// var what string

// func main() {
// 	pl("welcome to ", name)
// 	pl(name, " tells the time and the date  !!!! ")
// 	pl("pls enter your name :")
// 	fmt.Scan(&Input)
// 	pl("hi ", Input)

// 	pl("do you want to try ", name, "(yes or no)")
// 	fmt.Scan(&yesNo)
// 	if yesNo == "no" {
// 		fmt.Println("ok then want are you doing he ?")
// 		fmt.Scan(&what)
// 		fmt.Println("ok then bye")
// 		os.Exit(0)
// 	} else {
// 		now := time.Now()
// 		pl("ok then lets tell you the date !!!")
// 		pl(now.Day(), "/", now.Month(), "/", now.Local().Year())
// 		pl("and the time is :")
// 		pl(now.Hour(), ":", now.Local().Minute(), ":", now.Local().Second())

// 	}
// }

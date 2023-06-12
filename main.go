package main

import (
	"fmt"
	"math/rand"
	"time"
)

var yesNo string

func main() {
	var email string
	var noNo string
	var companyName = "levissima"
	var water = ""
	var name string
	var number int
	var Input string
	fmt.Println("welcome to", companyName)
	time.Sleep(500 * time.Millisecond)
	fmt.Println("enter your first name :")
	fmt.Scan(&name)
	time.Sleep(200 * time.Millisecond)
	fmt.Println("ok enter your email please : ")
	fmt.Scan(&email)
	time.Sleep(200 * time.Millisecond)
	for {
		fmt.Println("we have three type of water")
		time.Sleep(200 * time.Millisecond)
		fmt.Println("flavoured  , sparkling  , natural ")
		time.Sleep(200 * time.Millisecond)
		fmt.Println("choose your type of water ")
		fmt.Scan(&water)
		switch water {
		case "flavoured":

			rand.Seed(time.Now().UnixNano())

			fmt.Println()
			fmt.Println("this is flavoured water")
			time.Sleep(300 * time.Millisecond)
			fmt.Println("do you want to buy flavoured water (yes or no) ")
			fmt.Scan(&yesNo)
			if yesNo == "no" {
				fmt.Println("ok choose again")
				continue

			} else {

				time.Sleep(500 * time.Millisecond)
				fmt.Println("how many bottles : ")
				fmt.Scan(&number)
				time.Sleep(500 * time.Millisecond)
				fmt.Println("we will send an email to ", email)
				time.Sleep(200 * time.Millisecond)
				fmt.Printf("thanks for buying %v flavoured water bottles \n", number)
				fmt.Println("heres your order code:", rand.Int())
				time.Sleep(200 * time.Millisecond)
				fmt.Println("see you next time ", name)

			}
		case "sparkling":
			var yesNo string
			fmt.Println("this is sparkling water")
			fmt.Println("do you want to buy sparkling water (yes or no) ")
			fmt.Scan(&yesNo)
			if yesNo == "no" {
				fmt.Println("ok choose again")
				continue
			} else {

				time.Sleep(500 * time.Millisecond)
				fmt.Println("how many bottles : ")
				fmt.Scan(&number)
				time.Sleep(500 * time.Millisecond)
				fmt.Println("we will send an email to ", email)
				time.Sleep(200 * time.Millisecond)
				fmt.Printf("thanks for buying %v sparkling water bottles \n", number)
				fmt.Println("heres your order code:", rand.Int())
				time.Sleep(1200 * time.Millisecond)
				fmt.Println("see you next time ", name)

			}
		case "natural":
			var yesNo string
			fmt.Println("this is natural water")
			fmt.Println("do you want to buy natural water (yes or no) ")
			fmt.Scan(&yesNo)
			if yesNo == "no" {
				fmt.Println("ok choose again")
				continue
			} else {

				time.Sleep(500 * time.Millisecond)
				fmt.Println("how many bottles : ")
				fmt.Scan(&number)
				time.Sleep(500 * time.Millisecond)
				fmt.Println("we will send an email to ", email)
				time.Sleep(200 * time.Millisecond)
				fmt.Printf("thanks for buying %v natural water bottles \n", number)
				fmt.Println("heres your order code:", rand.Int())
				time.Sleep(1200 * time.Millisecond)

			}
		}

		fmt.Println("do you want to buy more water ? (yes or no)")
		fmt.Scan(&noNo)
		if noNo == "no" {
			fmt.Println("ok then bye see you next time", name)
			break

		} else {
			fmt.Println("do you want to use the same email that is ", email, "(yes or no)")
			fmt.Scan(&Input)
			if Input == "no" {
				fmt.Println("please enter your email :")
				fmt.Scan(&email)
				fmt.Println("we will use this email for your order", email)
				continue
			} else {
				for {

					fmt.Println("we have three type of water")
					time.Sleep(200 * time.Millisecond)
					fmt.Println("flavoured  , sparkling  , natural ")
					time.Sleep(200 * time.Millisecond)
					fmt.Println("choose your type of water ")
					fmt.Scan(&water)
					switch water {
					case "flavoured":

						rand.Seed(time.Now().UnixNano())

						fmt.Println()
						fmt.Println("this is flavoured water")
						time.Sleep(300 * time.Millisecond)
						fmt.Println("do you want to buy flavoured water (yes or no) ")
						fmt.Scan(&yesNo)
						if yesNo == "no" {
							fmt.Println("ok choose again ")
							continue
						} else {
							time.Sleep(500 * time.Millisecond)
							fmt.Println("how many bottles : ")
							fmt.Scan(&number)
							time.Sleep(500 * time.Millisecond)
							fmt.Println("we will send an email to ", email)
							time.Sleep(200 * time.Millisecond)
							fmt.Printf("thanks for buying %v flavoured water bottles \n", number)
							fmt.Println("heres your order code:", rand.Int())
							time.Sleep(200 * time.Millisecond)
							fmt.Println("see you next time ", name)

						}
					case "sparkling":
						var yesNo string
						fmt.Println("this is sparkling water")
						fmt.Println("do you want to buy sparkling water (yes or no) ")
						fmt.Scan(&yesNo)
						if yesNo == "no" {
							fmt.Println("ok choose again ")
							continue
						} else {
							time.Sleep(500 * time.Millisecond)
							fmt.Println("how many bottles : ")
							fmt.Scan(&number)
							time.Sleep(500 * time.Millisecond)
							fmt.Println("we will send an email to ", email)
							time.Sleep(200 * time.Millisecond)
							fmt.Printf("thanks for buying %v sparkling water bottles \n", number)
							fmt.Println("heres your order code:", rand.Int())
							time.Sleep(1200 * time.Millisecond)
							fmt.Println("see you next time ", name)

						}
					case "natural":
						var yesNo string
						fmt.Println("this is natural water")
						fmt.Println("do you want to buy natural water (yes or no) ")
						fmt.Scan(&yesNo)
						if yesNo == "no" {
							fmt.Println("ok choose again ")
							continue
						} else {
							var email string
							time.Sleep(500 * time.Millisecond)
							fmt.Println("how many bottles : ")
							fmt.Scan(&number)
							time.Sleep(500 * time.Millisecond)
							fmt.Println("we will send an email to ", email)
							time.Sleep(200 * time.Millisecond)
							fmt.Printf("thanks for buying %v natural water bottles \n", number)
							fmt.Println("heres your order code:", rand.Int())
							time.Sleep(1200 * time.Millisecond)
						}
					}
					fmt.Println("do you want to buy more water ? (yes or no)")
					fmt.Scan(&noNo)
					if noNo == "no" {
						break
					} else {

						fmt.Println("do you want to use the same email that is ", email, "(yes or no)")
						fmt.Scan(&Input)
						if Input == "no" {
							fmt.Println("please enter your email :")
							fmt.Scan(&email)
							fmt.Println("we will use this email for your order", email)
							continue
						}

					}
				}
				fmt.Println("are you sure (yes or no)")
				fmt.Scan(&Input)
				if Input == "yes" {
					fmt.Println("ok bye see you next time")
					break
				}
				if Input == "no" {
					fmt.Println("okay then lets continue")
					continue
				}
			}
		}
	}
}

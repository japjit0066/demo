package main

// import (
// 	"fmt"
// 	"os"
// 	"time"
// )

// var name = "jsr eletronics"
// var yesNo string
// var item = []string{"phones,", "", "laptops,", "", "pcs,", "", "pc-components,", "", "camera,", "", "headphones,", "", "mics,", "", "tech-furnitures,", "", "cables,", "", "machines,", "", "eletronic-home-gadgets,", "", "screwdriver-ltt,", "", "screwdriver-jsr,", "", "fitness-eletronics,", "", "mouses-and-keyboard", "accessories"}
// var choose string
// var phones = []string{"Apple,", "", "Huawei,", "", "", "Nothing,", "", "onePlus,", "", "Oppo,", "", "Samsung,", "", "vivo and Xiaomi"}
// var phone string
// var model string
// var input string
// var account string
// var pl = fmt.Println
// var sc = fmt.Scan

// func main() {

// 	pl("welcome to", name)
// 	time.Sleep(300 * time.Millisecond)
// 	pl("do you want to buy a computer, ")
// 	time.Sleep(300 * time.Millisecond)
// 	pl("phone or any piece of tech you came to the right place")
// 	time.Sleep(300 * time.Millisecond)
// 	pl("we have every thing even components !!!!!!!!!")
// 	time.Sleep(300 * time.Millisecond)
// 	pl("do you want to buy something ?  (yes or no)")
// 	sc(&yesNo)
// 	if yesNo == "no" {
// 		pl("ok, bye")
// 		os.Exit(0)
// 	} else {
// 		pl("what can i buy ? ")
// 		time.Sleep(600 * time.Millisecond)
// 		pl("exellent question we have : ")
// 		time.Sleep(600 * time.Millisecond)
// 		pl(item)
// 		pl("what do you what to buy ? :")
// 		sc(&choose)
// 		switch choose {
// 		case "phones":
// 			pl("this are phones")
// 			time.Sleep(600 * time.Millisecond)
// 			pl("you may be asking what brands have you got ?")
// 			time.Sleep(600 * time.Millisecond)
// 			pl("well we have :")
// 			pl(phones)
// 			pl("don't know which phone company to choose ? ")
// 			time.Sleep(600 * time.Millisecond)
// 			pl("we will help you with a list")
// 			time.Sleep(3 * time.Second)
// 			pl("")
// 			pl("apple and samsung = easy repair at any repair center , apple expensive and it is ios , samsung has some expensive phones and some cheap phones ,smasung is android , both of the devices have really good quality, and really trusted companies ")
// 			pl("")
// 			pl("Xiaomi , oppo = a little hard to repair at any repair center, they have some expensive phones and some cheap ones , good quality, both android but xiaomi has a custom miui design")
// 			pl("")
// 			pl("Oneplus, vivo = really hard to repair at any repair center, there all cheap phones , good quality")
// 			pl("")
// 			pl("huawei = no more google support but in old phones there is still support, a almost dead company , good quality , all phones are cheap ")
// 			pl("")
// 			pl("nothing = new company , good quality , quality equal to apple but the phone is android, the phone is a custom android , cheap ")
// 			time.Sleep(5 * time.Second)
// 			pl("")
// 			pl("now if you are ready , choose !!!!")
// 			sc(&phone)

// 			switch phone {
// 			case "apple":
// 				pl("this is", phone)
// 				time.Sleep(600 * time.Millisecond)
// 				pl("nice choice apple phones are really good ")
// 				time.Sleep(600 * time.Millisecond)
// 				pl("we have all the newest models : iphone-se ,  iphone-12-mini , iphone-12  , iphone-12-pro , iphone-12-promax , iphone-13-mini , iphone-13 , iphone-13-pro , iphone-13-promax  , iphone-14 , iphone-14-plus , iphone-14-pro , iphone-14-promax ")
// 				sc(&model)
// 				pl("are you sure you want this", model, ", (yes or no)")
// 				sc(&input)
// 				if input == "no" {
// 					pl("ok retry ")
// 					pl("we have all the newest models : iphone-se ,  iphone-12-mini , iphone-12  , iphone-12-pro , iphone-12-promax , iphone-13-mini , iphone-13 , iphone-13-pro , iphone-13-promax  , iphone-14 , iphone-14-plus , iphone-14-pro , iphone-14-promax ")
// 					sc(&model)
// 				}
// 				switch model {
// 				case "iphone-se":
// 					pl("this is", model)
// 					pl("#####################")
// 					pl("it costs $559")
// 					pl("#####################")
// 					pl("are you sure you want to buy an ", model, ", for a price of $559 , (yes or no)")
// 					sc(&input)
// 					if input == "no" {
// 						pl("ok bye")
// 						os.Exit(0)
// 					}
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("please enter your bank account : ")
// 					sc(&account)
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("thank you for buying a", model)
// 					pl("bye see you next time ")
// 					pl("just kidding if you want accessories go back and go to accessories not in phones ")
// 					pl("now i go away")

// 					os.Exit(0)
// 				case "iphone-12-mini":
// 					pl("this is", model)
// 					pl("#####################")
// 					pl("it costs $440")
// 					pl("#####################")
// 					pl("are you sure you want to buy an ", model, ", for a price of $440 , (yes or no)")
// 					sc(&input)
// 					if input == "no" {
// 						pl("ok bye")
// 						os.Exit(0)
// 					}
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("please enter your bank account : ")
// 					sc(&account)
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("thank you for buying a", model)
// 					pl("bye see you next time ")
// 					pl("just kidding if you want accessories go back and go to accessories not in phones ")
// 					pl("now i go away")

// 					os.Exit(0)

// 				case "iphone-12 ":
// 					pl("this is", model)
// 					pl("#####################")
// 					pl("it costs $450")
// 					pl("#####################")
// 					pl("are you sure you want to buy an ", model, ", for a price of $450 , (yes or no)")
// 					sc(&input)
// 					if input == "no" {
// 						pl("ok bye")
// 						os.Exit(0)
// 					}
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("please enter your bank account : ")
// 					sc(&account)
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("thank you for buying a", model)
// 					pl("bye see you next time ")
// 					pl("just kidding if you want accessories go back and go to accessories not in phones ")
// 					pl("now i go away")

// 					os.Exit(0)

// 				case "iphone-12-pro":
// 					pl("this is", model)
// 					pl("#####################")
// 					pl("it costs $868,90")
// 					pl("#####################")
// 					pl("are you sure you want to buy an ", model, ", for a price of $868,90 , (yes or no)")
// 					sc(&input)
// 					if input == "no" {
// 						pl("ok bye")
// 						os.Exit(0)
// 					}
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("please enter your bank account : ")
// 					sc(&account)
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("thank you for buying a", model)
// 					pl("bye see you next time ")
// 					pl("just kidding if you want accessories go back and go to accessories not in phones ")
// 					pl("now i go away")

// 					os.Exit(0)

// 				case "iphone-12-promax":
// 					pl("this is", model)
// 					pl("#####################")
// 					pl("it costs $1.045")
// 					pl("#####################")
// 					pl("are you sure you want to buy an ", model, ", for a price of $1.045 , (yes or no)")
// 					sc(&input)
// 					if input == "no" {
// 						pl("ok bye")
// 						os.Exit(0)
// 					}
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("please enter your bank account : ")
// 					sc(&account)
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("thank you for buying a", model)
// 					pl("bye see you next time ")
// 					pl("just kidding if you want accessories go back and go to accessories not in phones ")
// 					pl("now i go away")

// 					os.Exit(0)

// 				case "iphone-13-mini":
// 					pl("this is", model)
// 					pl("#####################")
// 					pl("it costs $629")
// 					pl("#####################")
// 					pl("are you sure you want to buy an ", model, ", for a price of $629 , (yes or no)")
// 					sc(&input)
// 					if input == "no" {
// 						pl("ok bye")
// 						os.Exit(0)
// 					}
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("please enter your bank account : ")
// 					sc(&account)
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("thank you for buying a", model)
// 					pl("bye see you next time ")

// 					pl("just kidding if you want accessories go back and go to accessories not in phones ")
// 					pl("now i go away")
// 					os.Exit(0)

// 				case "iphone-13":
// 					pl("this is", model)
// 					pl("#####################")
// 					pl("it costs $706")
// 					pl("#####################")
// 					pl("are you sure you want to buy an ", model, ", for a price of $706 , (yes or no)")
// 					sc(&input)
// 					if input == "no" {
// 						pl("ok bye")
// 						os.Exit(0)
// 					}
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("please enter your bank account : ")
// 					sc(&account)
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("thank you for buying a", model)
// 					pl("bye see you next time ")
// 					pl("just kidding if you want accessories go back and go to accessories not in phones ")
// 					pl("now i go away")
// 					os.Exit(0)

// 				case "iphone-13-pro":
// 					pl("this is", model)
// 					pl("#####################")
// 					pl("it costs $ 849")
// 					pl("#####################")
// 					pl("are you sure you want to buy an ", model, ", for a price of $ 849 , (yes or no)")
// 					sc(&input)
// 					if input == "no" {
// 						pl("ok bye")
// 						os.Exit(0)
// 					}
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("please enter your bank account : ")
// 					sc(&account)
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("thank you for buying a", model)
// 					pl("bye see you next time ")
// 					pl("just kidding if you want accessories go back and go to accessories not in phones ")
// 					pl("now i go away")
// 					os.Exit(0)

// 				case "iphone-13-promax":
// 					pl("this is", model)
// 					pl("#####################")
// 					pl("it costs $1.169")
// 					pl("#####################")
// 					pl("are you sure you want to buy an ", model, ", for a price of $1.169 , (yes or no)")
// 					sc(&input)
// 					if input == "no" {
// 						pl("ok bye")
// 						os.Exit(0)
// 					}
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("please enter your bank account : ")
// 					sc(&account)
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("thank you for buying a", model)
// 					pl("bye see you next time ")
// 					pl("just kidding if you want accessories go back and go to accessories not in phones ")
// 					pl("now i go away")
// 					os.Exit(0)

// 				case "iphone-14":
// 					pl("this is", model)
// 					pl("#####################")
// 					pl("it costs $725,99")
// 					pl("#####################")
// 					pl("are you sure you want to buy an ", model, ", for a price of $725,99 , (yes or no)")
// 					sc(&input)
// 					if input == "no" {
// 						pl("ok bye")
// 						os.Exit(0)
// 					}
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("please enter your bank account : ")
// 					sc(&account)
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("thank you for buying a", model)
// 					pl("bye see you next time ")
// 					pl("just kidding if you want accessories go back and go to accessories not in phones ")
// 					pl("now i go away")
// 					os.Exit(0)

// 				case "iphone-14-plus":
// 					pl("this is", model)
// 					pl("#####################")
// 					pl("it costs $787")
// 					pl("#####################")
// 					pl("are you sure you want to buy an ", model, ", for a price of $787 , (yes or no)")
// 					sc(&input)
// 					if input == "no" {
// 						pl("ok bye")
// 						os.Exit(0)
// 					}
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("please enter your bank account : ")
// 					sc(&account)
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("thank you for buying a", model)
// 					pl("bye see you next time ")
// 					pl("just kidding if you want accessories go back and go to accessories not in phones ")
// 					pl("now i go away")
// 					os.Exit(0)

// 				case "iphone-14-pro":
// 					pl("this is", model)
// 					pl("#####################")
// 					pl("it costs $976,00")
// 					pl("#####################")
// 					pl("are you sure you want to buy an ", model, ", for a price of $976,00 , (yes or no)")
// 					sc(&input)
// 					if input == "no" {
// 						pl("ok bye")
// 						os.Exit(0)
// 					}
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("please enter your bank account : ")
// 					sc(&account)
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("thank you for buying a", model)
// 					pl("bye see you next time ")
// 					pl("just kidding if you want accessories go back and go to accessories not in phones ")
// 					pl("now i go away")
// 					os.Exit(0)

// 				case "iphone-14-promax":
// 					pl("this is", model)
// 					pl("#####################")
// 					pl("it costs $1.087")
// 					pl("#####################")
// 					pl("are you sure you want to buy an ", model, ", for a price of $1.087 , (yes or no)")
// 					sc(&input)
// 					if input == "no" {
// 						pl("ok bye")
// 						os.Exit(0)
// 					}
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("please enter your bank account : ")
// 					sc(&account)
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("thank you for buying a", model)
// 					pl("bye see you next time ")
// 					pl("just kidding if you want accessories go back and go to accessories not in phones ")
// 					pl("now i go away")

// 					os.Exit(0)
// 				default:
// 					fmt.Println("something went wrong")
// 					os.Exit(0)
// 				}
// 			case "samsung":
// 				pl("this is", phone)
// 				time.Sleep(600 * time.Millisecond)
// 				pl("nice choice apple phones are really good ")
// 				time.Sleep(600 * time.Millisecond)
// 				pl("we have all the newest models : galaxy-a14 ,  galaxy-a04s , galaxy-s10 , galaxy-s20 , galaxy-s20e , galaxy-s20-fe , galaxy-s20-ultra , galaxy-s21 , galaxy-s21-ultra  , galaxy-s22 , galaxy-s22-ultra  , galaxy-s23 , galaxy-s23-ultra ")
// 				sc(&model)
// 				pl("are you sure you want this", model, ", (yes or no)")
// 				sc(&input)
// 				if input == "no" {
// 					pl("ok retry ")
// 					pl("we have all the newest models : galaxy-a14 ,  galaxy-a04s , galaxy-a40s , galaxy-s10-lite   , galaxy-s20 ,  , galaxy-s20-fe , galaxy-s20-ultra , galaxy-s21 , galaxy-s21-ultra  , galaxy-s22 , galaxy-s22-ultra  , galaxy-s23 , galaxy-s23-ultra  ")
// 					sc(&model)
// 				}
// 				switch model {
// 				case "galaxy-a14":
// 					pl("this is", model)
// 					pl("#####################")
// 					pl("it costs $189.58")
// 					pl("#####################")
// 					pl("are you sure you want to buy an ", model, ", for a price of $189.58, (yes or no)")
// 					sc(&input)
// 					if input == "no" {
// 						pl("ok bye")
// 						os.Exit(0)
// 					}
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("please enter your bank account : ")
// 					sc(&account)
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("thank you for buying a", model)
// 					pl("bye see you next time ")
// 					pl("just kidding if you want accessories go back and go to accessories not in phones ")
// 					pl("now i go away")

// 					os.Exit(0)
// 				case "galaxy-a04s":
// 					pl("this is", model)
// 					pl("#####################")
// 					pl("it costs $165")
// 					pl("#####################")
// 					pl("are you sure you want to buy an ", model, ", for a price of $165 , (yes or no)")
// 					sc(&input)
// 					if input == "no" {
// 						pl("ok bye")
// 						os.Exit(0)
// 					}
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("please enter your bank account : ")
// 					sc(&account)
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("thank you for buying a", model)
// 					pl("bye see you next time ")
// 					pl("just kidding if you want accessories go back and go to accessories not in phones ")
// 					pl("now i go away")

// 					os.Exit(0)
// 				case "galaxy-a40s":
// 					pl("this is", model)
// 					pl("#####################")
// 					pl("it costs $324.00")
// 					pl("#####################")
// 					pl("are you sure you want to buy an ", model, ", for a price of $324.00 , (yes or no)")
// 					sc(&input)
// 					if input == "no" {
// 						pl("ok bye")
// 						os.Exit(0)
// 					}
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("please enter your bank account : ")
// 					sc(&account)
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("thank you for buying a", model)
// 					pl("bye see you next time ")
// 					pl("just kidding if you want accessories go back and go to accessories not in phones ")
// 					pl("now i go away")
// 					pl("now i go away")
// 					os.Exit(0)

// 				case "galaxy-s10":
// 					pl("this is", model)
// 					pl("#####################")
// 					pl("it costs $649")
// 					pl("#####################")
// 					pl("are you sure you want to buy an ", model, ", for a price of $649 , (yes or no)")
// 					sc(&input)
// 					if input == "no" {
// 						pl("ok bye")
// 						os.Exit(0)
// 					}
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("please enter your bank account : ")
// 					sc(&account)
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("thank you for buying a", model)
// 					pl("bye see you next time ")
// 					pl("just kidding if you want accessories go back and go to accessories not in phones ")
// 					pl("now i go away")

// 					os.Exit(0)

// 				case "galaxy-s20":
// 					pl("this is", model)
// 					pl("#####################")
// 					pl("it costs $649")
// 					pl("#####################")
// 					pl("are you sure you want to buy an ", model, ", for a price of $649 , (yes or no)")
// 					sc(&input)
// 					if input == "no" {
// 						pl("ok bye")
// 						os.Exit(0)
// 					}
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("please enter your bank account : ")
// 					sc(&account)
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("thank you for buying a", model)
// 					pl("bye see you next time ")
// 					pl("just kidding if you want accessories go back and go to accessories not in phones ")

// 					pl("now i go away")
// 					os.Exit(0)

// 				case "galaxy-s20-fe":
// 					pl("this is", model)
// 					pl("#####################")
// 					pl("it costs $ ")
// 					pl("#####################")
// 					pl("are you sure you want to buy an ", model, ", for a price of $629 , (yes or no)")
// 					sc(&input)
// 					if input == "no" {
// 						pl("ok bye")
// 						os.Exit(0)
// 					}
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("please enter your bank account : ")
// 					sc(&account)
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("thank you for buying a", model)
// 					pl("bye see you next time ")
// 					pl("just kidding if you want accessories go back and go to accessories not in phones ")

// 					pl("now i go away")
// 					os.Exit(0)

// 				case "galaxy-s20-ultra":
// 					pl("this is", model)
// 					pl("#####################")
// 					pl("it costs $720")
// 					pl("#####################")
// 					pl("are you sure you want to buy an ", model, ", for a price of $720 , (yes or no)")
// 					sc(&input)
// 					if input == "no" {
// 						pl("ok bye")
// 						os.Exit(0)
// 					}
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("please enter your bank account : ")
// 					sc(&account)
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("thank you for buying a", model)
// 					pl("bye see you next time ")
// 					pl("just kidding if you want accessories go back and go to accessories not in phones ")

// 					pl("now i go away")
// 					os.Exit(0)

// 				case "galaxy-s21":
// 					pl("this is", model)
// 					pl("#####################")
// 					pl("it costs $ 849")
// 					pl("#####################")
// 					pl("are you sure you want to buy an ", model, ", for a price of $ 849 , (yes or no)")
// 					sc(&input)
// 					if input == "no" {
// 						pl("ok bye")
// 						os.Exit(0)
// 					}
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("please enter your bank account : ")
// 					sc(&account)
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("thank you for buying a", model)
// 					pl("bye see you next time ")
// 					pl("just kidding if you want accessories go back and go to accessories not in phones ")

// 					pl("now i go away")
// 					os.Exit(0)

// 				case "galaxy-s21-ultra":
// 					pl("this is", model)
// 					pl("#####################")
// 					pl("it costs $1.169")
// 					pl("#####################")
// 					pl("are you sure you want to buy an ", model, ", for a price of $1.169 , (yes or no)")
// 					sc(&input)
// 					if input == "no" {
// 						pl("ok bye")
// 						os.Exit(0)
// 					}
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("please enter your bank account : ")
// 					sc(&account)
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("thank you for buying a", model)
// 					pl("bye see you next time ")
// 					pl("just kidding if you want accessories go back and go to accessories not in phones ")

// 					pl("now i go away")
// 					os.Exit(0)

// 				case "galaxy-s22":
// 					pl("this is", model)
// 					pl("#####################")
// 					pl("it costs $725,99")
// 					pl("#####################")
// 					pl("are you sure you want to buy an ", model, ", for a price of $725,99 , (yes or no)")
// 					sc(&input)
// 					if input == "no" {
// 						pl("ok bye")
// 						os.Exit(0)
// 					}
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("please enter your bank account : ")
// 					sc(&account)
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("thank you for buying a", model)
// 					pl("bye see you next time ")
// 					pl("just kidding if you want accessories go back and go to accessories not in phones ")

// 					pl("now i go away")
// 					os.Exit(0)

// 				case "galaxy-s22-ultra":
// 					pl("this is", model)
// 					pl("#####################")
// 					pl("it costs $787")
// 					pl("#####################")
// 					pl("are you sure you want to buy an ", model, ", for a price of $787 , (yes or no)")
// 					sc(&input)
// 					if input == "no" {
// 						pl("ok bye")
// 						os.Exit(0)
// 					}
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("please enter your bank account : ")
// 					sc(&account)
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("thank you for buying a", model)
// 					pl("bye see you next time ")
// 					pl("just kidding if you want accessories go back and go to accessories not in phones ")

// 					pl("now i go away")
// 					os.Exit(0)

// 				case "galaxy-s23":
// 					pl("this is", model)
// 					pl("#####################")
// 					pl("it costs $976,00")
// 					pl("#####################")
// 					pl("are you sure you want to buy an ", model, ", for a price of $976,00 , (yes or no)")
// 					sc(&input)
// 					if input == "no" {
// 						pl("ok bye")
// 						os.Exit(0)
// 					}
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("please enter your bank account : ")
// 					sc(&account)
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("thank you for buying a", model)
// 					pl("bye see you next time ")
// 					pl("just kidding if you want accessories go back and go to accessories not in phones ")
// 					pl("now i go away")
// 					os.Exit(0)

// 				case "galaxy-s23-ultra":
// 					pl("this is", model)
// 					pl("#####################")
// 					pl("it costs $1.087")
// 					pl("#####################")
// 					pl("are you sure you want to buy an ", model, ", for a price of $1.087 , (yes or no)")
// 					sc(&input)
// 					if input == "no" {
// 						pl("ok bye")
// 						os.Exit(0)
// 					}
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("please enter your bank account : ")
// 					sc(&account)
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("thank you for buying a", model)
// 					pl("bye see you next time ")
// 					pl("just kidding if you want accessories go back and go to accessories not in phones ")
// 					pl("now i go away")

// 					os.Exit(0)
// 				default:
// 					fmt.Println("something went wrong")
// 					os.Exit(0)
// 				}

// 			case "xiaomi":
// 				pl("this is", phone)
// 				time.Sleep(600 * time.Millisecond)
// 				pl("nice choice apple phones are really good ")
// 				time.Sleep(600 * time.Millisecond)
// 				pl("we have all the newest models : galaxy-a14 ,  galaxy-a04s , galaxy-s10 , galaxy-s20 , galaxy-s20e , galaxy-s20-fe , galaxy-s20-ultra , galaxy-s21 , galaxy-s21-ultra  , galaxy-s22 , galaxy-s22-ultra  , galaxy-s23 , galaxy-s23-ultra ")
// 				sc(&model)
// 				pl("are you sure you want this", model, ", (yes or no)")
// 				sc(&input)
// 				if input == "no" {
// 					pl("ok retry ")
// 					pl("we have all the newest models : galaxy-a14 ,  galaxy-a04s , galaxy-a40s , galaxy-s10-lite   , galaxy-s20 ,  , galaxy-s20-fe , galaxy-s20-ultra , galaxy-s21 , galaxy-s21-ultra  , galaxy-s22 , galaxy-s22-ultra  , galaxy-s23 , galaxy-s23-ultra  ")
// 					sc(&model)
// 				}
// 				switch model {
// 				case "galaxy-a14":
// 					pl("this is", model)
// 					pl("#####################")
// 					pl("it costs $189.58")
// 					pl("#####################")
// 					pl("are you sure you want to buy an ", model, ", for a price of $189.58, (yes or no)")
// 					sc(&input)
// 					if input == "no" {
// 						pl("ok bye")
// 						os.Exit(0)
// 					}
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("please enter your bank account : ")
// 					sc(&account)
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("thank you for buying a", model)
// 					pl("bye see you next time ")
// 					pl("just kidding if you want accessories go back and go to accessories not in phones ")
// 					pl("now i go away")

// 					os.Exit(0)
// 				case "galaxy-a04s":
// 					pl("this is", model)
// 					pl("#####################")
// 					pl("it costs $165")
// 					pl("#####################")
// 					pl("are you sure you want to buy an ", model, ", for a price of $165 , (yes or no)")
// 					sc(&input)
// 					if input == "no" {
// 						pl("ok bye")
// 						os.Exit(0)
// 					}
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("please enter your bank account : ")
// 					sc(&account)
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("thank you for buying a", model)
// 					pl("bye see you next time ")
// 					pl("just kidding if you want accessories go back and go to accessories not in phones ")
// 					pl("now i go away")

// 					os.Exit(0)
// 				case "galaxy-a40s":
// 					pl("this is", model)
// 					pl("#####################")
// 					pl("it costs $324.00")
// 					pl("#####################")
// 					pl("are you sure you want to buy an ", model, ", for a price of $324.00 , (yes or no)")
// 					sc(&input)
// 					if input == "no" {
// 						pl("ok bye")
// 						os.Exit(0)
// 					}
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("please enter your bank account : ")
// 					sc(&account)
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("thank you for buying a", model)
// 					pl("bye see you next time ")
// 					pl("just kidding if you want accessories go back and go to accessories not in phones ")
// 					pl("now i go away")
// 					pl("now i go away")
// 					os.Exit(0)

// 				case "galaxy-s10":
// 					pl("this is", model)
// 					pl("#####################")
// 					pl("it costs $649")
// 					pl("#####################")
// 					pl("are you sure you want to buy an ", model, ", for a price of $649 , (yes or no)")
// 					sc(&input)
// 					if input == "no" {
// 						pl("ok bye")
// 						os.Exit(0)
// 					}
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("please enter your bank account : ")
// 					sc(&account)
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("thank you for buying a", model)
// 					pl("bye see you next time ")
// 					pl("just kidding if you want accessories go back and go to accessories not in phones ")
// 					pl("now i go away")

// 					os.Exit(0)

// 				case "galaxy-s20":
// 					pl("this is", model)
// 					pl("#####################")
// 					pl("it costs $649")
// 					pl("#####################")
// 					pl("are you sure you want to buy an ", model, ", for a price of $649 , (yes or no)")
// 					sc(&input)
// 					if input == "no" {
// 						pl("ok bye")
// 						os.Exit(0)
// 					}
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("please enter your bank account : ")
// 					sc(&account)
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("thank you for buying a", model)
// 					pl("bye see you next time ")
// 					pl("just kidding if you want accessories go back and go to accessories not in phones ")

// 					pl("now i go away")
// 					os.Exit(0)

// 				case "galaxy-s20-fe":
// 					pl("this is", model)
// 					pl("#####################")
// 					pl("it costs $ ")
// 					pl("#####################")
// 					pl("are you sure you want to buy an ", model, ", for a price of $629 , (yes or no)")
// 					sc(&input)
// 					if input == "no" {
// 						pl("ok bye")
// 						os.Exit(0)
// 					}
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("please enter your bank account : ")
// 					sc(&account)
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("thank you for buying a", model)
// 					pl("bye see you next time ")
// 					pl("just kidding if you want accessories go back and go to accessories not in phones ")

// 					pl("now i go away")
// 					os.Exit(0)

// 				case "galaxy-s20-ultra":
// 					pl("this is", model)
// 					pl("#####################")
// 					pl("it costs $720")
// 					pl("#####################")
// 					pl("are you sure you want to buy an ", model, ", for a price of $720 , (yes or no)")
// 					sc(&input)
// 					if input == "no" {
// 						pl("ok bye")
// 						os.Exit(0)
// 					}
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("please enter your bank account : ")
// 					sc(&account)
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("thank you for buying a", model)
// 					pl("bye see you next time ")
// 					pl("just kidding if you want accessories go back and go to accessories not in phones ")

// 					pl("now i go away")
// 					os.Exit(0)

// 				case "galaxy-s21":
// 					pl("this is", model)
// 					pl("#####################")
// 					pl("it costs $ 849")
// 					pl("#####################")
// 					pl("are you sure you want to buy an ", model, ", for a price of $ 849 , (yes or no)")
// 					sc(&input)
// 					if input == "no" {
// 						pl("ok bye")
// 						os.Exit(0)
// 					}
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("please enter your bank account : ")
// 					sc(&account)
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("thank you for buying a", model)
// 					pl("bye see you next time ")
// 					pl("just kidding if you want accessories go back and go to accessories not in phones ")

// 					pl("now i go away")
// 					os.Exit(0)

// 				case "galaxy-s21-ultra":
// 					pl("this is", model)
// 					pl("#####################")
// 					pl("it costs $1.169")
// 					pl("#####################")
// 					pl("are you sure you want to buy an ", model, ", for a price of $1.169 , (yes or no)")
// 					sc(&input)
// 					if input == "no" {
// 						pl("ok bye")
// 						os.Exit(0)
// 					}
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("please enter your bank account : ")
// 					sc(&account)
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("thank you for buying a", model)
// 					pl("bye see you next time ")
// 					pl("just kidding if you want accessories go back and go to accessories not in phones ")

// 					pl("now i go away")
// 					os.Exit(0)

// 				case "galaxy-s22":
// 					pl("this is", model)
// 					pl("#####################")
// 					pl("it costs $725,99")
// 					pl("#####################")
// 					pl("are you sure you want to buy an ", model, ", for a price of $725,99 , (yes or no)")
// 					sc(&input)
// 					if input == "no" {
// 						pl("ok bye")
// 						os.Exit(0)
// 					}
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("please enter your bank account : ")
// 					sc(&account)
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("thank you for buying a", model)
// 					pl("bye see you next time ")
// 					pl("just kidding if you want accessories go back and go to accessories not in phones ")

// 					pl("now i go away")
// 					os.Exit(0)

// 				case "galaxy-s22-ultra":
// 					pl("this is", model)
// 					pl("#####################")
// 					pl("it costs $787")
// 					pl("#####################")
// 					pl("are you sure you want to buy an ", model, ", for a price of $787 , (yes or no)")
// 					sc(&input)
// 					if input == "no" {
// 						pl("ok bye")
// 						os.Exit(0)
// 					}
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("please enter your bank account : ")
// 					sc(&account)
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("thank you for buying a", model)
// 					pl("bye see you next time ")
// 					pl("just kidding if you want accessories go back and go to accessories not in phones ")

// 					pl("now i go away")
// 					os.Exit(0)

// 				case "galaxy-s23":
// 					pl("this is", model)
// 					pl("#####################")
// 					pl("it costs $976,00")
// 					pl("#####################")
// 					pl("are you sure you want to buy an ", model, ", for a price of $976,00 , (yes or no)")
// 					sc(&input)
// 					if input == "no" {
// 						pl("ok bye")
// 						os.Exit(0)
// 					}
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("please enter your bank account : ")
// 					sc(&account)
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("thank you for buying a", model)
// 					pl("bye see you next time ")
// 					pl("just kidding if you want accessories go back and go to accessories not in phones ")
// 					pl("now i go away")
// 					os.Exit(0)

// 				case "galaxy-s23-ultra":
// 					pl("this is", model)
// 					pl("#####################")
// 					pl("it costs $1.087")
// 					pl("#####################")
// 					pl("are you sure you want to buy an ", model, ", for a price of $1.087 , (yes or no)")
// 					sc(&input)
// 					if input == "no" {
// 						pl("ok bye")
// 						os.Exit(0)
// 					}
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("please enter your bank account : ")
// 					sc(&account)
// 					pl("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 					pl("thank you for buying a", model)
// 					pl("bye see you next time ")
// 					pl("just kidding if you want accessories go back and go to accessories not in phones ")
// 					pl("now i go away")

// 					os.Exit(0)
// 				default:
// 					fmt.Println("something went wrong")
// 					os.Exit(0)
// 				}
// 			case "oppo":
// 				pl("this is", phone)
// 			case "oneplus":
// 				pl("this is", phone)
// 			case "vivo":
// 				pl("this is", phone)
// 			case "huawei":
// 				pl("this is", phone)
// 			case "nothing":
// 				pl("this is", phone)
// 			default:
// 				pl("something went wrong")
// 			}
// 		case "laptops":
// 			pl("this are laptops")
// 		case "pcs":
// 			pl("this are pcs")
// 		case "pc-components":
// 			pl("this are pc componets")
// 		case "camera":
// 			pl("this are camera")
// 		case "headphones":
// 			pl("this are headphones")
// 		case "mics":
// 			pl("this are mics")
// 		case "tech-furnitures":
// 			pl("this are tech furnitures")
// 		case "cables":
// 			pl("this are cables")
// 		case "machines":
// 			pl("this are machines")
// 		case "eletronic-home-gadgets":
// 			pl("this are eletronic home gadgets")
// 		case "screwdriver-ltt":
// 			pl("this are screwdriver ltt")
// 		case "screwdriver-jsr":
// 			pl("this are screwdriver jsr")
// 		case "fitness-eletronics":
// 			pl("this are fitness eletronics")
// 		case "mouse-and-keyboard":
// 			pl("this are mouse and keyboard")
// 		case "accessories":
// 			pl("this are accessories")
// 		default:
// 			pl("something went wrong ")
// 		}
// 	}
// }

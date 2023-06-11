package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// const conferenceTickets int = 50

// var conferenceName = "Go Conference"
// var remainingTickets uint = 50
// var bookings = make([]UserData, 0)

// type UserData struct {
// 	firstName       string
// 	lastName        string
// 	email           string
// 	numberOfTickets uint
// }

// var wg = sync.WaitGroup{}

// func main() {

// 	greetUsers()

// 	// for {
// 	firstName, lastName, email, userTickets := getUserInput()
// 	isValidName, isValidEmail := ValidateUserInput(firstName, lastName, email)

// 	if isValidName && isValidEmail {

// 		bookTicket(userTickets, firstName, lastName, email)

// 		wg.Add(1)
// 		go sendTicket(userTickets, firstName, lastName, email)

// 		firstNames := getFirstNames()
// 		fmt.Printf("The first names of bookings are: %v\n", firstNames)

// 		if remainingTickets == 0 {
// 			// end program
// 			fmt.Println("Our conference is booked out. Come back next year.")
// 			// break
// 		}
// 	} else {
// 		if !isValidName {
// 			fmt.Println("first name or last name you entered is too short")
// 		}
// 		if !isValidEmail {
// 			fmt.Println("email address you entered doesn't contain @ sign")
// 		}

// 	}
// 	//}
// 	wg.Wait()
// }

// func greetUsers() {
// 	fmt.Printf("Welcome to %v booking application\n", conferenceName)
// 	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
// 	fmt.Println("Get your tickets here to attend")
// }

// func getFirstNames() []string {
// 	firstNames := []string{}
// 	for _, booking := range bookings {
// 		firstNames = append(firstNames, booking.firstName)
// 	}
// 	return firstNames
// }

// func getUserInput() (string, string, string, uint) {
// 	var firstName string
// 	var lastName string
// 	var email string
// 	var userTickets uint

// 	fmt.Println("Enter your first name: ")
// 	fmt.Scan(&firstName)

// 	fmt.Println("Enter your last name: ")
// 	fmt.Scan(&lastName)

// 	fmt.Println("Enter your email address: ")
// 	fmt.Scan(&email)

// 	fmt.Println("Enter number of tickets: ")
// 	fmt.Scan(&userTickets)

// 	return firstName, lastName, email, userTickets
// }

// func bookTicket(userTickets uint, firstName string, lastName string, email string) {
// 	remainingTickets = remainingTickets - userTickets

// 	var userData = UserData{
// 		firstName:       firstName,
// 		lastName:        lastName,
// 		email:           email,
// 		numberOfTickets: userTickets,
// 	}

// 	bookings = append(bookings, userData)
// 	fmt.Printf("List of bookings is %v\n", bookings)

// 	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
// 	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
// }

// func sendTicket(userTickets uint, firstName string, lastName string, email string) {
// 	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
// 	time.Sleep(05 * time.Second)
// 	fmt.Println("#################")
// 	time.Sleep(03 * time.Second)
// 	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
// 	time.Sleep(02 * time.Second)
// 	fmt.Println("#################")

// 	fmt.Println("do you want to enter are newsletter (yes or no)")
// 	var yesNo = ""
// 	fmt.Scan(&yesNo)
// 	if yesNo == "no" {
// 		fmt.Println("ok bye see you next time")

// 	} else {
// 		var phoneNumber = ""

// 		fmt.Println("then give us your phone number so we will enter you in the newsletter")

// 		fmt.Scan(&phoneNumber)
// 		time.Sleep(500 * time.Millisecond)
// 		fmt.Printf("+ 39 %v \n", phoneNumber)

// 		fmt.Println("thank you so much for entering the newsletter ")
// 		time.Sleep(500 * time.Millisecond)
// 		fmt.Println("byee! see you next time")
// 	}
// 	time.Sleep(600 * time.Millisecond)
// 	fmt.Println("⚠️ ⚠️ ⚠️", "", "system shut down", "⚠️ ⚠️ ⚠️")

// 	wg.Done()
// }

package main

import (
	"fmt"
	"sync"
	"time"
)

const concertTickets int = 350

var concertName = "Kidung Jagad"
var RemainingTickets uint = 350
var bookings = make([]UserData, 0)

type UserData struct {
	firstName      string
	lastName       string
	email          string
	ticketsAmmount uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	firstName, lastName, email, userTickets := getUserInput()

	isValidName, isValidEmail, isValidAmmount := ValidateUserInput(firstName, lastName, email, userTickets, RemainingTickets)

	if isValidName && isValidEmail && isValidAmmount {

		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()
		fmt.Printf("Book lists: %v\n", firstNames)

		if RemainingTickets == 0 {
			fmt.Println("Tickets are sold out! See you at the concert!")
		}

	} else {
		if !isValidName {
			fmt.Println("First name or last name you entered is too short!")
		}
		if !isValidEmail {
			fmt.Println("Email address is invalid!")
		}
		if !isValidAmmount {
			fmt.Println("Number of tickets you book is invalid!")
		}
	}
	wg.Wait()
}

func greetUsers() {
	fmt.Println("ITS Student Choir Presents")
	fmt.Printf("%v: A Parade for 50th Years of ITS Student Choir\n", concertName)
	fmt.Printf("There is %v available, limited for %v Person\n", RemainingTickets, concertTickets)
	fmt.Printf("Grab your tickets now!\n")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	//User Input
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	RemainingTickets -= userTickets

	// Map for user
	var userData = UserData{
		firstName:      firstName,
		lastName:       lastName,
		email:          email,
		ticketsAmmount: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Println("Thank you for", firstName, lastName, "for booking", userTickets, "Tickets. You will receive a confirmation email at", email)
	fmt.Println(RemainingTickets, "tickets remaining for", concertName)
}

func sendTicket(userTickets uint, firsName string, lastName string, email string) {
	time.Sleep(30 * time.Second)

	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firsName, lastName)
	fmt.Println("###############################")
	fmt.Printf("Sending ticket:\n %v \n to %v\n", ticket, email)
	fmt.Println("###############################")
	wg.Done()
}

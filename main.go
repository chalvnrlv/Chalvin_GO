package main

import (
	"fmt"
	"strings"
)

func main() {
	concertName := "Kidung Jagad"
	const concertTickets int = 350
	var remainingTickets uint = 350
	bookings := []string{}

	fmt.Println("ITS Student Choir Presents")
	fmt.Printf("%v: A Parade for 50th Years of ITS Student Choir\n", concertName)
	fmt.Printf("There is %v available, limited for %v Person\n", remainingTickets, concertTickets)
	fmt.Printf("Grab your tickets now!\n")

	for {
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

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidAmmount := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidAmmount {
			remainingTickets -= userTickets
			bookings = append(bookings, firstName+" "+lastName)

			// fmt.Printf("Whole slice: %v\n", bookings)
			// fmt.Printf("First val: %v\n", bookings[0])
			// fmt.Printf("slice type: %T\n", bookings)
			// fmt.Printf("slice Len: %v\n", len(bookings))

			fmt.Println("Thank you for", firstName, lastName, "for booking", userTickets, "Tickets. You will receive a confirmation email at", email)
			fmt.Println(remainingTickets, "tickets remaining for", concertName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("Book lists: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Tickets are sold out! See you at the concert!")
				break
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
	}

}

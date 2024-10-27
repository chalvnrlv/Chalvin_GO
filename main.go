package main

import "fmt"

func main() {
	var concertName = "Kidung Jagad"
	const concertTickets = 350
	var remainingTickets = 350

	fmt.Println("ITS Student Choir Presents")
	fmt.Printf("%v: A Parade for 50th Years of ITS Student Choir\n", concertName)
	fmt.Printf("There is %v available, limited for %v Person\n", remainingTickets, concertTickets)
	fmt.Println("Grab your tickets now!")

	var userName string
	var userTickets int

	userName = "Maya"
	userTickets = 5
	fmt.Println(userName, "Have bought", userTickets, "Tickets")

}

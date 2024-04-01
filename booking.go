package main

import (
	"fmt"
	"strings"
)

func main() {
	var fightName string = "Floyd V Conor"
	const fightTickets int = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Println("Welcome to the Boxing Show booking app")
	fmt.Printf("Bookings for the fight %v are now open!!\n", fightName)
	fmt.Printf("Out of %v There are %v tickets left!!\n", fightTickets, remainingTickets)

	for remainingTickets > 0 {
		var firstName string
		var lastName string
		var userTickets uint
		var email string

		fmt.Println("Enter your first name")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name")
		fmt.Scan(&lastName)
		fmt.Println("Enter the number of tickets you want to book")
		fmt.Scan(&userTickets)
		fmt.Println("Enter the email you want to recieve the tickets")
		fmt.Scan(&email)

		var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
		var isValidEmail bool = strings.Contains(email, "@")
		var isValidTicket bool = userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicket {

			bookings = append(bookings, firstName+" "+lastName)

			remainingTickets = remainingTickets - userTickets

			fmt.Printf("Thank you %v %v for booking %v tickets, an email will be sent to %v for confirmation\n", firstName, lastName, userTickets, email)
			fmt.Printf("Out of %v There are %v tickets left!!\n", fightTickets, remainingTickets)

			var firstNames []string
			for _, name := range bookings {
				var names = strings.Fields(name)
				firstNames = append(firstNames, names[0])
			}

			fmt.Println("The people who have booked the tickets are", firstNames)

			if remainingTickets == 0 {
				fmt.Println("The tickets to the fight have been booked out, Browse other fights for tickets")
				break
			}

		} else {
			if !isValidName {
				fmt.Printf("First name or Last name is an invalid input, Please try again!!")
			}
			if !isValidTicket {
				fmt.Printf("We only have %v tickets available, %v tickets can not be booked!! Please try again\n", remainingTickets, userTickets)
			}
			if !isValidEmail {
				fmt.Printf("Please enter a valid email address")
			}
		}
	}
}

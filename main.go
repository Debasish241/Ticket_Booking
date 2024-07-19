package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go conference Mission"
	const confirmationTickets = 50
	var remainingTickets uint = 30
	var bookings = []string{}
	// making slice OR
	// bookings := []string{}

	fmt.Printf("conferenceTickets is %T , remainingTickets is %T , conferenceName is %T\n", confirmationTickets, remainingTickets, conferenceName)

	fmt.Println(conferenceName)
	fmt.Printf("Welcome to our %v ticketing application\n", conferenceName)
	fmt.Printf("We have total of  %v tickets and %v remained .\n", remainingTickets, confirmationTickets)
	fmt.Println("Get Your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email:")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		// isInValidCity:= city != "singapore"|| city != "London"

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank You %v %v for booking %v tickets. You will be received a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining \n", remainingTickets)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("All tickets have been sold out")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address is not comntain @ sign ")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets should be between 1 and 30")
			}

		}

	}

}

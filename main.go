package main

import (
	"booking_app/helper"
	"fmt"
	"time"
)

var conferenceName = "Go conference Mission"

const confirmationTickets = 50

var RemianingTickets uint = 30

// var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {
	greetUser()

	fmt.Printf("conferenceTickets is %T , RemianingTickets is %T , conferenceName is %T\n", confirmationTickets, RemianingTickets, conferenceName)

	fmt.Println(conferenceName)

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, RemianingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(RemianingTickets, lastName, firstName, userTickets, email)
			sendTicket(userTickets, lastName, email, firstName)
			firstNames := getFirstNames(bookings)
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if RemianingTickets == 0 {
				fmt.Println("All tickets have been sold out")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address does not contain '@'")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets should be between 1 and 30")
			}
		}
	}

	// Example of switch statement with commented lines
	// city := "London"
	// switch city {
	// case "New York":
	// 	// execute for New York booking system
	// case "Singapore":
	// 	// execute for Singapore booking system
	// case "Berlin", "London":
	// 	// execute for Berlin and London booking system
	// case "Mexico City":
	// 	// execute for Mexico City booking system
	// case "Hong Kong":
	// 	// execute for Hong Kong booking system
	// default:
	// 	fmt.Println("Invalid city")
	// }
}

func greetUser() {
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of  %v tickets and %v remained.\n", confirmationTickets, RemianingTickets)
	fmt.Println("Get Your tickets here to attend")
}

func getFirstNames(bookings []UserData) []string {
	firstNames := make([]string, len(bookings))
	for i, booking := range bookings {
		firstNames[i] = booking.firstName
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookTicket(remainingTickets uint, lastName string, firstName string, userTickets uint, email string) {
	RemianingTickets = RemianingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	// userData["firstname"] = firstName
	// userData["lastname"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("list of bookings %v\n", bookings)
	fmt.Printf("Thank You %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining \n", RemianingTickets)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var tickets = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Printf("############\n")
	fmt.Printf("Sending ticket to %v the email address: %v\n", tickets, email)
	fmt.Printf("############\n")
}

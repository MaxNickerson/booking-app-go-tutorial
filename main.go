package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	bookings := []string{}

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// asking user for their name
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email adress: ")
		fmt.Scan(&email)
		// fmt.Println(remainingTickets)
		// fmt.Println(&remainingTickets)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		// isValidCity := city == "Singapore" || city == "London"

		if isValidName && isValidEmail && isValidTicketNumber {

			remainingTickets = remainingTickets - userTickets
			// bookings[0] = firstName + " " + lastName
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("The whole slice: %v\n", bookings)
			fmt.Printf("The first value: %v\n", bookings[0])
			fmt.Printf("Slice type: %T\n", bookings)
			fmt.Printf("Slice length: %v\n", len(bookings))

			fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			printFirstNames(bookings)
		} else {
			if !isValidName {
				fmt.Println("first of last name entered is too short")
			}

			if !isValidEmail {
				fmt.Println("email adress does not contain @")
			}

			if !isValidTicketNumber {
				fmt.Println("number of tickets entered is invalid")
			}
			// fmt.Printf("Your input data is invalid, try again\n")

		}

	}

	// city := "London"

	// switch city {
	// case "New York":
	// 	// execute code for booking New York conference tickets
	// case "Singapore":
	// case "London":

	// default:
	// 	fmt.Print("No valid city selected")

	// }
}

func greetUsers(conferenceName string, conferenceTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

}

func printFirstNames(bookings []string) {
	firstNames := []string{}

	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	fmt.Printf("The first names of bookings are %v\n", firstNames)

}

package main

import (
	"fmt"
	"strings"
)

const conferenceTickets int = 50 //can not := using with const
var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = []string{}

func main() {

	greetUsers()
	//fmt.Printf("conferenceTickets is %T, remainingTikcets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	// compared to Java
	//Go --> int8,int16,int32,int64 | uInt: Positive, whole numbers
	//Java --> byte,short,int,long

	for {

		firstName, lastName, email, userTickets := getUserInput()
		fmt.Println(firstName)
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			// fmt.Printf("the whole slice: %v\n", bookings)
			// fmt.Printf("the first value: %v\n", bookings[0])
			// fmt.Printf("slice type: %T\n", bookings)
			// fmt.Printf("slice length: %v\n", len(bookings))

			bookTicket(userTickets, firstName, lastName, email)

			FirstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", FirstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Come back next year")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("firstname or lastname you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}

		}

	}

	// city := "London"

	// switch city{
	// 	case "New York":
	// 		// execute code for booking New York conference tickets
	// 	case "Singapore","Hong Kong":
	// 		// execute code for booking Singapore conference tickets
	// 	case "London","Berlin":
	// 		// execute code for booking London conference tickets
	// 	case "Mexico city":
	// 		// execute code for booking Mexico city conference tickets
	// 	default:
	// 		fmt.Println("No valid city selected")
	// }

}

func greetUsers() {
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their name
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address:")
	fmt.Scan(&email)

	fmt.Println("Enter your number of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickest remaining for %v\n", remainingTickets, conferenceName)
}

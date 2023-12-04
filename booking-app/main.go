package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50 //can not := using with const
var conferenceName = "Go Conference"
var RemainingTickets uint = 50

// var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()
	//fmt.Printf("conferenceTickets is %T, remainingTikcets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	// compared to Java
	//Go --> int8,int16,int32,int64 | uInt: Positive, whole numbers
	//Java --> byte,short,int,long

	firstName, lastName, email, userTickets := getUserInput()
	fmt.Println(firstName)
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, RemainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		// fmt.Printf("the whole slice: %v\n", bookings)
		// fmt.Printf("the first value: %v\n", bookings[0])
		// fmt.Printf("slice type: %T\n", bookings)
		// fmt.Printf("slice length: %v\n", len(bookings))

		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		FirstNames := getFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", FirstNames)

		if RemainingTickets == 0 {
			// end program
			fmt.Println("Our conference is booked out. Come back next year")
			// break
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
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, RemainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		// firstNames = append(firstNames, booking["firstName"])
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
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
	RemainingTickets = RemainingTickets - userTickets

	// create a map for a user
	// var userData = make(map[string]string)
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickest remaining for %v\n", RemainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###########################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("###########################")
	wg.Done()
}

/// 3 Levels of Scope
///Local--> Declaration within function or block (can be used only within that function or block)
///Package --> Declaration outside all function (can be used everywhere in the same package)
/// Global --> Declarartion outside all function & uppercase first letter (cna be used everywhere across all packages)

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

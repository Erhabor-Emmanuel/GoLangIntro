package main

import (
	"booking-app/shared"
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go conference"

const conferenceTicket int = 50

var remainingTicket uint = 50

// Slice
// However, we might want to define a size then along the way the size of the list will adjust based
// on the contents in it
// var bookings = make([]map[string]string, 0)
// This then creates an empty list of userData struct
var bookings = make([]UserData, 0)

// Struct
// the type there means we are creating a custom type in our application called userDetails
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	fmt.Println("----------using fmt.Println statement ------------")

	fmt.Println("----------using fmt.Printf statement ------------")
	fmt.Printf("conferenceName is %T and conferenceTicket is %T\n", conferenceName, conferenceTicket)
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("we have a total of %v tickets and %v are still remaining\n", conferenceTicket, remainingTicket)
	fmt.Println("Get your ticket here to attend")

	fmt.Println("---------- Using Type specific, Scan, and Pointer, Arrays ------------")
	//Pointers are variables that points to the memory address of another variable.
	var compName string
	//And this is acheived by adding the "&" sign to the variable
	fmt.Println(&compName)

	//An array is a combination of the size of the elements and the type of the elements is going to contain
	var booked [50]string
	//Assigning elements based on index positions
	booked[2] = "Nana"

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := shared.ValidateUserInput(firstName, lastName, email, userTickets, remainingTicket)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTickets(userTickets, firstName, lastName, email)

		fmt.Printf("User:  %v %v with email: %v booked %v tickets\n", firstName, lastName, email, userTickets)

		//call function print first names
		firstNames := getFirstNames()

		fmt.Printf("The firstnames of the booking list are %v\n", firstNames)
		fmt.Printf("These are all our bookings: %v\n", bookings)

		noTickets := remainingTicket == 0
		if noTickets {
			//end program
			fmt.Println("we have no tickets remaining")
			// break
		}
	} else {
		if !isValidName {
			fmt.Println("first name or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("email address you entered doesn't contain an @ sign")
		}
		if !isValidTicketNumber {
			fmt.Println("first name or last name you entered is too short")
		}
		// fmt.Println("Your input data is invalid, try again")
		// // fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTicket, userTickets)
		// continue
	}
	wg.Wait()

	fmt.Printf("%v tickets remaining for %v\n", remainingTicket, conferenceName)

	//Switch Statements
	city := "London"

	switch city {
	case "New York":
		//execute code for booking New York conference tickets
	case "Singapore", "Hong Kong":
		//execute code for booking Singapore and Hong Kong conference tickets
	case "London", "Berlin":
		//execute code for booking London and Berlin conference tickets
	case "Mexico":
		//execute code for booking New York conference tickets
	default:
		fmt.Println("No valid city selected")
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("Welcome to our", conferenceName, "booking application")
	fmt.Println("we have a total of", conferenceTicket, "tickets and", remainingTicket, "are still remaining")
	fmt.Println("Get your ticket here to attend")
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
	//ask user for name
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("How many tickets:")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTicket = remainingTicket - userTickets

	//create a map for a user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	//Append
	//adds the elements at the end of the slice
	//Grows the slice if a greater capacity is needed
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("The whole slice: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("Slice type: %T\n", bookings)
	fmt.Printf("Slice length: %v\n", len(bookings))
}

func sendTickets(userTickets uint, firstName string, lastNmae string, email string) {
	time.Sleep(10 * time.Second)
	//Sprintf===> helps you put together a string but instead of printing it out you
	//can save it into a string variable
	var ticket = fmt.Sprintf("%v  tickets for %v %v", userTickets, firstName, lastNmae)
	fmt.Println("################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("################")
	wg.Done()
}

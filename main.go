package main

import (
	"fmt"
	"sync"
	"time"
)

var conferenceName string = "Go Conference"

const totalEventTickets int = 50

var ticketsLeft uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstname       string
	lastname        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetConferenceAttendees()

	for {
		firstname, lastname, email, ticketQty := getUserData()
		isValidTicketNUmber, isUsernameValid, isEmailValid := userInputValidator(firstname, lastname, email, ticketQty)

		if isUsernameValid && isEmailValid && isValidTicketNUmber {

			bookEventTicket(firstname, lastname, email, ticketQty)

			wg.Add(1)
			go sendTicket(ticketQty, firstname, lastname, email)

			attendees := getAttendeesName()
			fmt.Printf("The first names of bookings are: %v\n", attendees)

			if ticketsLeft == 0 {
				fmt.Println("Our conference is booked out. Please come back next year.")
				break
			}
		} else {
			if !isUsernameValid {
				fmt.Println("Your first and last name is too short. Enter a valid firstname and lastname with at least 3 characters")
			}

			if !isEmailValid {
				fmt.Println("Your email is invalid. Enter a valid email")
			}

			if !isValidTicketNUmber {
				fmt.Println("Number of ticket you want to buy is invalid. Enter a number between 1 -  50")
			}

		}
	}
	wg.Wait()
}

func greetConferenceAttendees() {
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", totalEventTickets, ticketsLeft)
	fmt.Println("Get your tickets to the event here to attend. Thank You!")
}

func getAttendeesName() []string {
	attendees := []string{}

	for _, booking := range bookings {

		attendees = append(attendees, booking.firstname)
	}

	return attendees
}

func getUserData() (string, string, string, uint) {

	var firstname string
	var lastname string
	var email string
	var ticketQty uint
	// ask user to enter their name
	fmt.Println("Enter your firstname:...")
	fmt.Scan(&firstname)

	fmt.Println("Enter your last name:...")
	fmt.Scan(&lastname)

	fmt.Println("Enter your email address:...")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:...")
	fmt.Scan(&ticketQty)

	return firstname, lastname, email, ticketQty
}

func bookEventTicket(firstname string, lastname string, email string, ticketQty uint) []string {
	ticketsLeft = ticketsLeft - ticketQty

	// create s map for the user
	var userData = UserData{
		firstname:       firstname,
		lastname:        lastname,
		email:           email,
		numberOfTickets: ticketQty,
	}
	// userData["firstname"] = firstname
	// userData["lastname"] = lastname
	// userData["email"] = email
	// userData["ticketQty"] = fmt.Sprintf("%v", ticketQty)
	// userData["number_tickets"] = strconv.FormatUint(uint64(ticketQty), 10)

	fmt.Println(userData)
	bookings = append(bookings, userData)

	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstname, lastname, ticketQty, email)
	fmt.Printf("%v tickets remaining for %v.\n", ticketsLeft, conferenceName)

	return nil
}

func sendTicket(ticketQty uint, firstname string, lastname string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", ticketQty, firstname, lastname)
	fmt.Println("#######################################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#######################################")
	wg.Done()
}

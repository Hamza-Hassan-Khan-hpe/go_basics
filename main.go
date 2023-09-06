package main

import (
	"fmt"
	"strings"
)

// https://www.youtube.com/watch?v=yyUHQIec83I
func main() {
	var conference = "Go Conference"
	const totalTickets = 50 // constant -- vale can be used and referenced but not altered
	var ticketsLeft uint = 50

	fmt.Println("Welcome,", conference)
	fmt.Println("Total tickets available ", totalTickets)
	fmt.Println("tickets left", ticketsLeft)

	// function used for printing formatted data, for the other formatting %s see Go Doc
	fmt.Printf("Welcome Formatted Conference Name %v ", conference) // %v is for values

	/* ---- Data Types ---- */

	var userName string
	var userTickets int
	// ask user for input

	userName = "Foo"
	userTickets = 2

	var remainingTickets uint = 1        //-1 //unsigned int i.e. Vale of this var cant be negative it will throw an error if so
	fmt.Printf("\n%v", remainingTickets) // by default, it will assign an int, but sometimes we need to explicitly
	// declare the datatype

	fmt.Println("\n"+"User: "+userName, "\nTickets:", userTickets)

	// %T is to see the datatype

	fmt.Printf("username is of type %T \nuserTickets is of type %T", userName, userTickets)

	// Syntactic Sugar of GO Lang

	foo := "I am a string"
	// but we cant use this when we are explicitly declaring a datatype
	fmt.Println("\n" + foo)

	var custName string
	println("What is the Customer first name: ")
	fmt.Scan(&custName)

	var numberOfTickersRequired uint
	println("Number of tickets required: ")
	fmt.Scan(&numberOfTickersRequired)

	println("Thanks")
	fmt.Printf("\n Customer's Name: %v", custName)
	fmt.Printf("\n Total Tickets Booked: %v", numberOfTickersRequired)

	ticketsLeft = ticketsLeft - numberOfTickersRequired
	fmt.Printf("\nTotal Tickets Left : %v", ticketsLeft)
	lists()

}

func lists() {
	/* ---- Array's and Slices ---- */
	//var bookings = [50]string{"foo", "Alice", "Bob"}
	// arrays are fixed size
	//we need to define size,datatype and array itself ({}) -  50 means, an array of 50 items
	// cant mix datatypes ??

	var bookings [50]string // can also declare an array like this
	bookings[0] = "Alice Kennedy"
	bookings[1] = "Bob Browne"
	bookings[2] = "Charlie Allen"
	bookings[3] = "Echo Lopez"

	fmt.Printf("\n%v", bookings)

	/* ---- Slices ---- */

	// an abstraction of array , that runs array underneath
	// expands as we add items
	//is dynamic and doesn't need a size definition
	// dont need to worry about index (can use append)
	// to define just dont add size
	var bookingsSlice []string // this is a slice
	firstNames := []string{}   // can also use this syntax

	bookingsSlice = append(bookingsSlice, "Alice Kennedy")
	bookingsSlice = append(bookingsSlice, "Bob Browne")
	bookingsSlice = append(bookingsSlice, "Charlie D")

	fmt.Printf("\n%v", bookingsSlice)
	fmt.Printf("\n%v", len(bookingsSlice))

	// With just bookings you will get an index out of range error
	// because bookings is an array of 50 elements
	// there are only 4 items in the booking array so when it reaches the 5 items, there is nothing there
	// hence at names[0] there is nothing at that Index, so it crashes
	for _, i := range bookingsSlice { // underscore is a black identifier, this means there is a var here that we are not using

		var names = strings.Fields(i) // returns a slice with two elements split on the space
		fmt.Printf("\n%v", names[0])
		firstNames = append(firstNames, names[0]) // index 0 is the first names since fields made it an array of fist and last names
	}
	fmt.Printf("\n%v", firstNames)
}

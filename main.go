package main

import "fmt" // fmt is used for formatting I/O

func main() { //giving entry point
	var confname string = "Go conference"
	const conftickets int = 50
	var remainedtick uint = 50
	var bookings [50]string //making an array size of array_data type

	//fmt.Printf("Total %T , reamining %T \n", conftickets, remainedtick) //%T is palceholder for the variable and prints the type of variable

	fmt.Println("Welcome to Booking ", confname) //comma for sepration betwween message and variable
	fmt.Println("Total tickets", conftickets, "Remaining", remainedtick)

	fmt.Println("Get ur ticktes here ")

	var usrname string
	var lstname string
	var email string
	var usrticket uint

	//asking for input
	fmt.Println("Enter your name")
	fmt.Scan(&usrname) // it will read and pointer will point to address of stored name and print that out
	fmt.Println("Enter ur last name")
	fmt.Scan(&lstname)
	fmt.Println("Enter ur email")
	fmt.Scan(&email)
	fmt.Println("Enter no of tickets required")
	fmt.Scan(&usrticket)

	//fmt.Println(remainedtick)

	//fmt.Println(&remainedtick) // pointer for finding memory address

	//usrname = "Tom"
	//usrticket = 2

	//logic
	remainedtick = remainedtick - usrticket //always keep same data type
	bookings[0] = usrname + " " + lstname   //storing values to array

	fmt.Printf("Array for booking : %v \n", bookings)
	fmt.Printf("The first value : %v \n", bookings)
	fmt.Printf("Array type:%T \n", bookings)        //type of an array
	fmt.Printf("Array length:%v \n", len(bookings)) //length of an array

	fmt.Println(usrname)
	fmt.Printf("User %v %v booked %v tickets. Will be reciving mail soon at %v \n", usrname, lstname, usrticket, email) //printf formats the message and require a %v for refering the variable
	fmt.Printf(" Remaining tickets : %v", remainedtick)

}

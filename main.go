package main

import (
	"fmt" // fmt is used for formatting I/O
	"strings"
)

func main() { //giving entry point
	var confname string = "Go conference"
	const conftickets int = 50
	var remainedtick uint = 50
	//var bookings [50]string //making an array size of array_data type

	var bookings []string // making slice

	//fmt.Printf("Total %T , reamining %T \n", conftickets, remainedtick) //%T is palceholder for the variable and prints the type of variable

	fmt.Println("Welcome to Booking ", confname) //comma for sepration betwween message and variable
	fmt.Println("Total tickets", conftickets, "Remaining", remainedtick)
	fmt.Println(" ")
	fmt.Println("Get ur ticktes here ")

	for remainedtick > 0 && len(bookings) < 50 { //for loop is only loop here
		var usrname string
		var lstname string
		var email string
		var usrticket uint
		//asking for input
		fmt.Println("Please Enter ur details here")
		fmt.Println("   -------------------------------------   ")
		fmt.Println(" ")
		fmt.Println("Enter your name")
		fmt.Scan(&usrname) // it will read and pointer will point to address of stored name and print that out
		fmt.Println("Enter ur last name")
		fmt.Scan(&lstname)
		fmt.Println("Enter ur email")
		fmt.Scan(&email)
		fmt.Println("Enter no of tickets required")
		fmt.Scan(&usrticket)

		//validating user input

		var isValidname = len(usrname) >= 2 && len(lstname) >= 2     //validating name size
		var isAt = strings.Contains(email, "@")                      //validating @ in the input for email
		var isValidtick = usrticket > 0 && usrticket <= remainedtick // checking tickets

		//var isValidcity = city == "Delhi" || city == "Mumbai"
		if isValidname && isValidtick && isAt {
			remainedtick = remainedtick - usrticket //always keep same data type

			//bookings[0] = usrname + " " + lstname   //storing values to array

			bookings = append(bookings, usrname+" "+lstname) // making dynamic list using slice

			//fmt.Printf("Array for booking : %v \n", bookings)

			fmt.Printf("Slice for booking : %v \n", bookings)

			fmt.Printf("The first value : %v \n", bookings[0]) //reteriving the value from array

			//fmt.Printf("Array type:%T \n", bookings)//type of an array

			fmt.Printf("Slice type:%T \n", bookings) //type of a slice

			//fmt.Printf("Array length:%v \n", len(bookings))    //length of an array

			fmt.Printf("Slice length:%v \n", len(bookings))

			//fmt.Println(usrname)

			fmt.Printf("User %v %v booked %v tickets. Will be reciving mail soon at %v \n", usrname, lstname, usrticket, email) //printf formats the message and require a %v for refering the variable
			fmt.Printf(" Remaining tickets : %v \n", remainedtick)
			usrnames := []string{}             //another way of assigning a variable and making loop to store each entry
			for _, booking := range bookings { // made a nested for , index for slice index , booking as variable to loop and bookings to define the range
				var names = strings.Fields(booking) // splits the strings with white space as separator and returns a slice with the split elements
				//var usrname=names[0] //storing the value to array
				usrnames = append(usrnames, names[0]) // adding an element in slice
			}
			//fmt.Printf("These are all our bookings : %v \n", bookings)
			fmt.Printf("The first names of bookings are : %v \n", usrnames)

			var notickremain bool = remainedtick == 0 // storing a boolean experssion into variables

			if notickremain {
				//Stopping the loop
				fmt.Println("Oops !  Tickets sold out come next year")
				break
			}

		} else {

			if !isValidname {
				fmt.Println(" ")
				fmt.Println("Your given name is invalid")
				fmt.Println(" ")
			}

			if !isAt {
				fmt.Println(" ")
				fmt.Println("You missed @ sign")
				fmt.Println(" ")
			}

			if !isValidtick {
				fmt.Println(" ")
				fmt.Println("Invalid no of tickets")
				fmt.Println(" ")
			}
			continue
		}

		//fmt.Println(remainedtick)

		//fmt.Println(&remainedtick) // pointer for finding memory address

		//usrname = "Tom"
		//usrticket = 2

		//logic

	}

}

//_ is used to ignore a variable

package main

import (
	"fmt"
)

// array not used frequently
// more freq - slices are array with additional functionalities

type Car struct {
   NoT   int
   Lux   bool 
   Seat  bool 
   Make  string 
   Model string 
   Year  int
}

func main() {

	var myStrings [3]string  //similarly for int = [3]int

	myStrings[0] = "cat" // 1 - int
	myStrings[1] = "dog" // 2
	myStrings[2] = "fish" // 3

	fmt.Println("First element in array is", myStrings[0])
    
	// struct

    // hardway >>>
	//var myCar Car
	//myCar.NoT = 4 
	//myCar.Lux = false 
	//myCar.Make = "VW"

	// shorthand >>
	myCar := Car{
		NoT: 4,
		Lux: true,
		Seat: true,
		Make: "Volvo",
		Model: "XC90",
		Year: 1989,
	}

	fmt.Printf("My car is a %d %s %s", myCar.Year, myCar.Make, myCar.Model)
}
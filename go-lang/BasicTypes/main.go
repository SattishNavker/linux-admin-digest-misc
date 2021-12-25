package main

// [1] Basic Types = numbers/strings/booleans
// [2] Aggregate Types = arrays/structs
// [3] Reference Types = Pointers/slices??/maps/functions/channels
// [4] Interface Type (can store anything here)

import "log"

var myInt int
//var myInt16 int16 // never-used these 3
//var myInt32 int32 // old-phone/computer
//var myInt64 int64 // 64arch

var myUint uint // unsinged int
var myFloat float32 
var myFloat64 float64 // really large

func main() {

	myInt = 10
	myUint = 20 // only +ve

	myFloat = 10.1
	myFloat64 = 100.1 

	log.Println(myInt, myUint, myFloat, myFloat64)

	myString := "Satt"
	log.Println(myString)
	myString = "Nav" // IMP : strings are immutable in go

	var myBool = true 
	myBool = false
	//myBool = 0/yellow // ERR will not work only true/false allowed
	log.Println(myBool)
}
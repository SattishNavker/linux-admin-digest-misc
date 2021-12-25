package main

import (
	"myapp/packageone"
)

var blockVar = "pkgLevel"

func main() {
	var myVar = "myVarLocal"

	packageone.PrintMe(blockVar,myVar)
}
package main

import (
	"fmt"
	"log"
	"github.com/eiannone/keyboard"
	"strconv"
)

func main () {
	
	//rune stands for single-chara not string
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

    // wont exec imm instead after main (current-func) - cleanup in end
	defer func() {
		_ = keyboard.Close()
	}()

	coffees := make(map[int]string)

	coffees[1] = "Cappucino"
	coffees[2] = "Latte"
	coffees[3] = "Americano"
	coffees[4] = "Mocha"
	coffees[5] = "Macchiato"
	coffees[6] = "Espresso"

	fmt.Println("MENU")
	fmt.Println("----")
	fmt.Println("1 - Cappucino")
	fmt.Println("2 - Latte")
	fmt.Println("3 - Americano")
	fmt.Println("4 - Mocha")
	fmt.Println("5 - Macchiato")
	fmt.Println("6 - Espresso")
	fmt.Println("Q - Quit")

	for {
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		if char == 'q' || char == 'Q' {
			break
		}

		// %q for rune, %d for int, %s for string
		i, _ := strconv.Atoi(string(char))
		fmt.Println(fmt.Sprintf("You chose %s", coffees[i]))

	}
	fmt.Println("Exiting...")
}
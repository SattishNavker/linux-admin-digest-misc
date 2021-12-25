package main

import ( 
	"fmt"
	"os"
	"strings"
	"strconv"
	"bufio"
	"github.com/eiannone/keyboard"
	"log"
)

var reader *bufio.Reader
type User struct {
	UserName string
	Age int
	FavNum float64
	OwnsADog bool
}

func main() {
	reader = bufio.NewReader(os.Stdin)

	var user User

	user.UserName = readString("What is your name")
	user.Age = readInt("How old are you ?")
	user.FavNum = readFloat("What is your fav num ?")
	user.OwnsADog = readDog("Do you have a dog ?")
	//fmt.Println("Your name is", userName, "and you are ", age, "years old" )
	//inefficient, err with int-age >> fmt.Println("Your name is " + userName + ". You are ", age, "years old" )
	//fmt.Println(fmt.Sprintf("Your name is %s. Your are %d years old.", userName, age))

	fmt.Printf("Your name is %s. Your are %d years old.Your fav num is %.2f. Dog %t\n", 
				user.UserName, 
				user.Age, 
				user.FavNum,
				user.OwnsADog)  // much faster
}

func prompt() {
	fmt.Print("-> ")
}

func readString(s string) string {
	for {
		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)
		if userInput == "" {
			fmt.Println("Please enter a value")
		} else {
			return userInput
		}
	}
}

func readInt(s string) int {
	for {
		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)
		num, err := strconv.Atoi(userInput)

		if err != nil {
			fmt.Println("Please enter a whole num")
		} else {
			return num
		}
    }
}

func readFloat(s string) float64 {
	for {
		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)
		num, err := strconv.ParseFloat(userInput, 64)

		if err != nil {
			fmt.Println("Please enter a num")
		} else {
			return num
		}
    }
}

func readDog(q string) bool {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	for {
		fmt.Println(q)
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}
		//if char == 'n' || char == 'N' {
		//	return false
		//}
		//return true

		if strings.ToLower(string(char)) != "y" && strings.ToLower(string(char)) == "Y" {
			fmt.Println("Please press y or n")
		} else if char == 'n' || char == 'N' {
			return false
		} else if char == 'y' || char == 'Y' {
			return true
		}
	}
}
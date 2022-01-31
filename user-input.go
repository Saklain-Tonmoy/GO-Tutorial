package main

import (
	"fmt"
	// the two packages below are for taking user input
	"bufio"
	"os"
	// the strconv package is for type conversion
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter your name : ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Print("Enter your email : ")
	scanner.Scan()
	email := scanner.Text()

	fmt.Print("Enter the year you were born : ")
	scanner.Scan()
	// converting string to int. By default all the user inputs are string
	year, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	fmt.Println("####################")
	fmt.Printf("Name : %v \n", name)
	fmt.Printf("Email : %v \n", email)
	fmt.Printf("Year : %v \n", year)

}

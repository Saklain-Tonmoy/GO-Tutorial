package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

type BankDetails struct {
	accountNo   string
	name        string
	accountType string
	balance     float64
}

func openAccount(account *BankDetails) {

	fmt.Print("Enter Account No: ")
	scanner.Scan()
	account.accountNo = scanner.Text()

	fmt.Print("Enter Account Type: ")
	scanner.Scan()
	account.accountType = scanner.Text()

	fmt.Print("Enter Name: ")
	scanner.Scan()
	account.name = scanner.Text()

	fmt.Print("Enter Balance: ")
	scanner.Scan()
	account.balance, _ = strconv.ParseFloat(scanner.Text(), 10)

}

func displayAccountDetails(account *BankDetails) {
	fmt.Println("###### Account Details ######")
	fmt.Printf("Name of account holder: %v \n", account.name)
	fmt.Printf("Account no: %v \n", account.accountNo)
	fmt.Printf("Account type: %v \n", account.accountType)
	fmt.Printf("Balance: %v \n", account.balance)
}

func deposit(account *BankDetails) {
	fmt.Print("Enter deposit amount: ")
	scanner.Scan()
	amount, _ := strconv.ParseFloat(scanner.Text(), 10)
	if amount > 0 {
		account.balance += amount
	} else {
		fmt.Printf("The deposit amount %v is invalid \n", amount)
	}
}

func withdrawl(account *BankDetails) {
	fmt.Print("Enter withdraw amount: ")
	scanner.Scan()
	amount, _ := strconv.ParseFloat(scanner.Text(), 10)

	if account.balance > amount {
		account.balance -= amount
	} else {
		fmt.Printf("Your balance is less than %v \t Transaction failed \n", amount)
	}

}

func main() {

	fmt.Print("How many number of accounts do you want to create? ")
	scanner.Scan()
	n, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	if n <= 0 {
		fmt.Printf("Invalid number of accounts %v", n)
	} else if n == 1 {
		account := BankDetails{}
		openAccount(&account)
		displayAccountDetails(&account)
	} else {
		accounts := []BankDetails{}

		for i := 0; i < int(n); i++ {
			accounts = append(accounts, BankDetails{})
			openAccount(&accounts[i])
			//displayAccountDetails(&accounts[i])
		}

		for i := 0; i < len(accounts); i++ {
			displayAccountDetails(&accounts[i])
		}

	}

	// var i int
	// loop:
	// 	for i >= 0 {
	// 		fmt.Print("Enter a value: ")
	// 		scanner.Scan()
	// 		i, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	// 		switch i {
	// 		case int(i) == int(1):
	// 			fmt.Println("Pressed ONE")
	// 		case 0:
	// 			fmt.Println("Pressed ZERO and that's why loop ")
	// 			break
	// 		default:
	// 			fmt.Println("Pressed ####")
	// 		}
	// 	}

	// for {
	// 	fmt.Print("Enter a value: ")
	// 	scanner.Scan()
	// 	a, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	// 	switch a {
	// 	case 1:
	// 		fmt.Println("Pressed ONE")
	// 	case 0:
	// 		fmt.Println("Pressed ZERO and that's why loop ")
	// 		break
	// 	default:
	// 		fmt.Println("Pressed ####")
	// 	}
	// }

	// account := BankDetails{}

	// openAccount(&account)

	// displayAccountDetails(&account)

}

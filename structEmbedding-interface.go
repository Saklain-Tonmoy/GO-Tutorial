package main

import (
	"bufio"
	"fmt"
	"os"
)

type MyInterface interface {
	details() string
}

type Author struct {
	name  string
	email string
}

type Book struct {
	title  string
	author Author
}

func (book *Book) details() string {
	return fmt.Sprintf("Book Name : %s , Author Name : %s , Author Email : %s", book.title, book.author.name, book.author.email)
}

func print(myInterface MyInterface) {
	fmt.Println(myInterface.details())
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	book := Book{}

	fmt.Println("Enter Book Title: ")
	scanner.Scan()
	title := scanner.Text()
	fmt.Println("Enter Author Name: ")
	scanner.Scan()
	authorName := scanner.Text()
	fmt.Println("Enter Author Email: ")
	scanner.Scan()
	authorEmail := scanner.Text()

	book.title = title
	book.author.name = authorName
	book.author.email = authorEmail

	test := []MyInterface{&book}

	fmt.Printf("%T\n", test)

	for _, t := range test {
		fmt.Println(t.details())
		print(t)
	}

}

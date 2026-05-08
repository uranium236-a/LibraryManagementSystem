package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Book struct {
	ID     int
	Name   string
	Author string
	Genre  string
	Rating float32
}

type Library struct {
	Books  map[int]Book
	NextID int
}

func main() {

	var lib Library

	filename := "library.json"

	lib.Books = make(map[int]Book)

	err := lib.loadLibrary(filename)

	if err != nil {
		fmt.Println("Starting with new library...")
	}

	fmt.Println("=========Library Management System===========")

	reader := bufio.NewReader(os.Stdin)
	var command string

	for {
		fmt.Print("\n1. Add Book\n2. Remove Book\n3. Show Book\n4. Help\n5. Exit\n>> ")
		command, _ = reader.ReadString('\n')
		command = strings.TrimSpace(command)

		if command == "5" {
			break
		}

		switch command {
		case "1":
			fmt.Print("Enter Book name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)
			fmt.Print("Enter Author name: ")
			author, _ := reader.ReadString('\n')
			author = strings.TrimSpace(author)
			fmt.Print("Enter Genre name: ")
			genre, _ := reader.ReadString('\n')
			genre = strings.TrimSpace(genre)
			fmt.Print("Enter Rating: ")
			rating, _ := reader.ReadString('\n')
			rating = strings.TrimSpace(rating)

			rat, err := strconv.ParseFloat(rating, 32)

			if err != nil {
				fmt.Println("Invalid rating")
				break
			}

			lib.addBook(name, author, genre, float32(rat))
			lib.saveLibrary(filename)

		case "2":
			fmt.Print("Enter Book ID: ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			id, err := strconv.Atoi(input)

			if err != nil {
				fmt.Println("Invalid input")
				break
			}
			lib.removeBook(id)
			lib.saveLibrary(filename)

		case "3":
			fmt.Print(">> ")
			comm, _ := reader.ReadString('\n')
			comm = strings.TrimSpace(comm)
			lib.showBook(comm, reader)

		case "4":
			fmt.Println("1. Add Book:\n\tEnter all asked input properly")
			fmt.Println("2. Remov Book:\n\tEnter Book Id to remove the Book")
			fmt.Println("3. Show Book/Books:\n\tall :- All books will be displayed\n\tname :- Books with containing same Name will be displayed\n\tauthor :- Books containing same Author name will be displayed\n\tgenre :- Books containing same Genre will be displayed")
			fmt.Println("5. Exit the CLI")

		default:
			fmt.Println("Invalid input")

		}
	}

}

func (lib *Library) saveLibrary(filename string) error {
	file, err := os.Create(filename)

	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	encoder.SetIndent("", " ")

	return encoder.Encode(lib)
}

func (lib *Library) loadLibrary(filename string) error {
	file, err := os.Open(filename)

	if err != nil {
		return err
	}
	defer file.Close()

	stat, _ := file.Stat()

	if stat.Size() == 0 {
		return nil
	}

	decoder := json.NewDecoder(file)

	return decoder.Decode(lib)
}

func (lib *Library) addBook(name, author, genre string, rating float32) {
	book := Book{
		ID:     lib.NextID,
		Name:   name,
		Author: author,
		Genre:  genre,
		Rating: rating,
	}

	lib.Books[lib.NextID] = book
	lib.NextID += 1

	fmt.Println("Book added successfully")
}

func (lib *Library) showBook(command string, reader *bufio.Reader) {

	switch command {
	case "all":
		fmt.Println("\nAll Books:")
		for id, b := range lib.Books {
			fmt.Println(id, "\t", b.Name, "\t", b.Author, "\t", b.Genre, "\t", b.Rating)
		}
	case "name", "author", "genre":

		fmt.Println("Enter ", command, " name: ")
		temp, _ := reader.ReadString('\n')
		temp = strings.TrimSpace(temp)

		if temp == "" {
			fmt.Println("Empty Input")
			return
		}

		for _, b := range lib.Books {
			var field string

			switch command {
			case "name":
				field = b.Name
			case "author":
				field = b.Author
			case "genre":
				field = b.Genre
			}

			if strings.Contains(strings.ToLower(field), strings.ToLower(temp)) {
				fmt.Println(b.ID, "\t", b.Name, "\t", b.Author, "\t", b.Genre, "\t", b.Rating)
			}
		}

	default:
		fmt.Println("Invalid command")
	}

}

func (lib *Library) removeBook(id int) {

	_, ok := lib.Books[id]

	if ok {
		delete(lib.Books, id)
		fmt.Println("Book removed successfully")
	} else {
		fmt.Println("Book doesn't exist")
	}

}

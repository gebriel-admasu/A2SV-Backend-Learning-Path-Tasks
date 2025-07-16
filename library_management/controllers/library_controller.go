package controllers

import (
	"bufio"
	"fmt"
	"library_management/models"
	"library_management/services"
	"os"
	"strconv"
	"strings"
)

func StartConsole(library *services.Library) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\nLibrary Management System")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books by Member")
		fmt.Println("7. Add Member")
		fmt.Println("0. Exit")
		fmt.Print("Choose an option: ")
		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)
		choice, _ := strconv.Atoi(choiceStr)

		switch choice {
		case 1:
			fmt.Print("Enter Book ID: ")
			idStr, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idStr))
			fmt.Print("Enter Title: ")
			title, _ := reader.ReadString('\n')
			fmt.Print("Enter Author: ")
			author, _ := reader.ReadString('\n')
			book := models.Book{ID: id, Title: strings.TrimSpace(title), Author: strings.TrimSpace(author), Status: "Available"}
			library.AddBook(book)
			fmt.Println("Book added.")
		case 2:
			fmt.Print("Enter Book ID to remove: ")
			idStr, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idStr))
			library.RemoveBook(id)
			fmt.Println("Book removed.")
		case 3:
			fmt.Print("Enter Book ID to borrow: ")
			bookIDStr, _ := reader.ReadString('\n')
			bookID, _ := strconv.Atoi(strings.TrimSpace(bookIDStr))
			fmt.Print("Enter Member ID: ")
			memberIDStr, _ := reader.ReadString('\n')
			memberID, _ := strconv.Atoi(strings.TrimSpace(memberIDStr))
			err := library.BorrowBook(bookID, memberID)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Book borrowed.")
			}
		case 4:
			fmt.Print("Enter Book ID to return: ")
			bookIDStr, _ := reader.ReadString('\n')
			bookID, _ := strconv.Atoi(strings.TrimSpace(bookIDStr))
			fmt.Print("Enter Member ID: ")
			memberIDStr, _ := reader.ReadString('\n')
			memberID, _ := strconv.Atoi(strings.TrimSpace(memberIDStr))
			err := library.ReturnBook(bookID, memberID)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Book returned.")
			}
		case 5:
			books := library.ListAvailableBooks()
			fmt.Println("Available Books:")
			for _, b := range books {
				fmt.Printf("ID: %d, Title: %s, Author: %s\n", b.ID, b.Title, b.Author)
			}
		case 6:
			fmt.Print("Enter Member ID: ")
			memberIDStr, _ := reader.ReadString('\n')
			memberID, _ := strconv.Atoi(strings.TrimSpace(memberIDStr))
			books := library.ListBorrowedBooks(memberID)
			fmt.Printf("Borrowed Books by Member %d:\n", memberID)
			for _, b := range books {
				fmt.Printf("ID: %d, Title: %s, Author: %s\n", b.ID, b.Title, b.Author)
			}
		case 7:
			fmt.Print("Enter Member ID: ")
			idStr, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idStr))
			fmt.Print("Enter Member Name: ")
			name, _ := reader.ReadString('\n')
			member := &models.Member{ID: id, Name: strings.TrimSpace(name)}
			library.Members[id] = member
			fmt.Println("Member added.")
		case 0:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option.")
		}
	}
}

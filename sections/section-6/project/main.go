package main

import "fmt"

type Contact struct {
	ID    int
	Name  string
	Email string
	Phone string
}

var contactList []Contact
var contactIndexByName map[string]int
var nextID int = 1

func init() {
	contactList = make([]Contact, 0) // initialize the map with a capacity of 0
	contactIndexByName = make(map[string]int)
}

func addContact(name string, email string, phone string) {
	if _, exists := contactIndexByName[name]; exists {
		fmt.Printf("Contact with name %s already exists\n", name)
		return
	}

	newContact := Contact{
		ID:    nextID,
		Name:  name,
		Email: email,
		Phone: phone,
	}
	nextID++
	contactList = append(contactList, newContact)
	contactIndexByName[name] = newContact.ID
	fmt.Printf("Contact %s added successfully\n", name)
}

/*
Alternative: func findContactByName(name string) Contact
With a value return you can’t use nil, so you need a second return to indicate “not found”.
*/
func findContactByName(name string) *Contact { // return a pointer to the contact
	index, ok := contactIndexByName[name]
	if !ok {
		return nil
	}
	return &contactList[index]
}

func main() {
	addContact("Alice Wonderland", "alice@example.com", "111-2222")
	addContact("Bob The Builder", "bob@example.com", "333-4444")
	addContact("Charlie Brown", "charlie@example.com", "555-6666")
	addContact("Alice Wonderland", "alice.new@example.com", "777-8888") // Attempt to add duplicate

	alice := findContactByName("Alice Wonderland")
	if alice == nil {
		fmt.Println("No Bob contact found.")
	} else {
		fmt.Println("Alice contact found.", alice.Name)
	}
	fmt.Println("--------------------------------")
	fmt.Printf("Name: %v\n", alice)
	fmt.Printf("Type: %T\n", alice)
}

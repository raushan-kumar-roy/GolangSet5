package main

import "fmt"

type Friend struct {
	Name         string
	MobileNumber string
	AltMobile    string
	Address      string
	City         string
	College      bool
}

func main() {
	friends := []Friend{
		{Name: "Alice", MobileNumber: "123-456-7890", AltMobile: "", Address: "123 Main St", City: "New York", CollegeFriend: true},
		{Name: "Bob", MobileNumber: "234-567-8901", AltMobile: "345-678-9012", Address: "456 Main St", City: "New York"},
		{Name: "Charlie", MobileNumber: "345-678-9012", AltMobile: "", Address: "789 Main St", City: "Chicago"},
		{Name: "Dave", MobileNumber: "456-789-0123", AltMobile: "567-890-1234", Address: "321 Main St", City: "Chicago", CollegeFriend: true},
		{Name: "Eve", MobileNumber: "567-890-1234", AltMobile: "", Address: "654 Main St", City: "Los Angeles"},
		{Name: "Frank", MobileNumber: "678-901-2345", AltMobile: "789-012-3456", Address: "987 Main St", City: "Los Angeles"},
		{Name: "Gina", MobileNumber: "789-012-3456", AltMobile: "", Address: "246 Main St", City: "San Francisco", CollegeFriend: true},
		{Name: "Henry", MobileNumber: "890-123-4567", AltMobile: "901-234-5678", Address: "369 Main St", City: "San Francisco"},
		{Name: "Izzy", MobileNumber: "901-234-5678", AltMobile: "", Address: "159 Main St", City: "Seattle"},
		{Name: "Jane", MobileNumber: "012-345-6789", AltMobile: "123-456-7890", Address: "753 Main St", City: "Seattle", CollegeFriend: true},
	}
	if CollegeFriend == true {
		for _, friend := range friends {
			fmt.Println(friend.Name, friend.MobileNumber)
		}
	} else {
		fmt.Println("not a college friend")
	}

}

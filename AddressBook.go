package main

import "fmt"

type Friend struct {
	Name         string
	MobileNumber string
	AltMobile    string
	Address      string
	City         string
}

func main() {
	friends := []Friend{
		{Name: "Alice", MobileNumber: "123-456-7890", AltMobile: "", Address: "123 Main St", City: "New York"},
		{Name: "Bob", MobileNumber: "234-567-8901", AltMobile: "345-678-9012", Address: "456 Main St", City: "New York"},
		{Name: "Charlie", MobileNumber: "345-678-9012", AltMobile: "", Address: "789 Main St", City: "Chicago"},
		{Name: "Dave", MobileNumber: "456-789-0123", AltMobile: "567-890-1234", Address: "321 Main St", City: "Chicago"},
		{Name: "Eve", MobileNumber: "567-890-1234", AltMobile: "", Address: "654 Main St", City: "Los Angeles"},
		{Name: "Frank", MobileNumber: "678-901-2345", AltMobile: "789-012-3456", Address: "987 Main St", City: "Los Angeles"},
		{Name: "Gina", MobileNumber: "789-012-3456", AltMobile: "", Address: "246 Main St", City: "San Francisco"},
		{Name: "Henry", MobileNumber: "890-123-4567", AltMobile: "901-234-5678", Address: "369 Main St", City: "San Francisco"},
		{Name: "Izzy", MobileNumber: "901-234-5678", AltMobile: "", Address: "159 Main St", City: "Seattle"},
		{Name: "Jane", MobileNumber: "012-345-6789", AltMobile: "123-456-7890", Address: "753 Main St", City: "Seattle"},
	}

	for _, friend := range friends {
		fmt.Println(friend.Name, friend.MobileNumber)
	}

	fmt.Println("\nFriends with alternate phone numbers:")
	for _, friend := range friends {
		if friend.AltMobile != "" {
			fmt.Println(friend.Name)
		}
	}

	cityCounts := make(map[string]int)
	for _, friend := range friends {
		cityCounts[friend.City]++
	}

	fmt.Println("\nNumber of friends in each city:")
	for city, count := range cityCounts {
		fmt.Println(city, count)
	}
}

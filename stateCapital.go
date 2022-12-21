package main

import "fmt"

func main() {
	stateCapital := make(map[string]string)

	stateCapital["Gujarat"] = "Gandhinagar"
	stateCapital["Maharashtra"] = "Mumbai"
	stateCapital["Karnataka"] = "Bengaluru"
	stateCapital["Tamil Nadu"] = "Chennai"
	stateCapital["Andhra Pradesh"] = "Amaravati"
	stateCapital["Telangana"] = "Hyderabad"
	stateCapital["Punjab"] = "Chandigarh"
	stateCapital["Haryana"] = "Chandigarh"
	stateCapital["Uttar Pradesh"] = "Lucknow"
	stateCapital["Bihar"] = "Patna"

	fmt.Printf("stateCapital is of type %T\n", stateCapital)

	fmt.Println("stateCapital:", stateCapital)

	for state, capital := range stateCapital {
		fmt.Printf("%s → %s\n", state, capital)
	}
}

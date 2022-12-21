package main

import (
	"fmt"
	"sort"
)

func main() {
	states := make(map[string]string)
	states["Delhi"] = "Delhi"
	states["Kerala"] = "Thiruvananthapuram"
	states["Tamil Nadu"] = "Chennai"
	states["Maharashtra"] = "Mumbai"
	states["Uttar Pradesh"] = "Lucknow"

	stateSlice := make([]string, 0, len(states))
	for state := range states {
		stateSlice = append(stateSlice, state)
	}

	sort.Strings(stateSlice)

	for _, state := range stateSlice {
		fmt.Printf("%s: %s\n", state, states[state])
	}
}

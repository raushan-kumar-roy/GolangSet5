package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	numberDistribution := make(map[string]int)

	for i := 0; i < 1000; i++ {
		num := rand.Intn(26)

		if num <= 5 {
			numberDistribution[">0 && <=5"]++
		} else if num <= 10 {
			numberDistribution[">5 && <=10"]++
		} else if num <= 15 {
			numberDistribution[">10 && <=15"]++
		} else if num <= 20 {
			numberDistribution[">15 && <=20"]++
		} else {
			numberDistribution[">20 && <=25"]++
		}
	}

	fmt.Println(numberDistribution)
}

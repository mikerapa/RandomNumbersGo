package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Application started")
	loopLimit := 20
	minValue := 1
	maxValue := 10
	fmt.Println("minValue: ", minValue, "maxValue: ", maxValue, "loopLimit: ", loopLimit)
	var results map[int]bool
	results = createResultsCollection(maxValue, minValue)

	for i := 0; i <= loopLimit; i++ {
		newNumber := randomizeInt(minValue, maxValue)
		fmt.Println("new number:", newNumber)
		results[newNumber] = true
	}

	showUnclaimedNumbers(results)
}

func randomizeInt(MinValue int, MaxValue int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(MaxValue-MinValue+1) + MinValue
}

func createResultsCollection(UpperBound int, LowerBound int) (results map[int]bool) {
	results = make(map[int]bool)

	for k := LowerBound; k <= UpperBound; k++ {
		results[k] = false
	}
	return
}

func showUnclaimedNumbers(Results map[int]bool) {
	fmt.Print("Unclaimed numbers:")

	for i, k := range Results {
		if !k {
			fmt.Print(i, " ")
		}
	}
	fmt.Println("")
}

package main

import (
	"fmt"
	"sync"

	"github.com/RGood/countdown/pkg/helpers"
	"github.com/RGood/countdown/pkg/numbers"
)

func main() {
	wg := &sync.WaitGroup{}

	// Numbers to play with
	startingNums := []int{100, 8, 4, 8, 9, 6}

	// Goal number
	target := 317

	// Data struct containing results
	results := helpers.NewStringSet()

	// Stop on first match
	getFirst := false

	for index, value := range startingNums {
		var numsCopy = make([]int, index)
		copy(numsCopy, startingNums[:index])
		otherNums := append(numsCopy, startingNums[index+1:]...)

		wg.Add(1)
		go numbers.CalcNums(value, fmt.Sprintf("%d", value), otherNums, target, results, getFirst, wg)
	}

	wg.Wait()

	for _, result := range results.Values() {
		println(result)
	}

	fmt.Printf("Num Results:    %d\n", results.Size())
	fmt.Printf("Num Iterations: %d\n", numbers.RunCount)
}

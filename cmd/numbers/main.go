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
	startingNums := []int{1, 10, 4, 5, 2, 5}

	// Goal number
	target := 423

	// Data struct containing results
	results := helpers.NewStringSet()

	// Stop on first match
	getFirst := false

	startingNPPs := helpers.GenNPPs(startingNums)
	wg.Add(1)
	go numbers.CalcNums(startingNPPs, target, results, getFirst, wg)

	wg.Wait()

	for _, result := range results.Values() {
		println(result)
	}

	fmt.Printf("Num Results:    %d\n", results.Size())
	fmt.Printf("Num Iterations: %d\n", numbers.RunCount)
}

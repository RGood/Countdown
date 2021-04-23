package main

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/RGood/countdown/pkg/helpers"
	"github.com/RGood/countdown/pkg/numbers"
)

func main() {
	wg := &sync.WaitGroup{}

	// Numbers to play with
	startingNums := []int{100, 75, 50, 25, 10, 9}

	// Goal number
	target := 819

	// Data struct containing results
	results := helpers.NewStringSet()

	// Stop on first match
	getFirst := false

	startTime := time.Now()
	startingNPPs := helpers.GenNPPs(startingNums)
	wg.Add(1)
	go numbers.CalcNums(startingNPPs, target, results, getFirst, wg)

	wg.Wait()
	elapsedTime := time.Now().Sub(startTime)

	fmt.Printf("Starting Numbers: %s\n", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(startingNums)), ", "), "[]"))
	fmt.Printf("Target: %d\n", target)

	for _, result := range results.Values() {
		println(result)
	}

	fmt.Printf("Num Results:    %d\n", results.Size())
	fmt.Printf("Num Iterations: %d\n", numbers.RunCount)
	fmt.Printf("Elapsed Time:   %fs\n", elapsedTime.Seconds())
}

package numbers

import (
	"fmt"
	"sync"

	"github.com/RGood/countdown/pkg/helpers"
)

var RunCount int = 0

func CalcNums(curVal int, curPath string, options []int, target int, results *helpers.StringSet, getFirst bool, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	RunCount++

	if getFirst && results.Size() > 0 {
		return
	}

	if curVal == target {
		results.Add(curPath)
	}

	for index, option := range options {
		var numsCopy = make([]int, index)
		copy(numsCopy, options[:index])
		new_options := append(numsCopy, options[index+1:]...)

		waitGroup.Add(1)
		go CalcNums(curVal*option, curPath+fmt.Sprintf(" * %d", option), new_options, target, results, getFirst, waitGroup)
		waitGroup.Add(1)
		go CalcNums(curVal+option, curPath+fmt.Sprintf(" + %d", option), new_options, target, results, getFirst, waitGroup)
		waitGroup.Add(1)
		go CalcNums(curVal-option, curPath+fmt.Sprintf(" - %d", option), new_options, target, results, getFirst, waitGroup)
		if option != 0 && curVal%option == 0 {
			waitGroup.Add(1)
			go CalcNums(curVal/option, curPath+fmt.Sprintf(" / %d", option), new_options, target, results, getFirst, waitGroup)
		}
	}

}

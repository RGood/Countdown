package numbers

import (
	"sync"

	"github.com/RGood/countdown/pkg/helpers"
)

var RunCount int = 0

func GetAllExcluding(vars []*helpers.NumPathPair, indexes ...int) []*helpers.NumPathPair {
	indexMap := map[int]struct{}{}
	for _, index := range indexes {
		indexMap[index] = struct{}{}
	}

	elements := []*helpers.NumPathPair{}
	for index, npp := range vars {
		_, ok := indexMap[index]
		if !ok {
			elements = append(elements, npp)
		}
	}

	return elements

}

func Mutate(options []*helpers.NumPathPair) [][]*helpers.NumPathPair {
	mutations := [][]*helpers.NumPathPair{}
	for i1, npp1 := range options {
		for i2 := i1 + 1; i2 < len(options); i2++ {
			npp2 := options[i2]
			mutations = append(mutations, append(GetAllExcluding(options, i1, i2), npp1.Add(npp2)))
			mutations = append(mutations, append(GetAllExcluding(options, i1, i2), npp1.Mul(npp2)))

			if npp1.Num-npp2.Num >= 0 {
				mutations = append(mutations, append(GetAllExcluding(options, i1, i2), npp1.Sub(npp2)))
			}

			if npp2.Num != 0 && npp1.Num%npp2.Num == 0 {
				mutations = append(mutations, append(GetAllExcluding(options, i1, i2), npp1.Div(npp2)))
			}
		}
	}

	return mutations
}

func CalcNums(options []*helpers.NumPathPair, target int, results *helpers.StringSet, getFirst bool, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	RunCount++

	if getFirst && results.Size() > 0 {
		return
	}

	for _, npp := range options {
		if npp.Num == target {
			results.Add(npp.Path)
		}
	}

	for _, permutation := range Mutate(options) {
		waitGroup.Add(1)
		go CalcNums(permutation, target, results, getFirst, waitGroup)
	}

}

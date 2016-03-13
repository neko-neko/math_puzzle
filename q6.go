package main

import(
	"flag"
	"fmt"
	"github.com/neko-neko/math_puzzle/common"
)

const MAX_NUMBER = 10000

func main() {
	var iterateCount int;
	flag.IntVar(&iterateCount, "iterate", 10, "iterate count")
	flag.Parse()

	common.ExecuteIteration(iterateCount, func() {
		var hit int
		for i := 2; i <= MAX_NUMBER; i++ {
			if i % 2 != 0 {
				continue
			}
			var (
				initNumber int = i
				number int = i * 3 + 1
			)

			if calculate(initNumber, number) {
				hit++
			}
		}
		fmt.Printf("%d\n", hit)
	})
}

func calculate(initNumber int, number int) bool {
	if number % 2 == 0 {
		number = number / 2
	} else {
		number = number * 3 + 1
	}

	if number == initNumber {
		return true
	} else if number == 1 {
		return false
	}

	return calculate(initNumber, number)
}

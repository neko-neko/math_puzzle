package main

import(
	"flag"
	"fmt"
	"github.com/neko-neko/math_puzzle/common"
)

const MAX_CARDS = 100

func main() {
	var iterateCount int;
	flag.IntVar(&iterateCount, "iterate", 10, "iterate count")
	flag.Parse()

	common.ExecuteIteration(iterateCount, func() {
		var cards [MAX_CARDS]bool
		for i := 2; i <= MAX_CARDS; i++ {
			var reverseCardIndex = i - 1
			for reverseCardIndex < MAX_CARDS {
				cards[reverseCardIndex] = !cards[reverseCardIndex]
				reverseCardIndex += i
			}
		}

		for idx, card := range cards {
			if !card {
				fmt.Printf("%d\n", idx + 1);
			}
		}
	})
}
package main

import(
	"flag"
	"fmt"
	"github.com/neko-neko/math_puzzle/common"
)

func main() {
	var iterateCount int;
	flag.IntVar(&iterateCount, "iterate", 10, "iterate count")
	flag.Parse()

	common.ExecuteIteration(iterateCount, func() {
		var board = [][]int{{1, 2}, {2, 3}, {7, 8}, {8, 9}, {1, 4}, {3, 6}, {4, 7}, {6, 9}, {1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}}
		var result int
		calculateBoardPattern(board, &result)
		fmt.Printf("%d\n", result)
	})
}

func calculateBoardPattern(board [][]int, pattern *int) {
	if len(board) == 0 {
		return
	}

	for _, result := range board {
		var next [][]int
		for _, num := range board {
			if len(num) == 1 && len(result) == 1 {
				if num[0] != result[0] {
					next = append(next, num)
				}
			} else if len(num) == 1 {
				if num[0] != result[0] && num[0] != result[1] {
					next = append(next, num)
				}
			} else if len(result) == 1 {
				if result[0] != num[0] && result[0] != num[1] {
					next = append(next, num)
				}
			} else {
				if result[0] != num[0] && result[0] != num[1] && result[1] != num[0] && result[1] != num[1] {
					next = append(next, num)
				}
			}
		}
		*pattern++
		calculateBoardPattern(next, pattern)
	}
}
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
		var result []int
		for i := 4; i <= 500; i += 4 {
			side := (i / 4)
			square := side * side

			var list []int
			for height := 1; height <= side - 1; height++ {
				list = append(list, height * (i / 2 - height))
			}
			uniquePatterns(list, 2, square, &result)
		}
		fmt.Printf("count:%d\n", len(result))
	})
}

func uniquePatterns(list []int, r int, square int, patterns *[]int) {
	n := len(list)

	if r > n {
		return
	}

	indexes := make([]int, r)
	for i := range indexes {
		indexes[i] = i
	}

	result := make([]int, r)
	for idx, v := range indexes {
		result[idx] = list[v]
	}

	for {
		i := r - 1
		for ; i >= 0 && indexes[i] == i+n-r; i -= 1 {
		}

		if i < 0 {
			return
		}

		indexes[i] += 1
		for j := i + 1; j < r; j += 1 {
			indexes[j] = indexes[j-1] + 1
		}

		for ; i < len(indexes); i += 1 {
			result[i] = list[indexes[i]]
		}

		if result[0] + result[1] == square {
			if div(result[0], result[1]) == 1 {
				*patterns = append(*patterns, result[0] + result[1])
			}
		}
	}
}

func div(a int, b int) int {
	if b == 0 {
		return a
	} else {
		return div(b, a % b)
	}
}
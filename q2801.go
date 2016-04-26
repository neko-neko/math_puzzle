package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// parse stdin
	var users [4][5][5]int
	var x int
	var y int
	var userIndex int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		slice := strings.Split(line, ",")
		for _, s := range slice {
			users[userIndex][y][x], _ = strconv.Atoi(s)
			x++
			if (x == 5) {
				x = 0
				y++
			}
			if (y == 5) {
				x = 0
				y = 0
				userIndex++
			}
		}
	}

	// bingo calculate
	bingoPatterns := createBingoPatterns(users)

	// calculate min pattern
	min := 75
	for i := 0; i < 12; i++ {
		for j := 0; j < 12; j++ {
			for k := 0; k < 12; k++ {
				for l := 0; l < 12; l++ {
					var bingoMap map[int]bool = make(map[int]bool)
					for m := 0; m < 5; m++ {
						bingoMap[bingoPatterns[0][i][m]] = true
						bingoMap[bingoPatterns[1][j][m]] = true
						bingoMap[bingoPatterns[2][k][m]] = true
						bingoMap[bingoPatterns[3][l][m]] = true
					}
					if (len(bingoMap) < min) {
						min = len(bingoMap)
					}
				}
			}
		}
	}
	fmt.Println(min)
}

// create bingo patterns
func createBingoPatterns(users [4][5][5]int) [4][12][5]int {
	var bingoPatterns [4][12][5]int
	pattern := 0

	for u := 0; u < 4; u++ {
		for y := 0; y < 5; y++ {
			// X
			for x := 0; x < 5; x++ {
				bingoPatterns[u][pattern][x] = users[u][y][x]
			}
			pattern++
			// Y
			for i := 0; i < 5; i++ {
				bingoPatterns[u][pattern][i] = users[u][i][y]
			}
			pattern++
		}
		// Slant
		for i := 0; i < 5; i++ {
			bingoPatterns[u][pattern][i] = users[u][i][i]
			bingoPatterns[u][pattern + 1][i] = users[u][i][4 - i]
		}
		pattern = 0
	}
	return bingoPatterns
}
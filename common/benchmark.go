package common

import (
	"fmt"
	"time"
	"math"
)

func ExecuteIteration(iterateCount int, function func()) {
	var results []float64
	for i := 0; i < iterateCount; i++ {
		fmt.Printf("------------Execution[%d]------------\n", i + 1)
		result := benchmark(function)
		results = append(results, result)
	}

	max, min, average := calculateResults(results)
	fmt.Printf("min:%fms\tmax:%fms\taverage:%fms\n", min, max, average)
}

func benchmark(function func()) float64 {
	start := time.Now()
	function()
	end := time.Now()
	return float64(end.Sub(start).Nanoseconds()) / float64(1000000)
}

func calculateResults(results []float64) (float64, float64, float64) {
	var (
		min float64 = math.MaxFloat64
		max float64 = math.SmallestNonzeroFloat64
		sum float64
	)
	for _, result := range results {
		if result > max {
			max = result
		}
		if result < min {
			min = result
		}
		sum += result
	}

	return max, min, sum / float64(len(results))
}
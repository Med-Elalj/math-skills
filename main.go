package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	a, err := os.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}
	data := strings.Split(string(a), "\n")
	numbers := make([]int, 0, len(data))
	for _, num := range data {
		if num == "" {
			continue
		}
		numFloat, err := strconv.Atoi(num)
		if err != nil {
			continue
		}
		numbers = append(numbers, numFloat)
	}
	fmt.Printf("Average: %.0f\n", math.Round(average(numbers)))
	fmt.Printf("Median: %.0f\n", math.Round(median(numbers)))
	fmt.Printf("Variance: %.0f\n", math.Round(variance(numbers)))
	fmt.Printf("Standard Deviation: %.0f\n", math.Round(standardDeviation(numbers)))
}

func average(numbers []int) float64 {
	sum := 0
	for _, num := range numbers {
		sum += num

	}
	return float64(sum) / float64(len(numbers))
}

func median(numbers []int) float64 {
	n := len(numbers)
	// sortedNumbers := make([]int, n)
	sort.Ints(numbers)
	if n%2 == 1 {
		return float64(numbers[n/2])
	}
	return float64(numbers[n/2-1]+numbers[n/2]) / 2.0
}

func variance(numbers []int) float64 {
	average := average(numbers)
	sumOfSquares := 0.0
	for _, num := range numbers {
		sumOfSquares += (float64(num) - average) * (float64(num) - average)
	}
	return sumOfSquares / float64(len(numbers))
}

func standardDeviation(numbers []int) float64 {
	return math.Sqrt(variance(numbers))
}

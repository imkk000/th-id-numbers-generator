package main

import (
	"flag"
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	var batchSize int
	flag.IntVar(&batchSize, "batch", 1, "input batch size")
	flag.Parse()

	var seed int64
	var source rand.Source
	var r *rand.Rand
	var m runtime.MemStats
	var rLimit = make([][]int, 14)
	for i := 0; i < len(rLimit); i++ {
		rLimit[i] = []int{0, 10}
	}
	// custom random limit
	rLimit[13] = []int{1, 9}
	var digits = make([]int, 13)
	for batchSize != 0 {
		batchSize--
		for i := len(digits); i >= 2; i-- {
			runtime.ReadMemStats(&m)
			seed = time.Now().UnixNano() + int64(m.Alloc) + int64(m.Frees)
			source = rand.NewSource(seed)
			r = rand.New(source)
			digits[i-1] = r.Intn(rLimit[i][1]-rLimit[i][0]) + rLimit[i][0]
			fmt.Print(digits[i-1])
		}
		fmt.Println(CalculateChecksum(digits))
	}
}

func CalculateChecksum(input []int) int {
	var sum int
	for i := len(input); i >= 2; i-- {
		sum += i * input[i-1]
	}
	sum %= 11
	if sum > 1 {
		return 11 - sum
	}
	return 1 - sum
}

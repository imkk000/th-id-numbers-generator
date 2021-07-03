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
	flag.IntVar(&batchSize, "batch", 10, "input batch size")

	var sum, n int
	var seed int64
	var source rand.Source
	var r *rand.Rand
	var m runtime.MemStats
	for batchSize != 0 {
		batchSize--
		for i := 13; i >= 2; i-- {
			runtime.ReadMemStats(&m)
			seed = time.Now().UnixNano() + int64(m.Alloc) + int64(m.Frees)
			source = rand.NewSource(seed)
			r = rand.New(source)
			n = r.Intn(10)
			sum += i * n
			fmt.Print(n)
		}
		sum %= 11
		if sum > 1 {
			sum = 11 - sum
		} else {
			sum = 1 - sum
		}
		fmt.Println(sum)
	}
}

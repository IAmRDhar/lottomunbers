package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberToGenerate = 1000000
)

func main() {
	start := time.Now()
	fmt.Println("Start")

	list := lottoNumbers(numberToGenerate)

	d := time.Now().Sub(start)
	fmt.Println("End ", len(list), " ", d.Seconds())
}

func lottoNumbers(n int) [][]int {
	var parallelCount = 10
	var iterations = n / parallelCount
	var list = make([][]int, parallelCount*iterations)
	var wg sync.WaitGroup

	wg.Add(parallelCount)
	for i := 0; i < parallelCount; i++ {
		go func() {
			defer wg.Done()

			r := rand.New(rand.NewSource(time.Now().UnixNano()))

			for i := 0; i < iterations; i++ {
				list[parallelCount*iterations-1] = add(r)
			}
		}()
	}
	wg.Wait()

	return list
}

func add(r *rand.Rand) []int {
	return []int{
		r.Intn(49),
		r.Intn(49),
		r.Intn(49),
		r.Intn(49),
		r.Intn(49),
		r.Intn(49),
		r.Intn(49),
	}
}

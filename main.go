//This version is much faster. The speedup was achieved by not doing the extra work the original version did.
package main

import (
	"fmt"
	"math/rand"
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
	list := make([][]int, 0, n)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < n; i++ {
		list = append(list, []int{
			rand.Intn(49),
			rand.Intn(49),
			rand.Intn(49),
			rand.Intn(49),
			rand.Intn(49),
			rand.Intn(49),
			rand.Intn(49),
		})
	}

	return list
}

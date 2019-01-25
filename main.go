/**
 * Another thing to notice is the 1.00M
 * allocations per operation. Go allocates
 * memory on the stack or the heap. Stack
 * memory is managed at compile time. Heap
 * memory, on the other hand, is managed at
 * runtime. This includes deciding where on
 * the heap to allocate the memory and
 * garbage collection when that memory is no longer needed
 */

/**
 * The next thing to deal with is the WaitGroup.
 * It needs to be on the heap because it is accessed
 * from different goroutines. What doesn't need to
 * be on the heap are the many references to it.
 * This can be fixed by initially getting a pointer to it.
 *
 * The function and &rand.Rand literal escapes can be
 * removed by moving the creation of r into the closer
 * and then converting the closure into a separate function.
 */
package main

import (
	"fmt"
	"math/rand"
	"runtime"
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
	total := 7 * n
	all := make([]int, total)
	list := make([][]int, n)
	seed := time.Now().UnixNano()
	wg := &sync.WaitGroup{}
	workers := runtime.GOMAXPROCS(-1) // one for each proc

	for i := 0; i < workers; i++ {
		work := total / workers
		begin := work * i
		end := begin + work

		if i == workers-1 {
			end += n % workers
		}

		wg.Add(1)
		//Removing duplicate references to wg by passing
		//a pointer to the wait group reference created
		//initially
		go lottoNumbersWorker(wg, seed+int64(i), all, begin, end)
	}

	wg.Wait()

	for i := range list {
		list[i] = all[i : i+7]
	}

	return list
}

func lottoNumbersWorker(wg *sync.WaitGroup, seed int64, all []int, begin, end int) {
	defer wg.Done()

	r := rand.New(rand.NewSource(seed))

	for i := begin; i < end; i++ {
		all[i] = r.Intn(49)
	}
}

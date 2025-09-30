package main

import (
	"fmt"
	"sync"
	"time"
)

func hello(s string, id int) {
	fmt.Println(s, " - ", id)
	time.Sleep(1 * time.Second)

}
func main() {
	hello("whole begins", -1)

	var wg_whole sync.WaitGroup
	wg_whole.Go(func() {
		for i := 0; i < 6; i++ {
			hello("hello", i)
		}
	})
	wg_whole.Wait()

	hello("whole ends", -1)

	hello("add begins", -1)

	var wg_add sync.WaitGroup
	for i := range 6 {
		wg_add.Add(1)
		localI := i
		go func(i int) {
			defer wg_add.Done()
			fmt.Println("iteration: ", i)
			fmt.Println("iteration copy: ", localI)
			hello("hello", i)
		}(i)
	}
	wg_add.Wait()

	hello("add ends", -1)
}

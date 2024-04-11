package main

import (
	"fmt"
	"time"
)

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

func doUpdateWrong(c Counter) {
	c.Increment()
	fmt.Println("[-] doUpdateWrong: ", c.String())
}

func doUpdateRight(c *Counter) {
	c.Increment()
	fmt.Println("[+] doUpdateRight: ", c.String())
}

func main() {

	c := Counter{
		total:       1,
		lastUpdated: time.Now(),
	}

	fmt.Println(c.String())
	// Take the address of c and create a pointer that points to c.
	// `(&c).Increment()` calls the `Increment` method on that pointer.
	(c).Increment()
	fmt.Println(c.String())

	doUpdateWrong(c) // total will not update -- value.
	fmt.Println(c.String())

	doUpdateRight(&c) // total will update -- pointer.
	fmt.Println(c.String())

}

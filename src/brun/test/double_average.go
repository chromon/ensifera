package main

import (
	"ensifera/infra/algo"
	"fmt"
)

func main() {
	count, amount := int64(10), int64(100 * 100)
	for i := count; i > 0; i-- {
		t := algo.DoubleAverage(i, amount)
		amount -= t
		fmt.Print(float64(t) / float64(100), ",")
	}
	fmt.Print()
}

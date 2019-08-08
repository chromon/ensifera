package main

import (
	"ensifera/infra/algo"
	"fmt"
)

func main() {
	count, amount := int64(10), int64(100)
	for i := int64(0); i < count; i++ {
		t := algo.BeforeShuffle(count, amount * 100)
		fmt.Print(t, ",")
	}
	fmt.Print()
}

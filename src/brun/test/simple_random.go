package main

import (
	"ensifera/infra/algo"
	"fmt"
)

func main() {
	count, amount := int64(10), int64(100)
	for i := int64(0); i < count; i++ {
		res := algo.SimpleRandom(count, amount * 100)
		fmt.Print(float64(res) / float64(100), ", ")
	}
	fmt.Print()
}

package main

import (
	"ensifera/infra/algo"
	"fmt"
)

func main() {
	fmt.Print(algo.AfterShuffle(int64(10), int64(100 * 100)))
}

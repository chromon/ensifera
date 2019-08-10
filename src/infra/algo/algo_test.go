package algo

import (
	."github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSimpleRandom(t *testing.T) {
	ForTest("test simple random", t, SimpleRandom)
}

func TestBeforeShuffle(t *testing.T) {
	ForTest("test before shuffle", t, BeforeShuffle)
}

func TestDoubleRandom(t *testing.T) {
	ForTest("double random", t, DoubleRandom)
}

func TestDoubleAverage(t *testing.T) {
	ForTest("test double average.", t, DoubleAverage)
}

func ForTest(message string, t *testing.T, fn func(count, amount int64) int64) {
	count, amount := int64(10), int64(100 * 100)
	remain := amount
	sum := int64(0)
	for i := count; i > 0; i-- {
		t := fn(i, remain)
		remain -= t
		sum += t
	}

	Convey(message, t, func() {
		So(sum, ShouldEqual, amount)
	})
}

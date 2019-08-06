package algo

import (
	"math/rand"
	"time"
)

const min = int64(1)

// 简单随机算法
// count 红包数量
// amount 红包金额
func SimpleRandom(count, amount int64) int64 {

	// 当红包数量剩余一个的时候，直接返回数量
	if count == 1 {
		return amount
	}

	// 计算最大可调用金额，随机数 + 最小金额作为红包序列元素
	// 缺点：大概率先大后小
	max := amount - min * count
	rand.Seed(time.Now().UnixNano())
	res := rand.Int63n(max) + min

	return res
}


package algo

import (
	"math/rand"
	"time"
)

/*
	简单随机算法优化：关键在于控制每次抢红包时，随机数的上限
	二倍均值算法：
		控制每次抢红包时随机数的最大数，从而限制随机范围实现的

	假设红包总金额是 M ，共有 N 个红包，平均值 A = M / N
	那么每次随机的红包金额中位数就是 A ，随机数均匀分布在 A 的两边
	并且两边跨度不超过平均值 A ，也就是最大不超过 A 的两倍
	随机范围也就是 2 倍平均值 A 之间
 */

func DoubleAverage(count, amount int64) int64 {
	if count == 1 {
		return amount
	}
	// 计算出最大可用金额
	max := amount - min * count
	// 计算最大可用平均值（max < count 时，avg = 0）
	avg := max / count
	// 计算二倍均值，并在其基础上加上最小金额，防止出现 0 值
	avg2 := 2 * avg + min
	// 随机红包金额序列元素，把二倍均值作为随机的最大数
	rand.Seed(time.Now().UnixNano())
	a := rand.Int63n(avg2) + min

	return a
}
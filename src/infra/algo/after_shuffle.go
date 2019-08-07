package algo

import "math/rand"

/*
	后洗牌算法：2次随机，先随机再洗牌
		随机生成序列，范围为0 ~ 最大可用金额
		随机数 + 最小金额作为红包序列元素
		生成的序列在洗牌打乱
	缺点：
		需要事先生成整个序列
		只适用于发红包时生成
	性能：
		与简单随机相差不多，略慢万分之一
 */
func AfterShuffle(count, amount int64) []int64 {

	// 红包序列
	res := make([]int64, 0)
	// 计算最大金额
	max := amount - min * count
	// 剩余最大金额
	remain := max
	// 随机生成初始红包序列
	for i := int64(0); i < count; i++ {
		a := SimpleRandom(count - i, remain)
		remain -= a
		res = append(res, a)
	}
	// 洗牌
	rand.Shuffle(len(res), func(i, j int) {
		res[i], res[j] = res[j], res[i]
	})

	return res
}

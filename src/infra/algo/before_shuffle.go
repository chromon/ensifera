package algo

import (
    "math/rand"
    "time"
)

/*
	先洗牌算法：2次随机，先洗牌再随机，先计算一个种子金额，将种子金额
		作为随机最大金额，然后将随机出来的金额加上最小金额作为红包序列元素。
	1. 生成一个一定长度的红包种子金额序列
		种子金额序列长度控制在 3 ~ 1/3 剩余红包数量
		序列元素为最大可用剩余金额/(n + 1)，n为红包种子金额序列的索引数
		最大可用剩余金额 = 剩余金额 - 最小金额 * 剩余数量
	2. 随机在种子红包序列中选择一个元素作为随机数，即再随机一个数字，取模计算索引
	3. 按索引从种子金额序列拿出一个金额作为基数
	4. 基数作为最大数随机 + 最小金额作为红包序列元素
 */

func BeforeShuffle(count, amount int64) int64 {

    if count == 1 {
        return amount
    }

    // 计算出最大可调用金额
    max := amount - min * count
    // 生成红包种子金额序列
    seeds := make([]int64, 0)
    // 红包种子金额序列长度：3 ~ 1/2 * count
    size := count / 2
    if size < 3 {
        size = 3
    }
    for i := int64(0); i < size; i++ {
        a := max / (i + 1)
        seeds = append(seeds, a)
    }
    rand.Seed(time.Now().UnixNano())
    // 从红包种子金额序列中随机选择一个作为随机基数
    index := rand.Int63n(int64(len(seeds)))
    // 使用随机基数作为最大数，随机出一个数字作为红包金额序列元素
    res := rand.Int63n(seeds[index])

    return res + min
}
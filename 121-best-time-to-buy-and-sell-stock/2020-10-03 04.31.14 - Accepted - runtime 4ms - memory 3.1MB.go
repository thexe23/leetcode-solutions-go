func maxProfit(prices []int) int {
    min := 1<<31 - 1
    res := 0
    for i := 0; i < len(prices); i++ {
        if prices[i] < min {
            min = prices[i]
        }else if p := prices[i] - min; p > res {
            res = p
        }
    }
    return res
}
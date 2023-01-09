func maxProfit(prices []int) int {
    sold, hold, rest := 0, -1<< 31, 0
    for _, v := range prices {
        temp := sold
        sold = max(rest, sold)
        rest = hold + v
        hold = max(temp - v, hold)
    }
    return max(sold, rest)
}
func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

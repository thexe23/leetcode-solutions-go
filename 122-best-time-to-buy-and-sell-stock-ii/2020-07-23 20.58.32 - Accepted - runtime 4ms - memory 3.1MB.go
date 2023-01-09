func maxProfit(prices []int) int {
    maxProfit := 0;
    if len(prices) < 2 {
        return maxProfit
    }
    for i := 1; i < len(prices); i++{
        if prices[i] > prices[i-1]{
            maxProfit +=  prices[i]-prices[i-1]
        }
    }
    return maxProfit
}
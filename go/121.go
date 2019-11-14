func maxProfit(prices []int) int {
    res, buy := 0, math.MaxInt32
    for _, price := range prices {
        buy = min(buy, price)
        res = max(res, price - buy)
    }
    return res
}

func min(a, b int) int {
    if a <= b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}

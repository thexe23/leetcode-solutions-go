func maxDistToClosest(seats []int) int {
    res, n, last := 0, len(seats), -1
    for i := 0; i < len(seats); i++ {
        if seats[i] == 1 {
            if last == -1 {
                res = i
            }else {
                res = max(res, (i - last) / 2)
            }
            last = i
        }
    }
    res = max(res, n - last - 1)
    return res
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
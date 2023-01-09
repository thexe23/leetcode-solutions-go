func maxDistToClosest(seats []int) int {
    count := 0
    res := 0
    index := len(seats)
    for i := 0; i < len(seats); i++ {
        if seats[i] == 0 {
            index = min(index, i)
            count++
            if index == 0  {
                count++
            }
            if i == len(seats) - 1 {
                count *= 2
            }
            res = max(res, count)
        }else {
            count = 0
            index = len(seats)
        }
    }
    return (res + 1) / 2
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
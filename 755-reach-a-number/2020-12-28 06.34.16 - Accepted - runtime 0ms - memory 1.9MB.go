func reachNumber(target int) int {
    target = abs(target)
    k := 0
    for target > 0 {
        k++
        target -= k
    }
    if target % 2 == 0 {
        return k
    }
    return k + 1 + k % 2
}

func abs(x int) int {
    if x > 0 {
        return x
    }
    return -x
}
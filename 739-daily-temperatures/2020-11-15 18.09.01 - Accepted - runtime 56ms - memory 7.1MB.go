func dailyTemperatures(T []int) []int {
    stack := []int{}
    res := make([]int, len(T))
    
    for i := 0; i < len(T); i++ {
        for len(stack) != 0 && T[i] > T[stack[len(stack) - 1]] {
            index := stack[len(stack) - 1]
            stack = stack[:len(stack) - 1]
            res[index] = i - index
        }
        stack = append(stack, i)
    }
    
    return res
}
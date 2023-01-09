func leastInterval(tasks []byte, n int) int {
    taskMap := make([]int, 26)
    max, maxCount := 0, 0
    for _, v := range tasks{
        taskMap[v - 'A']++
        if taskMap[v - 'A'] == max {
            maxCount++
        }else if taskMap[v - 'A'] > max {
            maxCount = 1
            max = taskMap[v - 'A']
        }
    }
    partCount := max - 1
    partLen := n - (maxCount - 1)
    empty := partCount * partLen
    available := len(tasks) - maxCount * max
    idle := empty - available
    if idle < 0 {
        idle = 0
    }
    return len(tasks) + idle
}
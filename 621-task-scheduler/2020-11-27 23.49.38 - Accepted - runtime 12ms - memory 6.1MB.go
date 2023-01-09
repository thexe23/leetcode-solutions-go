func leastInterval(tasks []byte, n int) int {
    counter := make([]int, 26)
    max, maxCount := 0, 0 
    for _, v := range tasks {
        counter[v - 'A']++
        if max == counter[v - 'A'] {
            maxCount++
        }else if max < counter[v - 'A'] {
            max = counter[v - 'A']
            maxCount = 1
        }
    }
    partCount := max - 1
    partLen := n - (maxCount - 1)
    empty := partCount * partLen
    available := len(tasks) - maxCount * max
    idles := empty - available
    if idles < 0 {
        idles = 0
    }
    return len(tasks) + idles
}
func topKFrequent(nums []int, k int) []int {
    bucket := make([][]int, len(nums) + 1)
    res := make([]int, k)
    freq := make(map[int]int)
    
    for _, n := range nums {
        freq[n]++
    }
    
    for num, time := range freq {
        bucket[time] = append(bucket[time], num)
    }
    var j = 0
    for i := len(bucket) - 1; j < k; i-- {
        for _, v := range bucket[i] {
            res[j] = v
            j++
        }
    }
    return res
}
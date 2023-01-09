func lengthOfLIS(nums []int) int {
    N := len(nums)
    tail := make([]int, N)
    size := 0
    for _, v := range nums{
        i := 0
        j := size
        for i != j {
            m := i + (j - i)/2
            if tail[m] < v {
                i = m + 1
            }else {
                j = m
            }
        }
        tail[i] = v
        if i == size {
            size++
        }
    }
    return size
}
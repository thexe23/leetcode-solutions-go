func permute(nums []int) [][]int {
    n := len(nums)
    res := [][]int{}
    cur := []int{}
    visited := make([]bool, n, n)
    backtrack(&res, nums, cur, n, visited)
    return res
}

func backtrack(res *[][]int, nums, cur []int, n int, visited []bool) {
    if len(cur) == n {
        temp := make([]int, n)
        copy(temp, cur)
        *res = append(*res, temp)
        return
    }
    for i := 0; i < n; i++ {
        if visited[i] == true {
            continue
        }
        cur = append(cur, nums[i])
        visited[i] = true
        backtrack(res, nums, cur, n, visited)
        cur = cur[:len(cur) - 1]
        visited[i] = false
    }
}
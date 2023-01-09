func permuteUnique(nums []int) [][]int {
    sort.Ints(nums)
    res := [][]int{}
    row := []int{}
    visited := make([]bool, len(nums))
    backtrack(nums, visited, 0, &res, row)
    return res
}

func backtrack(nums []int, visited []bool, count int, res *[][]int, row []int) {
    if count == len(nums) {
        temp := make([]int, len(nums))
        copy(temp, row)
        *res = append(*res, temp)
        return
    }
    
    for i := 0; i < len(nums); i++ {
        if i > 0 && nums[i] == nums[i - 1] && !visited[i - 1]{
            continue
        }
        if !visited[i] {
            row = append(row, nums[i])
            visited[i] = true
            backtrack(nums, visited, count + 1, res, row)
            row = row[:len(row) - 1]
            visited[i] = false
        }
    }
}
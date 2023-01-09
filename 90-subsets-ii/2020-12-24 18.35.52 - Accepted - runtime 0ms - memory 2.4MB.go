func subsetsWithDup(nums []int) [][]int {
    res := [][]int{}
    sort.Ints(nums)
    backtrack(nums, &res, []int{}, 0)
    return res
}

func backtrack(nums []int, res *[][]int, row []int, index int) {
    temp := make([]int, len(row))
    copy(temp, row)
    *res = append(*res, temp)
    
    for i := index; i < len(nums); i++ {
        if i != index && nums[i] == nums[i - 1] {
            continue
        }
        row = append(row, nums[i])
        backtrack(nums, res, row, i + 1)
        row = row[:len(row) - 1]
    }
}
func subsets(nums []int) [][]int {
    if len(nums) == 0 {
        return [][]int{}
    }
    
    res := [][]int{{}}
    findAllSubsets(nums, &res, 0, []int{})
    return res
}

func findAllSubsets(nums []int, res *[][]int, begin int, row []int) {
    for i := begin; i < len(nums); i++ {
        row = append(row, nums[i])
        temp := make([]int, len(row))
        copy(temp, row)
        *res = append(*res, temp)
        findAllSubsets(nums, res, i + 1, row)
        row = row[:len(row) - 1]
    }
}
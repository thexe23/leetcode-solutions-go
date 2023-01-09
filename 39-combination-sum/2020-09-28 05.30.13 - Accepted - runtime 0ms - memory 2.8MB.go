func combinationSum(candidates []int, target int) [][]int {
    row, res := []int{}, [][]int{}
    findAll(target, 0, row, candidates, &res)
    return res
}

func findAll(target, index int, row []int, candidates []int, res *[][]int) {
    if target == 0 {
        temp := make([]int, len(row))
        copy(temp, row)
        *res = append(*res, temp)
        return
    }
    
    for i := index; i < len(candidates); i++ {
        if candidates[i] > target {
            continue
        }
        row = append(row, candidates[i])
        findAll(target - candidates[i], i, row, candidates, res)
        row = row[:len(row) - 1]
    }
}
func combinationSum2(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    row, res := []int{}, [][]int{}
    findAll(target, 0, candidates, row, &res)
    return res
}

func findAll(target, index int, candidates, row []int, res *[][]int) {
    if target == 0 {
            temp := make([]int, len(row))
            copy(temp, row)
            *res = append(*res, temp)
    }
    
    for i := index; i < len(candidates); i++ {
        if candidates[i] > target {
            break
        }
        
        if i > index && candidates[i] == candidates[i - 1] {
            continue
        }
        row = append(row, candidates[i])
        findAll(target - candidates[i], i + 1, candidates, row, res)
        row = row[:len(row) - 1]
    }
}
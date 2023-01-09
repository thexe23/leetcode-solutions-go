func minIncrementForUnique(A []int) int {
    if len(A) < 2 {
        return 0
    }
    sort.Ints(A)
    min := A[0]
    res := 0
    for i := 1; i < len(A); i++ {
        if A[i] > min {
            min = A[i]
        }else{
            min++
            res += min - A[i]
            A[i] = min
        }
    }
    return res
}
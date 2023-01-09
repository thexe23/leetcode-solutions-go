func convert(s string, numRows int) string {
    next := 2*numRows - 2
    l := len(s)
    res := make([]byte, len(s))
    if numRows == 1 {
        return s
    }
    k := 0
    for i := 0; i < numRows; i++ {
        for j := i; j < l; j = j + next {
            res[k] = s[j]
            k++
            if i != 0 && i != numRows - 1 && j + next - 2 * i < l{
                res[k] = s[j + next - 2 * i]
                k++
            }
        }
    }
    return string(res)
}
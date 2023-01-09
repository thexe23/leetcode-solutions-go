func convert(s string, numRows int) string {
    next := 2*numRows - 2
    l := len(s)
    res := []byte{}
    if numRows == 1 {
        return s
    }
    for i := 0; i < numRows; i++ {
        for j := i; j < l; j = j + next {
            res = append(res, s[j])
            if i != 0 && i != numRows - 1 && j + next - 2 * i < l{
                res = append(res, s[j + next - 2 * i])
            }
        }
    }
    return string(res)
}
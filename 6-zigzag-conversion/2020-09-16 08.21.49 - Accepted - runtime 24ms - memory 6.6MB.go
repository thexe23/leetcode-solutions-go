func convert(s string, numRows int) string {
    if numRows == 0 || len(s) == 0 || numRows == 1{
        return s
    }
    zig := make([][]byte, numRows)
    res := ""
    i := 0
    k := 0
    for k < len(s) {
        for i < numRows && k < len(s) {
            zig[i] = append(zig[i], s[k])
            i++
            k++
        }
        i--
        for i > 0 && k < len(s) {
            i--
            zig[i] = append(zig[i], s[k])
            k++
        }
        i++
    }
    for i = 0; i < len(zig); i++ {
        for j := 0; j < len(zig[i]); j++ {
                res += string(zig[i][j])
        }
    }
    return res
}
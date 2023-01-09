func convert(s string, numRows int) string {
    if len(s) == 0 || numRows <= 1{
        return s
    }
    res := make([]byte, len(s))
    colStep := numRows + numRows - 2
    diagStep := colStep - 2
    i := 0
    for row := 0; row < numRows; row++ {
        diag := row > 0 && row < numRows - 1
        for j := row; j < len(s); j = j + colStep{
            res[i] += s[j]
            i++
            if diag && j + diagStep < len(s) {
                res[i] += s[j+diagStep]
                i++
            }
        }
        if diag {
            diagStep -= 2
        }
    }
    return string(res)
}
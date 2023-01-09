func convert(s string, numRows int) string {
    if len(s) == 0 || numRows <= 1{
        return s
    }
    var res string = ""
    colStep := numRows + numRows - 2
    diagStep := colStep - 2
    for row := 0; row < numRows; row++ {
        diag := row > 0 && row < numRows - 1
        for j := row; j < len(s); j = j + colStep{
            res += string(s[j])
            if diag && j + diagStep < len(s) {
                res += string(s[j+diagStep])
            }
        }
        if diag {
            diagStep -= 2
        }
    }
    return res
}
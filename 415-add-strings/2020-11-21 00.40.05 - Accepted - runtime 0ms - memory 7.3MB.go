func addStrings(num1 string, num2 string) string {
    res := ""
    re := 0
    for i, j := len(num1) - 1, len(num2) - 1; i >= 0 || j >= 0 || re > 0; i, j = i - 1, j - 1 {
        num := 0
        if i >= 0 {
            num += int(num1[i] - '0')
        }
        if j >= 0 {
            num += int(num2[j] - '0')
        }
        num += re
        res = strconv.Itoa(num % 10) + res
        re = num / 10
    }
    return res
}
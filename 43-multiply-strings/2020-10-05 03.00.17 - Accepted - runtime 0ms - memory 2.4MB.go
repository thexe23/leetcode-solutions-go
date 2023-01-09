func multiply(num1 string, num2 string) string {
    n := len(num1)
    m := len(num2)
    res := make([]int, m + n)
    for i := m - 1; i >= 0; i-- {
        for j := n - 1; j >= 0; j-- {
            mul := int(num1[j] - '0') * int(num2[i] - '0')
            p1, p2 := i + j, i + j + 1
            sum := mul + res[p2]
            res[p2] = sum % 10
            res[p1] += sum / 10
        }
    }
    
    i := 0
    for i < len(res) && res[i] == 0 {
        i++
    }
    
    resStr := []byte{}
    for i < len(res) {
        resStr = append(resStr, byte(res[i]) + '0')
        i++
    }
    if len(resStr) == 0{
        return "0"
    }
    return string(resStr)
}
func countAndSay(n int) string {
    x := "1"
    for j := 1; j < n; j++ {
        res := []byte{}
        num := 1
        i := 1
        for ;i < len(x); i++ {
            if x[i] == x[i - 1] {
                num++
            }else {
                res = append(res, byte(num + '0'), x[i - 1])
                num = 1
            }
        }
        res = append(res, byte(num + '0'), x[i - 1])
        x = string(res)
    }
    return x
}
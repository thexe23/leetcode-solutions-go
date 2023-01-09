func countAndSay(n int) string {
    if n == 1 {
        return "1"
    }
    
    x := countAndSay(n - 1)
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
    return string(res)
}
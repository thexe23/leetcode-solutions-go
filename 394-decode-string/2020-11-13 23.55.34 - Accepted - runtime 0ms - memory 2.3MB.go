func decodeString(s string) string {
    numStack := []int{}
    strStack := []string{}
    str, num := "", 0
    for i := 0; i < len(s); i++{
        if '0' <= s[i] && s[i] <= '9'{
            digit := int(s[i] - '0')
            num = num * 10 + digit   
        }else if s[i] == '[' {
            strStack = append(strStack, str)
            numStack = append(numStack, num)
            str = ""
            num = 0
        }else if s[i] == ']' {
            num = numStack[len(numStack) - 1]
            numStack = numStack[:len(numStack) - 1]
            prevStr := strStack[len(strStack) - 1]
            strStack = strStack[:len(strStack) - 1]
            for num > 0{
                prevStr += str
                num--
            }
            str = prevStr
        }else {
             str += string(s[i])
        }
    }
    return str
}
func evalRPN(tokens []string) int {
    stack := []int{}
    for _, v := range tokens {
        if v == "+" || v == "-" || v == "*" || v == "/" {
            a := stack[len(stack) - 2]
            b := stack[len(stack) - 1]
            stack = stack[:len(stack) - 2]
            if v == "+" {
                stack = append(stack, a + b)
            }else if v == "-" {
                stack = append(stack, a - b)
            }else if v == "*" {
                stack = append(stack, a * b)
            }else if v == "/" {
                stack = append(stack, a / b)
            }
        }else {
            num, _ := strconv.Atoi(v)
            stack = append(stack, num)
        }
    }
    return stack[0]
}
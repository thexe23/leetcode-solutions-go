func myAtoi(str string) int {
    const INT_MAX = 2147483647
    const INT_MIN = -2147483648 
    i := 0
    var sign float64 = 1
    var result float64 = 0
    
    if (len(str) == 0){
        return 0
    }
    
    for (i < len(str) && str[i] == ' ') {
        i++
    }
    
    if i < len(str) && (str[i] == '+' || str[i] == '-') {
        if str[i] == '+' {
            sign = 1
        }else{
            sign = -1
        }
        i++
    }
    
    for (i < len(str) && str[i] >= '0' && str[i] <= '9') {
        result = result * 10 + float64(str[i] - '0')
        i++
    }
    result = result * sign
    if result > INT_MAX {
        result = INT_MAX
    }
    if result < INT_MIN {
        result = INT_MIN
    }
    
    int_res := int(result)
    return int_res
}
func reverse(x int) int {
    const INT_MIN = -2147483648
    const INT_MAX = 2147483647
    
    reverse := 0
    for ;x != 0; x = x/10 {
        reverse += x%10
        reverse = reverse*10
    }
        reverse = reverse/10
    
     if reverse < INT_MAX && reverse > INT_MIN{
        return reverse
     }else{
        return 0
     }
}
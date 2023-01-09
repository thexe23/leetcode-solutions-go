func addBinary(a string, b string) string {
    res := []byte{}
    re := 0
    i:= len(a) -1
    j := len(b) - 1
    
    for i >= 0 || j >= 0 || re == 1 {
        sum := 0
        if i >= 0 {
            sum += int(a[i] - '0')
            i--
        }
        
        if j >= 0 {
            sum += int(b[j] - '0')
            j--
        }
        
        sum += re
        res = append([]byte{byte(sum % 2) + '0'}, res...)
        re = sum / 2
    }
    return string(res)
}
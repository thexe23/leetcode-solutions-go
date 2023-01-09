func reverseBits(num uint32) uint32 {
    pow := 31
    var res uint32
    for pow >= 0 {
        res += (num & 1) << pow
        num >>= 1
        pow--
    }
    return res
}
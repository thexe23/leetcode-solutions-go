func singleNumber(nums []int) []int {
    mask := 0
    for _, v := range nums {
        mask ^= v
    }
    mask &= -mask
    a, b := 0, 0
    //divide into two part
    for _, v := range nums {
        if v & mask == 0 {
            a ^= v
        }else {
            b ^= v
        }
    }
    return []int{a, b}
}
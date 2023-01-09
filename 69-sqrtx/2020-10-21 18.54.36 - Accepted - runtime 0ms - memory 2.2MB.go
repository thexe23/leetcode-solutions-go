func mySqrt(x int) int {
    start := 1
    end := 1<< 16
    
    for start <= end {
        mid := start + (end - start)/2
        if x / mid < mid {
            end = mid - 1
        }else {
            if x / (mid + 1) < mid + 1 {
                return mid
            }
            start = mid + 1
        }
    }
    return 0
}
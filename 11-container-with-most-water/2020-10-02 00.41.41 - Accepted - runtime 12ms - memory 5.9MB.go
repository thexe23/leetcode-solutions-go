func maxArea(height []int) int {
    res := 0
    left := 0
    right := len(height) - 1
    for left < right {
        h := min(height[right], height[left])
        res = max(res, (right - left) * h )
        if height[left] < height[right] {
            left++
        }else{
            right--
        }
    }
    return res
}

func max(x, y int) int {
    if x >y {
        return x
    }
    return y
}

func min(x, y int) int {
    if x >y {
        return y
    }
    return x
}
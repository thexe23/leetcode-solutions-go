func trap(height []int) int {
    // Two Pointer
    if len(height) == 0 {
        return 0
    }
    res := 0
    leftMax := height[0]
    rightMax := height[len(height) - 1]
    left := 0
    right := len(height) - 1
    for left < right {
        if height[left] <= height[right] {
            if height[left] >= leftMax {
                leftMax = height[left]
            }else{
                res += (leftMax - height[left])
            }
            left++
        }else{
            if height[right] >= rightMax {
                rightMax = height[right]
            }else{
                res += (rightMax - height[right])
            }
            right--
        }
    }
    return res
}
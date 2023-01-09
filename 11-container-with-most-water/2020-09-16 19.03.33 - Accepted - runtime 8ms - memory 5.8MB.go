func maxArea(height []int) int {
    res := 0
    i := 0
    j := len(height) - 1
    for i < j {
        h := min(height[i], height[j])
        w := j - i
        res = max(res, h * w)
        if height[i] < height[j] {
            i++
        }else{
            j--
        }
    }
    return res
}

func min(x,y int) int {
    if x < y {
        return x
    }else{
        return y
    }
}

func max(x, y int) int {
    if x > y {
        return x
    }else{
        return y
    }
}
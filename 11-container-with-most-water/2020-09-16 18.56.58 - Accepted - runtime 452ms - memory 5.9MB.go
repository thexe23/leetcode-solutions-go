func maxArea(height []int) int {
    res := 0
    for i := 0; i < len(height); i++ {
        for j := i; j < len(height); j++ {
            h := min(height[i], height[j])
            w := j - i
            res = max(res, h * w)
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
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    m, n := len(nums1), len(nums2)
    if m > n {
        return findMedianSortedArrays(nums2, nums1)
    }
    
    half := (m + n + 1) / 2
    lo, hi := 0, m
    for lo <= hi {
        i := (lo + hi) / 2
        j := half - i
        if i > 0 && nums1[i - 1] > nums2[j] {
            hi = i - 1
        }else if i < m && nums1[i] < nums2[j - 1] {
            lo = i + 1
        }else {
            leftMax := 0
            if i == 0 {
                leftMax = nums2[j - 1]
            }else if j == 0 {
                leftMax = nums1[i - 1]
            }else {
                leftMax = max(nums1[i - 1], nums2[j - 1])
            }
            
            if (m + n) % 2 == 1 {
                return float64(leftMax)
            }
            
            rightMin := 1 << 31 - 1
            if i == m {
                rightMin = nums2[j]
            }else if j == n {
                rightMin = nums1[i]
            }else {
                rightMin = min(nums1[i], nums2[j])
            }
            
            return (float64(leftMax) + float64(rightMin)) / 2
        }
    }
    return 0.0
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
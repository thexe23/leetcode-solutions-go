func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    m := len(nums1)
    n := len(nums2)
    
    if m > n {
        return findMedianSortedArrays(nums2, nums1)
    }
    
    imin, imax, half := 0, m, (m + n + 1) / 2
    for imin <= imax {
        i := (imin + imax) / 2
        j := half - i
        if i < m && nums1[i] < nums2[j - 1] {
            imin = i + 1
        }else if i > 0 && nums1[i - 1] > nums2[j] {
            imax = i - 1
        }else {
            leftMax := 0
            if i == 0 {
                leftMax = nums2[j - 1]
            }else if j == 0 {
                leftMax = nums1[i - 1]
            }else {
                leftMax = max(nums1[i - 1], nums2[j - 1])
            }
            if (m + n) % 2 != 0 {
                return float64(leftMax)
            }
            
            rightMin := 0
            if i == m {
                rightMin = nums2[j]
            }else if j == n {
                rightMin = nums1[i]
            }else {
                rightMin = min(nums1[i], nums2[j])
            }
            return float64(rightMin + leftMax) / 2
        }
    }
    return 0.0
}

func min(x, y int) int {
    if x > y {
        return y
    }
    return x
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
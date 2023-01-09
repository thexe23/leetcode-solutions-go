func merge(nums1 []int, m int, nums2 []int, n int)  {
    i, j := m - 1, n - 1
    target := m + n -1
    for target >= 0 {
        if j < 0 || (i >= 0 && nums1[i] >= nums2[j]) {
            nums1[target] = nums1[i]
            i--
        }else {
            nums1[target] = nums2[j]
            j--
        }
        target--
    }
    
}
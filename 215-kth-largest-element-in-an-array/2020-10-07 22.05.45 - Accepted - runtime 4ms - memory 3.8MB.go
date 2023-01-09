import "math/rand"

var N []int

func swap (a, b int) {
    N[a], N[b] = N[b], N[a]
}

func partition(left, right, pivotIndex int) int {
    pivot := N[pivotIndex]
    
    swap(pivotIndex, right)
    storeIndex := left
    
    for i := left; i <= right; i++ {
        if N[i] < pivot {
            swap(storeIndex, i)
            storeIndex++
        }
    }
    
    swap(storeIndex, right)   
    return storeIndex
}

func quickselect(left, right, kSmallest int) int {
    if left == right {
        // list contains one element
        return N[left]
    }
    
    // select a random pivotIndex
    pivotIndex := left + rand.Intn(right - left)
    
    pivotIndex = partition(left, right, pivotIndex)
    
    if kSmallest == pivotIndex {
        return N[kSmallest]
    } else if kSmallest < pivotIndex {
        //go left
        return quickselect(left, pivotIndex - 1, kSmallest)       
    }
    
    return quickselect(pivotIndex + 1, right, kSmallest)
}

func findKthLargest(nums []int, k int) int {
    N = make([]int, len(nums))
    copy(N, nums)
    
    length := len(N)
    return quickselect(0, length - 1, length - k)
}
func numSquares(n int) int {
    for n & 3 == 0 {
        n >>= 2
    }
    if n & 7 == 7 {
        return 4
    }
    
    if isSquare(n) {
        return 1
    }
    sqr := int(math.Sqrt(float64(n)))
    for i := 1; i <= sqr; i++ {
        if isSquare(n - i * i) {
            return 2
        }
    }
    return 3
}

func isSquare(x int) bool {
    temp := int(math.Sqrt(float64(x)))
    return temp * temp == x
}
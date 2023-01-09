func lengthOfLastWord(s string) int {
    start := -1
    end := len(s) - 1
    for end >= 0 && s[end] == ' '{
        end--
    }
    
    for i := 0; i < end; i++ {
        if s[i] == ' '{
            start = i
        }
    }
    
    return end - start
}
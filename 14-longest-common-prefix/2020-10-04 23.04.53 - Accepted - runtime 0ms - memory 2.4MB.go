func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    prefix := strs[0]
    for i := 0; i < len(strs); i++ {
        for strings.HasPrefix(strs[i], prefix) == false{
            prefix = prefix[:len(prefix) - 1]
        }
    }
    return prefix
}
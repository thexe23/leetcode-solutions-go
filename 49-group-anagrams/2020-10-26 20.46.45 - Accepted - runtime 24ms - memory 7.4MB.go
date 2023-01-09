func groupAnagrams(strs []string) [][]string {
    m := make(map[string][]string)
    res := [][]string{}
    for _, str := range strs {
        key := hash(str)
        if m[key] != nil {
            m[key] = append(m[key], str)
        }else {
            m[key] = []string{str}
        }
    }
    
    for _, item := range m {
        res = append(res, item)
    }
    return res
}

func hash(str string) string {
    res := make([]byte, 26)
    for i := 0; i < len(str); i++ {
        res[str[i] - 'a']++
    }
    return string(res)
}
func groupAnagrams(strs []string) [][]string {
    m := make(map[string][]string)
    res := [][]string{}
    for _, str := range strs {
        s := strings.Split(str, "")
        sort.Strings(s)
        key := strings.Join(s, "")
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
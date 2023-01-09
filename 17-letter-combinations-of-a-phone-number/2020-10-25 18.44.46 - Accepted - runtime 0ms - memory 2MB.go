var letters = [][]byte {
    {},
    {},
    {'a', 'b', 'c'},
    {'d', 'e', 'f'},
    {'g', 'h', 'i'},
    {'j', 'k', 'l'},
    {'m', 'n', 'o'},
    {'p', 'q', 'r', 's'},
    {'t', 'u', 'v'},
    {'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
    if digits == ""{
        return nil
    }
    res := []string{}
    combine := []byte{}
    backtrack(digits, &res, combine, 0)
    return res
}

func backtrack(digits string, res *[]string, combine []byte, count int) {
    if count == len(digits) {
        temp := make([]byte, len(digits))
        copy(temp, combine)
        *res = append(*res, string(temp))
        return
    }
    num := digits[count] - '0'
    for i := 0; i < len(letters[num]); i++ {
        combine = append(combine, letters[num][i])
        backtrack(digits, res, combine, count + 1)
        combine = combine[:len(combine) - 1]
    }
}


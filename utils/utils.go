package utils

// TODO: support a generic interable
func SwapFromSlice(s []int, index int) []int {
    elementOne := s[index]

    s[index] = s[index + 1]
    s[index + 1] = elementOne

    return s
}


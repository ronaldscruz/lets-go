package strutils

// "func (r)everse" is wrong because if the first letter of a function is in lowercase,
// it means that it's private, so it will not be exported

// Reverse : reverse string ("abc" => "cba")
func Reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

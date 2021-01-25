package main

import "fmt"

func main() {
	var data string
	fmt.Scanf("%s", &data)
	fmt.Printf("%s", isBalancedBrackets(data))
}
func isBalancedBrackets(data string) string {
	result := "NO"
	brackets := make(map[byte]byte)
	brackets['('] = ')'
	brackets['['] = ']'
	brackets['{'] = '}'
	sli := make([]byte, 0)
	if len(data) == 0 {
		return "YES"
	}
	if len(data)%2 == 1 {
		return "NO"
	}
	for i, b := range data {
		// jika diawali oleh salah satu tanda tutup
		if i == 0 && (data[i] == ')' || data[i] == ']' || data[i] == '}') {
			return "NO"
		}
		if b == '(' || b == '[' || b == '{' {
			sli = append(sli, byte(b)) //tampung jika diawali tanda buka
		} else if b == ')' || b == ']' || b == '}' {
			left := sli[len(sli)-1]
			sli = sli[:len(sli)-1]
			if brackets[left] == byte(b) {
				result = "YES"
			} else {
				return "NO"
			}
		}
		if i == len(data)-1 && len(sli) != 0 {
			result = "NO"
		}
	}
	return result
}

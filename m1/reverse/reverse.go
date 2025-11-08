package reverse

import "strconv"

func String(s string) string {
	runes := []rune(s)
	return string(runes)
}

func Int(i int) int {
	i, _ = strconv.Atoi(String(strconv.Itoa(i)))
	return i
}

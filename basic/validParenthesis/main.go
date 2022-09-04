package main

func validParenthesis(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	st := []rune{}
	open := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	for _, r := range s {
		if closer, ok := open[r]; ok {
			st = append(st, closer)
			continue
		}

		l := len(st) - 1
		if l < 0 || r != st[l] { // 組み合わせが正しければオープンの次はクローズが必ずくる
			return false
		}

		st = st[:l] // ひとつずつスライス
	}

	return len(st) == 0 // 空になれば組み合わせは正しい
}

func main() {
	validParenthesis("[[]]")
}

package string

var validParenthesesMap = map[string]string{
	"(": ")",
	"{": "}",
	"[": "]",
}

// 使用入栈法，入栈的只是右括号，也即是说，当开始循环时，遇到左括号A则把其对应的右括号a入栈，而遇到右括号a，则出栈并判断是否和右括号a一样
// 两个成对的括号的位置一定是相对某个位置对称的，所以出栈的a一定得是之前A的成对者才行。
// O(N) Time，O(N) Space
func IsValidParentheses(s string) bool {
	// O(N+1) space
	// 放一串是为了避免一开始就是右括号的溢出
	stack := []string{""}
	// O(N) time
	for _, c := range s {
		// 如果 `c` 是左括号，则入栈
		if bc, ok := validParenthesesMap[string(c)]; ok {
			stack = append(stack, bc)
		} else {
			// 那就是右括号，得出栈，看是否成对
			if len(stack) > 1 && stack[len(stack)-1] == string(c) {
				stack = stack[:len(stack)-1]
				continue
			} else {
				return false
			}
		}
	}

	return 1 == len(stack)
}

func isValidParentheses(s string) bool {
	l := len(s)

	// 非偶数长度
	if l%2 != 0 || l == 0 {
		return false
	}

	m := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	stack := make([]rune, 0, l/2)
	for i, c := range s {
		if v, ok := m[c]; ok {
			stack = append(stack, v)
		} else if len(stack) == 0 || c != stack[len(stack)-1] {
			return false
		} else {
			// 出栈
			stack = stack[:len(stack)-1]
		}

		// 还没到头栈空间已经用完说明不是，对称的肯定只会用一半
		if i-cap(stack) > 1 && len(stack) == cap(stack) {
			return false
		}
	}

	return len(stack) == 0
}

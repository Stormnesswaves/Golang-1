func isValid(s string) bool {
    // 长度为奇数，一定不可能完全配对
    if len(s)%2 == 1 {
        return false
    }

    // 用切片当栈，存放左括号
    stack := make([]rune, 0, len(s))

    // 遍历字符串中的每一个字符
    for _, ch := range s {
        switch ch {
        // 如果是左括号，就压栈
        case '(', '{', '[':
            stack = append(stack, ch)

        // 如果是右括号，就检查栈顶是否有对应的左括号
        case ')', '}', ']':
            // 栈为空，说明右括号多了
            if len(stack) == 0 {
                return false
            }

            // 取出栈顶元素
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1] // 弹栈

            // 检查是否匹配
            if (ch == ')' && top != '(') ||
                (ch == ']' && top != '[') ||
                (ch == '}' && top != '{') {
                return false
            }
        }
    }

    // 所有字符处理完后，如果栈是空的，说明全部成功配对
    return len(stack) == 0
}

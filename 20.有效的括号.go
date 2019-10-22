/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 *
 * https://leetcode-cn.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (39.54%)
 * Likes:    1132
 * Dislikes: 0
 * Total Accepted:    141.5K
 * Total Submissions: 355K
 * Testcase Example:  '"()"'
 *
 * 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
 *
 * 有效字符串需满足：
 *
 *
 * 左括号必须用相同类型的右括号闭合。
 * 左括号必须以正确的顺序闭合。
 *
 *
 * 注意空字符串可被认为是有效字符串。
 *
 * 示例 1:
 *
 * 输入: "()"
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: "()[]{}"
 * 输出: true
 *
 *
 * 示例 3:
 *
 * 输入: "(]"
 * 输出: false
 *
 *
 * 示例 4:
 *
 * 输入: "([)]"
 * 输出: false
 *
 *
 * 示例 5:
 *
 * 输入: "{[]}"
 * 输出: true
 *
 */
package main

import (
	"errors"
)

// @lc code=start
func isValid(s string) bool {
	stack := NewStack()
	for _, value := range s {
		switch value {
		case '(', '[', '{':
			stack.Push(value)
		case ')':
			last, err := stack.Pop()
			if err != nil || last != '(' {
				return false
			}
		case ']':
			last, err := stack.Pop()
			if err != nil || last != '[' {
				return false
			}
		case '}':
			last, err := stack.Pop()
			if err != nil || last != '{' {
				return false
			}
		}
	}
	return stack.IsEmpty()

}

type Stack struct {
	elements []interface{}
}

func NewStack() Stack {
	stack := Stack{}
	return stack
}

// 判断栈是否为空
func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

// 入栈
func (s *Stack) Push(data interface{}) {
	// 把当前的元素放在栈顶的位置
	s.elements = append(s.elements, data)
}

// 出栈
func (s *Stack) Pop() (interface{}, error) {
	// 判断是否是空栈
	if s.IsEmpty() {
		return nil, errors.New("Stack: This Stack is empty , pop error")
	}
	// 把栈顶的元素赋值给临时变量tmp
	tmp := s.elements[s.Size()-1]

	s.elements = s.elements[:s.Size()-1]

	return tmp, nil
}

// 栈的元素的长度
func (s *Stack) Size() int {
	size := len(s.elements)
	return size
}

// 清空栈
func (s *Stack) Clear() {
	s.elements = nil
}

// @lc code=end

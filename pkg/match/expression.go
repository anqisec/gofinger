package match

import (
	"github.com/fuyoumingyan/gofinger/pkg/module"
	"regexp"
	"strconv"
)

// splitUnits 解析 rule 并将其转换为表达式
func splitUnits(expression string, info module.Info) []string {
	pattern := `(\w+\s*(!?=)\s*"(?:\\"|[^"])*"\s*)|([&,|]{2})|[(,)]`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(expression, -1)
	var units []string
	for _, match := range matches {
		match = unEscapeAndSpace(match)
		if match != "" {
			if match == "||" || match == "&&" || match == "(" || match == ")" {
				units = append(units, match)
			} else {
				units = append(units, strconv.FormatBool(matchSingleRule(match, info)))
			}
		}
	}
	return units
}
func infixToPostfix(expression string, info module.Info) []string {
	operatorPrecedence := map[string]int{
		"||": 1,
		"&&": 2,
	}
	units := splitUnits(expression, info)
	var output []string
	var stack module.Stack

	for _, token := range units {
		if token == "(" {
			// 左括号 => 入栈
			stack.Push(token)
		} else if token == ")" {
			// 右括号 => 一直弹出, 直到遇到左括号
			for stack.Top() != "(" && stack.Top() != "" {
				output = append(output, stack.Pop())
			}
			// 弹出左括号
			stack.Pop()
		} else if operatorPrecedence[token] > 0 {
			// 如果是该符号优先级比栈顶低 => 弹出
			for stack.Top() != "" && operatorPrecedence[token] <= operatorPrecedence[stack.Top()] {
				output = append(output, stack.Pop())
			}
			// 该符号优先级比栈顶高了, 入栈
			stack.Push(token)
		} else {
			// 不是符号, 直接入栈
			output = append(output, token)
		}
	}
	// 遍历万之后, 将栈这符号弹出
	for stack.Top() != "" {
		output = append(output, stack.Pop())
	}
	return output
}

package match

import (
	"github.com/fuyoumingyan/gofinger/core/module"
	"strconv"
	"strings"
)

func MatchRules(rules string, info module.Info) bool {
	stack := module.Stack{}
	tokens := infixToPostfix(rules, info)
	var c bool
	if len(tokens) == 1 {
		parseBool, err := strconv.ParseBool(tokens[0])
		if err != nil {
			return false
		}
		return parseBool
	}
	for i := 0; i < len(tokens); i++ {
		if strings.Contains(tokens[i], "||") || strings.Contains(tokens[i], "&&") {
			a, _ := strconv.ParseBool(stack.Pop())
			b, _ := strconv.ParseBool(stack.Pop())
			if strings.Contains(tokens[i], "||") {
				c = a || b
			} else {
				c = a && b
			}
			stack.Push(strconv.FormatBool(c))
		} else {
			stack.Push(tokens[i])
		}
	}
	return c
}

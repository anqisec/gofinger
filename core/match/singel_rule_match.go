package match

import (
	"gofinger/core/module"
	"regexp"
	"strings"
)

// matchSingleRule 匹配单个规则
func matchSingleRule(rule string, info module.Info) bool {
	rule = unEscapeAndSpace(rule)
	re := regexp.MustCompile(`"(.*?)"$`)
	match := re.FindStringSubmatch(rule)
	var ruleStr string
	if len(match) > 1 {
		ruleStr = strings.TrimSpace(match[1])
	} else {
		return false
	}
	if len(strings.TrimSpace(ruleStr)) == 0 {
		return false
	}
	if strings.Contains(rule, "title") {
		return matchEqual(rule, matchTitle(info.Title, ruleStr))
	}
	if strings.Contains(rule, "body") {
		return matchEqual(rule, matchBody(info.Body, ruleStr))
	}
	if strings.Contains(rule, "header") || strings.Contains(rule, "banner") || strings.Contains(rule, "server") {
		return matchEqual(rule, matchHeader(info.Header, ruleStr))
	}
	if strings.Contains(rule, "icon_hash") {
		return matchEqual(rule, matchIcoHash(info.IcoHashs, ruleStr))
	}
	if strings.Contains(rule, "cert") {
		return matchEqual(rule, matchCert(info.Cert, ruleStr))
	}
	return false
}

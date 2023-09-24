package request

import (
	"regexp"
)

func GetJSRedirectURL(body string) string {
	rules := []string{
		`window\.location\.href\("([^"]+)"\)`,
		`window\.location\.replace\("([^"]+)"\)`,
		`location\.replace\("([^"]+)"\)`,
	}
	var redirectURL string
	for _, rule := range rules {
		re := regexp.MustCompile(rule)
		matches := re.FindStringSubmatch(body)
		if len(matches) > 1 {
			redirectURL = matches[1]
		}
	}
	return redirectURL
}

package request

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"regexp"
	"strings"
)

func GetTitle(body string) string {
	re := regexp.MustCompile("<title>(.*?)</title>")
	matches := re.FindStringSubmatch(body)
	var title string
	if len(matches) >= 2 {
		title = matches[1]
		if len(strings.TrimSpace(title)) != 0 {
			return title
		}
	} else {
		document, err := goquery.NewDocumentFromReader(strings.NewReader(body))
		if err != nil {
			log.Println(err)
			return "<nil>"
		}
		title = document.Find("title").Text()
		title = strings.ReplaceAll(title, "\n", "")
		title = strings.TrimSpace(title)
		if len(strings.TrimSpace(title)) != 0 {
			return title
		}
	}
	return "<nil>"
}

package template

import (
	_ "embed"
	"github.com/fuyoumingyan/gofinger/core/module"
	"html/template"
	"log"
	"os"
	"strings"
)

//go:embed template.html
var htmlTemplate string

func GetHtmlResult(screenshotResult []module.Result) {
	for i := range screenshotResult {
		screenshotResult[i].Screenshot = strings.ReplaceAll(screenshotResult[i].Screenshot, "/result", "")
	}
	tmpl, err := template.New("html").Parse(htmlTemplate)
	if err != nil {
		return
	}
	file, err := os.Create("./result/results.html")
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	// 将数据应用到模板并输出到文件
	err = tmpl.Execute(file, screenshotResult)
	if err != nil {
		log.Println(err)
	}
}

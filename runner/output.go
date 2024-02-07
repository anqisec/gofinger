package runner

import (
	"encoding/csv"
	"fmt"
	"github.com/fuyoumingyan/gofinger/pkg/module"
	"github.com/fuyoumingyan/gofinger/pkg/options"
	out "github.com/fuyoumingyan/gofinger/pkg/output"
	"github.com/fuyoumingyan/gofinger/pkg/template"
	"github.com/fuyoumingyan/gofinger/pkg/utils"
	"github.com/projectdiscovery/gologger"
	"os"
	"strings"
)

type output struct {
	file             *os.File
	writer           *csv.Writer
	builder          *strings.Builder
	option           *options.Options
	fingerRunner     *FingerRunner
	requestRunner    *RequestRunner
	saveCSV          bool
	saveHtml         bool
	screenshotResult []module.Result
	windowsWidth     int
}

func NewOutputRunner(option *options.Options, fingerRunner *FingerRunner, requestRunner *RequestRunner) *output {
	o := new(output)
	o.option = option
	if len(option.Urls) > 1 {
		file, err := os.Create("./result/results.csv")
		if err != nil {
			gologger.Error().Msg(err.Error())
		}
		_, err = file.WriteString("\xEF\xBB\xBF")
		if err != nil {
			gologger.Error().Msg(err.Error())
		}
		o.file = file
		o.writer = csv.NewWriter(o.file)
		o.saveCSV = true
		o.writer.Write([]string{"url", "Title", "Finger"})
	}
	if option.Screenshot {
		o.saveHtml = true
	}
	o.builder = new(strings.Builder)
	o.windowsWidth = out.GetWindowWith()
	o.fingerRunner = fingerRunner
	o.requestRunner = requestRunner
	return o
}
func (o *output) RunEnumeration() {
	for result := range o.fingerRunner.result {
		if o.saveCSV {
			o.writer.Write([]string{result.Url, result.Title, result.Fingers})
		}
		if o.saveHtml {
			result.Screenshot = fmt.Sprintf("./result/screenshots/%v.png", utils.Md5(result.Url))
			o.screenshotResult = append(o.screenshotResult, result)
		}
		o.builder.WriteString(result.Url)
		o.builder.WriteString(" [ ")
		o.builder.WriteString(result.Title)
		o.builder.WriteString(" ] [ ")
		o.builder.WriteString(result.Fingers)
		o.builder.WriteString(" ]")
		o.Print(o.builder.String())
		fmt.Fprintf(os.Stdout, "All: %d RequestSuccess: %d RequestFaild: %d GetFinger: %d\r", o.requestRunner.allIndex, o.requestRunner.successIndex, o.requestRunner.faildIndex, o.fingerRunner.index)
	}
	if o.saveCSV {
		o.writer.Flush()
		o.file.Close()
	}
	o.Print("fingerprint identification complete .")
	if o.option.Screenshot {
		o.Print("Start taking screenshots of URLs.")
	}
	if o.saveHtml {
		screenshotRunner := NewScreenshotRunner(o.screenshotResult)
		screenshotRunner.RunEnumeration()
		template.GetHtmlResult(o.screenshotResult)
	}
	if o.option.Screenshot {
		o.Print("Screenshots completed.")
	}
}

func (o *output) Print(str string) {
	o.builder.Reset()
	o.builder.WriteString(str)
	screenWidth := o.windowsWidth - len(o.builder.String()) - 30
	for screenWidth > 0 {
		o.builder.WriteByte(' ')
		screenWidth--
	}
	gologger.Info().Msgf("%s\n", o.builder.String())
	o.builder.Reset()
}

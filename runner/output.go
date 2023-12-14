package runner

import (
	"encoding/csv"
	"fmt"
	"gofinger/core/options"
	out "gofinger/core/output"
	"log"
	"os"
	"strings"
)

type output struct {
	file          *os.File
	writer        *csv.Writer
	builder       *strings.Builder
	option        *options.Options
	fingerRunner  *FingerRunner
	requestRunner *RequestRunner
	wirteToFile   bool
	windowsWidth  int
}

func NewOutputRunner(option *options.Options, fingerRunner *FingerRunner, requestRunner *RequestRunner) *output {
	o := new(output)
	o.option = option
	if option.Output != "" {
		file, err := os.Create(option.Output)
		if err != nil {
			log.Println(err)
		}
		file.WriteString("\xEF\xBB\xBF")
		o.file = file
		o.writer = csv.NewWriter(o.file)
		o.wirteToFile = true
		o.writer.Write([]string{"url", "Title", "Finger"})
	}
	o.builder = new(strings.Builder)
	o.windowsWidth = out.GetWindowWith()
	o.fingerRunner = fingerRunner
	o.requestRunner = requestRunner
	return o
}
func (o *output) RunEnumeration() {
	for result := range o.fingerRunner.result {
		if o.wirteToFile {
			_ = o.writer.Write([]string{result.Url, result.Title, result.Fingers})
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
	if o.wirteToFile {
		o.writer.Flush()
		o.file.Close()
		o.Print(fmt.Sprintf("The results are saved in %v .", o.option.Output))
	}
	o.Print("fingerprint identification complete .")
}

func (o *output) Print(str string) {
	o.builder.Reset()
	o.builder.WriteString(str)
	screenWidth := o.windowsWidth - len(o.builder.String()) - 30
	for screenWidth > 0 {
		o.builder.WriteByte(' ')
		screenWidth--
	}
	log.Printf("%s\n", o.builder.String())
	o.builder.Reset()
}

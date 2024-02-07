package runner

import (
	"github.com/fuyoumingyan/gofinger/pkg/options"
	"github.com/fuyoumingyan/gofinger/pkg/template"
	"os"
	"os/signal"
)

type Runner struct {
	requestRunner *RequestRunner
	fingerRunner  *FingerRunner
	output        *output
}

func NewRunner(options *options.Options) *Runner {
	r := new(Runner)
	r.requestRunner = NewRequestRunner(options)
	r.fingerRunner = NewFingerRunner(options, r.requestRunner)
	r.output = NewOutputRunner(options, r.fingerRunner, r.requestRunner)
	return r
}
func (r *Runner) RunEnumeration() {
	go r.fingerRunner.RunEnumeration()
	go r.requestRunner.RunEnumeration()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			r.output.Print("CTRL+C pressed: Exiting")
			if r.output.saveCSV {
				r.output.writer.Flush()
				r.output.file.Close()
				template.GetHtmlResult(r.output.screenshotResult)
				r.output.Print("The results are saved in ./result/results.csv .")
				r.output.Print("The results are saved in ./result/results.html .")
			}
			os.Exit(1)
		}
	}()
	r.output.RunEnumeration()
}

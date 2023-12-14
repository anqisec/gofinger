package runner

import (
	"gofinger/core/options"
	"log"
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
			log.Println("CTRL+C pressed: Exiting")
			if r.output.wirteToFile {
				r.output.writer.Flush()
				r.output.file.Close()
				log.Printf("The results are saved in %v .", r.output.option.Output)
			}
			os.Exit(1)
		}
	}()
	r.output.RunEnumeration()
}

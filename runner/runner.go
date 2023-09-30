package runner

import (
	"gofinger/core/options"
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
	r.output.RunEnumeration()
}

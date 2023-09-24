package runner

import (
	"gofinger/core/options"
	"sync"
)

type Runner struct {
	requestRunner *RequestRunner
	fingerRunner  *FingerRunner
	output        *output
	wg            sync.WaitGroup
}

func NewRunner(options *options.Options) *Runner {
	r := new(Runner)
	r.requestRunner = NewRequestRunner(options)
	r.fingerRunner = NewFingerRunner(options, r.requestRunner)
	r.output = NewOutputRunner(options, r.fingerRunner, r.requestRunner)
	r.wg = sync.WaitGroup{}
	return r
}
func (r *Runner) RunEnumeration() {
	go r.requestRunner.RunEnumeration()
	go r.fingerRunner.RunEnumeration()
	r.output.RunEnumeration()
}

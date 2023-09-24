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
	// 启动请求协程
	go r.requestRunner.RunEnumeration()
	// 启动指纹识别协程
	go r.fingerRunner.RunEnumeration()
	// 启动输出
	r.output.RunEnumeration()
}

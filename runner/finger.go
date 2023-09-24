package runner

import (
	"gofinger/core/data"
	"gofinger/core/match"
	"gofinger/core/module"
	"gofinger/core/options"
	"strings"
	"sync"
	"sync/atomic"
)

type FingerRunner struct {
	wg            sync.WaitGroup
	option        *options.Options
	limit         chan struct{}
	fingerDatas   module.FingerData
	result        chan module.Result
	requestRunner *RequestRunner
	index         uint64
}

func NewFingerRunner(option *options.Options, requestRunner *RequestRunner) *FingerRunner {
	f := new(FingerRunner)
	f.option = option
	f.wg = sync.WaitGroup{}
	f.fingerDatas = data.GetFingerData(option)
	f.limit = make(chan struct{}, 500)
	f.requestRunner = requestRunner
	f.result = make(chan module.Result, len(requestRunner.Targets))
	return f
}
func (f *FingerRunner) RunEnumeration() {
	for info := range f.requestRunner.UrlInfo {
		f.limit <- struct{}{}
		f.wg.Add(1)
		go f.run(info)
	}
	f.wg.Wait()
	close(f.result)
}

func (f *FingerRunner) run(info module.Info) {
	defer f.wg.Done()
	var fingers []string
	for _, fingerData := range f.fingerDatas {
		if match.MatchRules(fingerData.Rule, info) {
			fingers = append(fingers, fingerData.CMS)
		}
	}
	if len(fingers) == 0 {
		fingers = append(fingers, "<nil>")
	}
	reslut := module.Result{
		Url:     info.Url,
		Title:   info.Title,
		Fingers: strings.Join(fingers, ", "),
	}
	f.result <- reslut
	atomic.AddUint64(&f.index, 1)
}

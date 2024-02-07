package runner

import (
	"github.com/fuyoumingyan/gofinger/pkg/module"
	"github.com/fuyoumingyan/gofinger/pkg/options"
	"github.com/fuyoumingyan/gofinger/pkg/request"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
)

type RequestRunner struct {
	Options      *options.Options
	Client       http.Client
	UrlInfo      chan module.Info
	Targets      chan string
	wg           sync.WaitGroup
	successIndex uint64
	faildIndex   uint64
	allIndex     uint64
	limit        chan struct{}
}

func NewRequestRunner(options *options.Options) *RequestRunner {
	r := new(RequestRunner)
	r.Options = options
	r.Client = request.GetClient(options)
	r.wg = sync.WaitGroup{}
	r.limit = make(chan struct{}, options.Thread)
	r.Targets = make(chan string, len(options.Urls))
	for _, url := range options.Urls {
		r.Targets <- strings.TrimSpace(url)
	}
	close(r.Targets)
	r.allIndex = uint64(len(r.Targets))
	r.UrlInfo = make(chan module.Info, len(r.Targets))
	return r
}

func (r *RequestRunner) RunEnumeration() {
	defer close(r.limit)
	r.wg.Add(len(r.Targets))
	for target := range r.Targets {
		r.limit <- struct{}{}
		go r.run(target)
	}
	r.wg.Wait()
	close(r.UrlInfo)
}
func (r *RequestRunner) run(url string) {
	defer func() {
		<-r.limit
		r.wg.Done()
	}()
	info := request.SendRequest(url, r.Client)
	if len(info.Url) == 0 {
		atomic.AddUint64(&r.faildIndex, 1)
		return
	}
	r.UrlInfo <- info
	atomic.AddUint64(&r.successIndex, 1)
}

package runner

import (
	"github.com/go-rod/rod"
	"gofinger/core/module"
	"gofinger/core/screenshot"
	"log"
	"sync"
)

type ScreenshotRunner struct {
	Browser          *rod.Browser
	wg               sync.WaitGroup
	limit            chan struct{}
	screenshotResult []module.Result
}

func NewScreenshotRunner(screenshotResult []module.Result) *ScreenshotRunner {
	s := new(ScreenshotRunner)
	s.wg = sync.WaitGroup{}
	s.limit = make(chan struct{}, 10)
	s.Browser = screenshot.GetBrowser()
	s.screenshotResult = screenshotResult
	return s
}

func (s *ScreenshotRunner) RunEnumeration() {
	for _, result := range s.screenshotResult {
		s.limit <- struct{}{}
		s.wg.Add(1)
		go s.run(result)
	}
	s.wg.Wait()
}

func (s *ScreenshotRunner) run(result module.Result) {
	defer func(s *ScreenshotRunner) {
		<-s.limit
		s.wg.Done()
	}(s)
	err := screenshot.GetScreenshot(s.Browser, result)
	if err != nil {
		log.Printf("%s screenshot error : %v", result.Url, err)
	} else {
		log.Printf("%s screenshot successfully !", result.Url)
	}
}

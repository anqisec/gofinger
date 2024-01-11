package screenshot

import (
	"github.com/fuyoumingyan/gofinger/core/module"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
	"os"
	"time"
)

func GetBrowser() *rod.Browser {
	browser := rod.New().MustConnect().MustIncognito()
	err := browser.ControlURL("").IgnoreCertErrors(true)
	if err != nil {
		return nil
	}
	return browser
}

func GetScreenshot(browser *rod.Browser, result module.Result) error {
	page, err := browser.Timeout(2 * time.Minute).Page(proto.TargetCreateTarget{URL: result.Url})
	if err != nil {
		return err
	}
	defer page.Close()
	err = page.WaitStable(3 * time.Second)
	//err = page.WaitLoad()
	if err != nil {
		return err
	}
	view := proto.EmulationSetDeviceMetricsOverride{
		Width:  1920,
		Height: 1080,
	}
	err = page.SetViewport(&view)
	if err != nil {
		return err
	}
	screenshot, err := page.Screenshot(false, nil)
	if err != nil {
		return err
	}

	err = os.WriteFile(result.Screenshot, screenshot, 0666)
	if err != nil {
		return err
	}
	return nil
}

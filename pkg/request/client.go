package request

import (
	"crypto/tls"
	"errors"
	"github.com/fuyoumingyan/gofinger/pkg/options"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func GetClient(options *options.Options) http.Client {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	checkRedirect := func(req *http.Request, via []*http.Request) error {
		for _, prevReq := range via {
			if !strings.Contains(req.Header.Get("Cookie"), "rememberMe") {
				req.Header.Add("Cookie", prevReq.Header.Get("Cookie"))
			}
		}
		if len(via) >= 3 {
			return errors.New("stopped after 3 redirects")
		}
		return nil
	}
	if options.Proxy != "" {
		proxyFunc := func(_ *http.Request) (*url.URL, error) {
			return url.Parse(options.Proxy)
		}
		transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			Proxy:           proxyFunc,
		}
	}
	client := http.Client{
		Timeout:       time.Duration(options.Timeout) * time.Second,
		Transport:     transport,
		CheckRedirect: checkRedirect,
	}
	return client
}

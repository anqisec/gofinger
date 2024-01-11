package request

import (
	"github.com/fuyoumingyan/gofinger/core/module"
	"github.com/fuyoumingyan/gofinger/core/utils"
	"html"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func SendRequest(urlStr string, client http.Client) module.Info {
	if !strings.HasPrefix(urlStr, "http") {
		urlStr = "http://" + urlStr
	}
	parse, err := url.Parse(urlStr)
	if err != nil {
		log.Println(err)
		return module.Info{}
	}
	urlStr = utils.GetHealthUrl(parse)
	ip := utils.GetIP(parse)
	request, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		log.Println(err)
		return module.Info{}
	}
	cookie := &http.Cookie{
		Name:  "rememberMe",
		Value: "me",
	}
	request.AddCookie(cookie)
	request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	request.Header.Set("User-Agent", GetRandomUA())
	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return module.Info{}
	}
	body := html.UnescapeString(GetBody(response))
	title := GetTitle(body)
	response.Body.Close()
	redirectURL := GetJSRedirectURL(urlStr, body)
	if redirectURL != "" && title == "<nil>" {
		return SendRequest(redirectURL, client)
	}
	icoHashs := getICOHash(response.Request.URL.String(), body, client)
	info := module.Info{
		Url:        urlStr,
		Title:      title,
		Body:       body,
		Header:     GetHeader(response),
		IcoHashs:   icoHashs,
		Cert:       getCertContent(urlStr),
		UniqueHash: utils.Md5(body + ip),
	}
	return info
}

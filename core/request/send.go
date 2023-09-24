package request

import (
	"gofinger/core/module"
	"log"
	"net/http"
	"strings"
)

func SendRequest(url string, client http.Client) module.Info {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()
	if !strings.Contains(url, "http") {
		log.Println(url, " is an invalid URL")
		return module.Info{}
	}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return module.Info{}
	}
	cookie := &http.Cookie{
		Name:  "rememberMe",
		Value: "me",
	}
	request.AddCookie(cookie)
	request.Header.Set("Accept", "*/*;q=0.8")
	request.Header.Set("Connection", "close")
	request.Header.Set("User-Agent", GetRandomUA())
	response, err := client.Do(request)
	defer response.Body.Close()
	body := GetBody(response)
	if redirectURL := GetJSRedirectURL(body); redirectURL != "" {
		return SendRequest(redirectURL, client)
	}
	icoHashs := getICOHash(response.Request.URL.String(), body, client)
	info := module.Info{
		Url:      url,
		Title:    GetTitle(body),
		Body:     body,
		Header:   GetHeader(response),
		IcoHashs: icoHashs,
		Cert:     getCertContent(url),
	}
	return info
}

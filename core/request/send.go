package request

import (
	"gofinger/core/module"
	"gofinger/core/utils"
	"log"
	"net/http"
)

func SendRequest(url string, client http.Client) module.Info {
	url = utils.AddSchemeIfNotExists(url)
	if url == "" {
		log.Printf("%s is an invalid url .\n", url)
		return module.Info{}
	}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
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
	if err != nil {
		log.Println(err.Error())
		return module.Info{}
	}
	body := GetBody(response)
	response.Body.Close()
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

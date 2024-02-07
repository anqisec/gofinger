package request

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/projectdiscovery/gologger"
	"github.com/twmb/murmur3"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// getICOHash 获取 ico hash
func getICOHash(url string, body string, client http.Client) []string {
	var icoHashs []string
	icoHash := getSingleICOHash(joinURLAndPath(url, "favicon.ico"), client)
	if icoHash != "" {
		icoHashs = append(icoHashs, icoHash)
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		gologger.Error().Msg(err.Error())
		return icoHashs
	}
	rules := []string{
		"link[rel=\"shortcut icon\"]",
		"link[rel=\"icon\"]",
	}
	var icoUrl string
	for _, rule := range rules {
		iconPath, exists := doc.Find(rule).Attr("href")
		if exists {
			if strings.Contains(iconPath, "http") {
				icoUrl = iconPath
			} else {
				icoUrl = joinURLAndPath(url, iconPath)
			}
		}
	}
	icoHash = getSingleICOHash(icoUrl, client)
	if icoHash != "" {
		icoHashs = append(icoHashs, icoHash)
	}
	return icoHashs
}

// getSingleICOHash 获取单个 ico 的 hash
func getSingleICOHash(icoUrl string, client http.Client) string {
	resp, err := client.Get(icoUrl)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return ""
		}
		return Mmh3Hash32(StandBase64(body))
	}
	return ""
}

// joinURLAndPath 拼接 URL 和路径
func joinURLAndPath(baseURL string, path string) string {
	parsedBaseURL, err := url.Parse(baseURL)
	if err != nil {
		gologger.Error().Msg(err.Error())
		return ""
	}
	var builder strings.Builder
	builder.WriteString(parsedBaseURL.Scheme)
	builder.WriteString("://")
	builder.WriteString(parsedBaseURL.Host)
	baseUrl := builder.String()
	builder.Reset()
	builder.WriteString(baseUrl)
	builder.WriteString("/")
	builder.WriteString(strings.TrimLeft(path, "/"))
	return builder.String()
}

// Mmh3Hash32 32 位 mmh3 hash
func Mmh3Hash32(raw []byte) string {
	var h32 = murmur3.New32()
	_, err := h32.Write(raw)
	if err == nil {
		return fmt.Sprintf("%d", int32(h32.Sum32()))
	} else {
		return "0"
	}
}

// StandBase64 Base64 编码, 每 76 字符加上换行符
func StandBase64(braw []byte) []byte {
	bckd := base64.StdEncoding.EncodeToString(braw)
	var buffer bytes.Buffer
	for i := 0; i < len(bckd); i++ {
		ch := bckd[i]
		buffer.WriteByte(ch)
		if (i+1)%76 == 0 {
			buffer.WriteByte('\n')
		}
	}
	buffer.WriteByte('\n')
	return buffer.Bytes()
}

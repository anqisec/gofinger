package request

import (
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"net/http"
	"strings"
)

func GetBody(response *http.Response) string {
	contentType := response.Header.Get("Content-Type")
	if isNotAudioVideoFileContentType(contentType) {
		if strings.Contains(strings.ToLower(contentType), "utf-8") {
			body, err := io.ReadAll(response.Body)
			if err != nil {
				return ""
			}
			return string(body)
		}
		bodyReader := bufio.NewReader(response.Body)
		e := determiEncoding(bodyReader)
		utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
		utf8Body, err := io.ReadAll(utf8Reader)
		if err != nil {
			return ""
		}
		return string(utf8Body)
	}
	return ""
}

// isNotAudioVideoFileContentType 排除音频、视频、文件
func isNotAudioVideoFileContentType(contentType string) bool {
	contentType = strings.ToLower(contentType)
	audioPrefixes := []string{"audio/", "application/ogg", "application/x-mpegurl"}
	videoPrefixes := []string{"video/", "application/x-mpegurl"}
	filePrefixes := []string{"application/octet-stream", "application/pdf"}
	for _, prefix := range append(audioPrefixes, append(videoPrefixes, filePrefixes...)...) {
		if strings.HasPrefix(contentType, prefix) {
			return false // 匹配到音频、视频、文件的Content-Type，返回false
		}
	}
	return true
}

// determiEncoding 获取编码类型
func determiEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

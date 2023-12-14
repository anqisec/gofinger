package request

import (
	"bufio"
	"bytes"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"net/http"
	"regexp"
	"strings"
)

//func GetBody(response *http.Response) string {
//	contentType := response.Header.Get("Content-Type")
//	if isNotAudioVideoFileContentType(contentType) {
//		if strings.Contains(strings.ToLower(contentType), "utf-8") {
//			body, err := io.ReadAll(response.Body)
//			if err != nil {
//				return ""
//			}
//			return string(body)
//		}
//		println(contentType)
//		bodyReader := bufio.NewReader(response.Body)
//		e := determiEncoding(bodyReader)
//		println(e)
//		utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
//		utf8Body, err := io.ReadAll(utf8Reader)
//		if err != nil {
//			return ""
//		}
//		return string(utf8Body)
//	}
//	return ""
//}

func DecodeAuto(response *http.Response) string {
	bodyReader := bufio.NewReader(response.Body)
	e := determiEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	utf8Body, err := io.ReadAll(utf8Reader)
	if err != nil {
		return ""
	}
	return string(utf8Body)
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

// Decodegbk converts GBK to UTF-8
func Decodegbk(s []byte) string {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, simplifiedchinese.GBK.NewDecoder())
	d, e := io.ReadAll(O)
	if e != nil {
		return ""
	}
	return string(d)
}

func DecodeKorean(s []byte) string {
	koreanDecoder := korean.EUCKR.NewDecoder()
	d, e := koreanDecoder.Bytes(s)
	if e != nil {
		return ""
	}
	return string(d)
}

// GetBody DecodeData ExtractTitle from a response
func GetBody(response *http.Response) string {
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return ""
	}
	contentTypes, ok := response.Header["Content-Type"]
	if ok {
		contentType := strings.ToLower(strings.Join(contentTypes, ";"))
		if !isNotAudioVideoFileContentType(contentType) {
			return ""
		}
		switch {
		case strings.Contains(contentType, "charset=gb2312") || strings.Contains(contentType, "charset=gbk"):
			return Decodegbk(data)
		case strings.Contains(contentType, "euc-kr"):
			return DecodeKorean(data)
		case strings.Contains(contentType, "utf-8"):
			return string(data)
		}
		reContentType := regexp.MustCompile(`(?im)\s*charset="(.*?)"|charset=(.*?)"\s*`)
		matches := reContentType.FindStringSubmatch(string(data))
		var mcontentType = ""
		if len(matches) != 0 {
			for i, v := range matches {
				if string(v) != "" && i != 0 {
					mcontentType = v
				}
			}
			mcontentType = strings.ToLower(mcontentType)
		}
		switch {
		case strings.Contains(mcontentType, "gb2312") || strings.Contains(mcontentType, "gbk"):
			return Decodegbk(data)
		default:
			return string(data)
		}
	}
	return string(data)
}

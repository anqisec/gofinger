package utils

import (
	"bufio"
	"net/url"
	"os"
	"strings"
)

func DeduplicateEmptyStrings(slice []string) []string {
	seen := make(map[string]bool)
	result := []string{}
	for _, s := range slice {
		if s != "" && !seen[s] {
			seen[s] = true
			result = append(result, s)
		}
	}
	return result
}
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

func AddSchemeIfNotExists(inputURL string) string {
	if strings.HasPrefix(inputURL, "http") || strings.HasPrefix(inputURL, "https") {
		return inputURL
	}
	// ip 地址无协议会解析失败
	httpUrl := "http://" + inputURL
	parsed, err := url.Parse(httpUrl)
	if err != nil {
		return ""
	}
	if parsed.Port() != "" && (parsed.Port() == "443" || parsed.Port() == "8181" || parsed.Port() == "8443" || parsed.Port() == "9443") {
		return "https://" + inputURL
	} else {
		return httpUrl
	}
}

package utils

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"net"
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

func GetIP(parsed *url.URL) string {
	addrs, err := net.LookupIP(parsed.Hostname())
	if err != nil {
		return ""
	}
	ips := make([]string, 0, len(addrs))
	for _, addr := range addrs {
		ips = append(ips, addr.String())
	}
	return strings.Join(ips, ", ")
}

func GetHealthUrl(parsed *url.URL) string {
	if parsed.Scheme == "https" {
		return parsed.String()
	}
	if parsed.Port() != "" && (parsed.Port() == "443" || parsed.Port() == "8181" || parsed.Port() == "8443" || parsed.Port() == "9443") {
		parsed.Scheme = "https"
		return parsed.String()
	}
	if parsed.Port() != "" && parsed.Port() == "80" {
		return "http://" + parsed.Hostname()
	}
	return parsed.String()
}
func Md5(str string) string {
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}

func JoinURL(baseURL, path string) (string, error) {
	base, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}
	p, err := url.Parse(path)
	if err != nil {
		return "", err
	}
	finalURL := base.ResolveReference(p).String()
	return finalURL, nil
}
func Mkdir(path string) error {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		if os.RemoveAll(path) != nil {
			return err
		}
	}
	err := os.MkdirAll(path, 0750)
	if err != nil {
		return err
	}
	return nil
}

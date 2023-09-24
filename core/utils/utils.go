package utils

import (
	"bufio"
	"os"
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

package match

func matchTitle(title, ruleStr string) bool {
	if CaseInsensitiveContains(title, ruleStr) {
		return true
	}
	return false
}
func matchCert(cert, ruleStr string) bool {
	if CaseInsensitiveContains(cert, ruleStr) {
		return true
	}
	return false
}
func matchBody(body, ruleStr string) bool {
	if CaseInsensitiveContains(body, ruleStr) {
		return true
	}
	return false
}
func matchHeader(header, ruleStr string) bool {
	if CaseInsensitiveContains(header, ruleStr) {
		return true
	}
	return false
}
func matchIcoHash(icoHashs []string, ruleStr string) bool {
	if inSlice(ruleStr, icoHashs) {
		return true
	}
	return false
}

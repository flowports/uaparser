package uaparser

import "regexp"

var othersRegex = regexp.MustCompile(`YaBrowser|UC Browser|UCBrowser`)
var msChromiumEdgeRegex = regexp.MustCompile(`Edg\/([^\s]+)`)
var msOriginalEdgeRegex = regexp.MustCompile(`Edge\/([^\s]+)`)
var oldOperaRegex = regexp.MustCompile(`Opera\/([^\s]+)`)
var newOperaRegex = regexp.MustCompile(`OPR\/([^\s]+)`)
var safariRegex = regexp.MustCompile(`Safari\/([^\s]+)`)
var safariVersionRegex = regexp.MustCompile(`Version\/([^\s]+)`)
var chromeRegex = regexp.MustCompile(`Chrome\/([^\s]+)`)
var firefoxRegex = regexp.MustCompile(`Firefox\/([^\s]+)`)

func parseBrowserVersion(userAgent string, os string) (string, string) {
	if othersRegex.MatchString(userAgent) {
		return "Other", "Other"
	}

	if firefoxRegex.MatchString(userAgent) {
		match := firefoxRegex.FindStringSubmatch(userAgent)
		return "Firefox", match[1]
	}

	if msChromiumEdgeRegex.MatchString(userAgent) {
		match := msChromiumEdgeRegex.FindStringSubmatch(userAgent)
		return "Edge", match[1]
	}

	if msOriginalEdgeRegex.MatchString(userAgent) {
		match := msOriginalEdgeRegex.FindStringSubmatch(userAgent)
		return "Edge", match[1]
	}

	isApple := os == "iOS" || os == "macOS" || os == "iPadOS"
	if isApple && safariVersionRegex.MatchString(userAgent) && safariRegex.MatchString(userAgent) {
		match := safariVersionRegex.FindStringSubmatch(userAgent)
		return "Safari", match[1]
	}

	if newOperaRegex.MatchString(userAgent) {
		match := newOperaRegex.FindStringSubmatch(userAgent)
		return "Opera", match[1]
	}

	if oldOperaRegex.MatchString(userAgent) {
		match := oldOperaRegex.FindStringSubmatch(userAgent)
		return "Opera", match[1]
	}

	if chromeRegex.MatchString(userAgent) {
		match := chromeRegex.FindStringSubmatch(userAgent)
		return "Chrome", match[1]
	}

	return "Other", "Other"
}

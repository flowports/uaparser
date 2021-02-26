package uaparser

import "strings"

func parseOS(userAgent string) string {
	if strings.Contains(userAgent, "(Windows") {
		return "Windows"
	}
	if strings.Contains(userAgent, "(iPhone") {
		return "iOS"
	}
	if strings.Contains(userAgent, "(iPad") {
		return "iPadOS"
	}
	if strings.Contains(userAgent, "(Macintosh") {
		return "macOS"
	}
	if strings.Contains(userAgent, "Android") {
		return "Android"
	}
	if strings.Contains(userAgent, "FreeBSD") {
		return "FreeBSD"
	}
	if strings.Contains(userAgent, "Linux") {
		return "Linux"
	}
	return "Other"
}

package uaparser

// Browser is a model derived by a UserAgent string
type Browser struct {
	Name    string
	Version string
	OS      string
}

// Parse extracts browser information from a UserAgent string
func Parse(userAgent string) Browser {
	os := parseOS(userAgent)
	name, version := parseBrowserVersion(userAgent, os)

	return Browser{
		Name:    name,
		Version: version,
		OS:      os,
	}
}

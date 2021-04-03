package uaparser_test

import (
	"testing"

	"github.com/flowports/uaparser"
)

func TestValidUserAgents(t *testing.T) {
	testCases := []struct {
		input           string
		expectedName    string
		expectedOS      string
		expectedVersion string
	}{
		{
			input:           "Mozilla/5.0 (X11; FreeBSD amd64; rv:85.0) Gecko/20100101 Firefox/85.0",
			expectedName:    "Firefox",
			expectedOS:      "FreeBSD",
			expectedVersion: "85.0",
		},
		{
			input:           "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36",
			expectedName:    "Chrome",
			expectedOS:      "Windows",
			expectedVersion: "79.0.3945.130",
		},
		{
			input:           "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36 Edge/18.18363",
			expectedName:    "Edge",
			expectedOS:      "Windows",
			expectedVersion: "18.18363",
		},
		{
			input:           "Mozilla/5.0 (X11; CrOS x86_64 13310.93.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.133 Safari/537.36",
			expectedName:    "Chrome",
			expectedOS:      "Other",
			expectedVersion: "85.0.4183.133",
		},
		{
			input:           "Mozilla/5.0 (iPhone; CPU iPhone OS 13_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.1 Mobile/15E148 Safari/604.1",
			expectedName:    "Safari",
			expectedOS:      "iOS",
			expectedVersion: "13.1",
		},
		{
			input:           "Mozilla/5.0 (iPad; CPU OS 14_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 Mobile/15E148 Safari/604.1",
			expectedName:    "Safari",
			expectedOS:      "iPadOS",
			expectedVersion: "14.0",
		},
		{
			input:           "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36 Edg/86.0.622.58",
			expectedName:    "Edge",
			expectedOS:      "macOS",
			expectedVersion: "86.0.622.58",
		},
		{
			input:           "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:57.0) Gecko/20100101 Firefox/57.0",
			expectedName:    "Firefox",
			expectedOS:      "Linux",
			expectedVersion: "57.0",
		},
		{
			input:           "Mozilla/5.0 (Linux; Android 6.0.1; RedMi Note 5 Build/RB3N5C; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/68.0.3440.91 Mobile Safari/537.36",
			expectedName:    "Chrome",
			expectedOS:      "Android",
			expectedVersion: "68.0.3440.91",
		},
		{
			input:           "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.183 YaBrowser/20.9.0 Yowser/2.5 Safari/537.36",
			expectedName:    "Other",
			expectedOS:      "Windows",
			expectedVersion: "Other",
		},
		{
			input:           "Mozilla/5.0 (Linux; U; Android 6.0.1; zh-CN; F5121 Build/34.0.A.1.247) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/40.0.2214.89 UCBrowser/11.5.1.944 Mobile Safari/537.36",
			expectedName:    "Other",
			expectedOS:      "Android",
			expectedVersion: "Other",
		},
		{
			input:           "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.66 Safari/537.36 OPR/72.0.3815.378",
			expectedName:    "Opera",
			expectedOS:      "Windows",
			expectedVersion: "72.0.3815.378",
		},
		{
			input:           "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_0_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.66 Safari/537.36 OPR/72.0.3815.378",
			expectedName:    "Opera",
			expectedOS:      "macOS",
			expectedVersion: "72.0.3815.378",
		},
		{
			input:           "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.66 Safari/537.36 Opera/72.0.3815.378",
			expectedName:    "Opera",
			expectedOS:      "Linux",
			expectedVersion: "72.0.3815.378",
		},
		{
			input:           "Mozilla/5.0",
			expectedName:    "Other",
			expectedOS:      "Other",
			expectedVersion: "Other",
		},
	}

	for _, testCase := range testCases {
		browser := uaparser.Parse(testCase.input)
		if browser.Name != testCase.expectedName {
			t.Fatalf("%s != %s", browser.Name, testCase.expectedName)
		}

		if browser.OS != testCase.expectedOS {
			t.Fatalf("%s != %s", browser.OS, testCase.expectedOS)
		}

		if browser.Version != testCase.expectedVersion {
			t.Fatalf("%s != %s", browser.Version, testCase.expectedVersion)
		}
	}
}

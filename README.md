# github.com/flowports/uaparser

A minimalist UserAgent parser for popular agents.

📦 Zero Dependencies

## What it does

uaparser transforms a [UserAgent string](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/User-Agent) into a Browser object with its normalized Name, Version and Operating System name. Only popular OS and Browsers are supported.

- **Supported Operating Systems:** Windows, iOS, iPadOS, macOS, Android, FreeBSD and Linux
- **Supported Browsers:** Firefox, Edge, Safari, Opera, Chrome

## How to use

```go
import (
    "github.com/flowports/uaparser"
)

func main() {
    browser := uaparser.Parse("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36")

    println(browser.OS) // Windows
    println(browser.Name) // Chrome
    println(browser.Version) // 79.0.3945.130
}
```

Invalid or unknown UserAgents are parsed into a Browser object with `"Other"` on each field.

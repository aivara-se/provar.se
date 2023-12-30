package metadata

import "github.com/mileusna/useragent"

// ParseUserAgent parses a user agent string.
func ParseUserAgent(ua string) useragent.UserAgent {
	return useragent.Parse(ua)
}

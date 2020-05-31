package useragent

import (
	"github.com/hiddenyte/genius-cli/useragent"
)

var userAgents []useragent.UserAgent

// GetUserAgent retrieves a user agent that supports rendering the specifed address. Returns nil, if there is none.
func GetUserAgent(address string) useragent.UserAgent {
	for _, ua := range userAgents {
		if ua.Supports(address) {
			return ua
		}
	}
	return nil
}

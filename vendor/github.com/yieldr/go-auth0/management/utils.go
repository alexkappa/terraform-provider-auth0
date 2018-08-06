package management

import (
	"net/http"
	"strconv"
	"time"

	"github.com/PuerkitoBio/rehttp"
)

func wrapRetry(c *http.Client) *http.Client {
	if c == nil {
		c = http.DefaultClient
	}
	return &http.Client{
		Transport: rehttp.NewTransport(
			c.Transport,
			func(attempt rehttp.Attempt) bool {
				if attempt.Response == nil {
					return false
				}
				return attempt.Response.StatusCode == http.StatusTooManyRequests
			},
			func(attempt rehttp.Attempt) time.Duration {
				resetAt := attempt.Response.Header.Get("X-RateLimit-Reset")
				resetAtUnix, err := strconv.ParseInt(resetAt, 10, 64)
				if err != nil {
					resetAtUnix = time.Now().Add(5 * time.Second).Unix()
				}
				return time.Unix(resetAtUnix, 0).Sub(time.Now())
			},
		),
	}
}

func wrapUserAgent(c *http.Client) *http.Client {
	if c == nil {
		c = http.DefaultClient
	}
	return &http.Client{
		Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			r.Header.Set("User-Agent", "Go-Auth0-SDK/v0")
			return c.Transport.RoundTrip(r)
		}),
	}
}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (rf roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return rf(r)
}

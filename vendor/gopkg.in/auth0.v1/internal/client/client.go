package client

import (
	"context"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"time"

	"github.com/PuerkitoBio/rehttp"
	"golang.org/x/oauth2/clientcredentials"
)

func WrapRetry(c *http.Client) *http.Client {
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

func WrapUserAgent(c *http.Client) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(func(req *http.Request) (*http.Response, error) {
			req.Header.Set("User-Agent", "Go-Auth0-SDK/v0")
			return c.Transport.RoundTrip(req)
		}),
	}
}

type RoundTripFunc func(*http.Request) (*http.Response, error)

func (rf RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return rf(req)
}

func WrapDebug(c *http.Client) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(func(req *http.Request) (*http.Response, error) {
			res, err := c.Transport.RoundTrip(req)
			if err != nil {
				return res, err
			}
			reqBytes, _ := httputil.DumpRequest(req, true)
			resBytes, _ := httputil.DumpResponse(res, true)
			log.Printf("%s\n%s\b\n", reqBytes, resBytes)
			return res, nil
		}),
	}
}

func OAuth2(domain, clientID, clientSecret string) *http.Client {
	return (&clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     "https://" + domain + "/oauth/token",
		EndpointParams: url.Values{
			"audience": {"https://" + domain + "/api/v2/"},
		},
	}).Client(context.Background())
}

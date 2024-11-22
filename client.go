package htt

import (
	"crypto/tls"
	"github.com/sekiju/htt/round_tripper"
	"net/http"
	"time"
)

// Client is the global HTTP client used by the package.
var Client = DefaultClient()

// DefaultClient creates a new HTTP client with custom transport settings.
func DefaultClient() *http.Client {
	client := http.DefaultClient
	client.Transport = round_tripper.NewHeaderRoundTripper(
		&http.Transport{
			TLSClientConfig: &tls.Config{
				CurvePreferences: []tls.CurveID{tls.CurveP256, tls.CurveP384, tls.CurveP521, tls.X25519},
			},
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
		map[string]string{
			"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8",
			"Accept-Language": "en-US,en;q=0.5",
			"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:134.0) Gecko/20100101 Firefox/134.0",
		},
	)

	return client
}

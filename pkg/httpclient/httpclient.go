// pkg/httpclient/httpclient.go

// Add retry logic (with github.com/hashicorp/go-retryablehttp)

// Add tracing headers for distributed tracing

package httpclient

import (
	"net/http"
	"time"
)

func New() *http.Client {
	return &http.Client{
		Timeout: 5 * time.Second,
	}
}

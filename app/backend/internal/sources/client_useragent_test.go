package sources

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestClientUserAgent checks that Source requests name HolyArrs and not the
// Go http client's default, which site protection blocks as a bot.
func TestClientUserAgent(t *testing.T) {
	old := userAgent
	defer SetUserAgentVersion(old[len("HolyArrs/"):])
	SetUserAgentVersion("9.9.9")

	var got string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got = r.Header.Get("User-Agent")
		w.Write([]byte("<error></error>"))
	}))
	defer ts.Close()

	c := New(Config{URL: ts.URL, Format: "newznab", AllowLocal: true})
	_, _ = c.Caps(context.Background())
	if got != "HolyArrs/9.9.9" {
		t.Fatalf("User-Agent = %q, want HolyArrs/9.9.9", got)
	}
}

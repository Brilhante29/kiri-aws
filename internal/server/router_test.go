package server

import (
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestRouter_PrefixMatchRespectsBoundary regression-tests a bug where
// `extractRoutePrefix` and `Router.ServeHTTP` matched prefixes by raw
// string-prefix, so a path like `/kiri-audit-bad-bucket` was routed to
// the `/kiri` prefix router (which only handles `/_kiri/*` and
// similar), shadowing the S3 wildcard route `PUT /{bucket}` and
// returning 404. The fix requires the prefix to be followed by `/` or
// end-of-string.
func TestRouter_PrefixMatchRespectsBoundary(t *testing.T) {
	t.Parallel()

	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	r := NewRouter(logger)

	called := ""

	// `/kiri` is a registered prefix (used by /_kiri/health etc.).
	r.Handle("GET", "/kiri/health", func(w http.ResponseWriter, _ *http.Request) {
		called = "kiri-health"

		w.WriteHeader(http.StatusOK)
	})

	// `/{bucket}` is the S3-style wildcard route. With the bug, a
	// PUT to `/kiri-audit-bad-bucket` would be sent to the /kiri
	// prefix router (no matching pattern) → 404.
	r.Handle("PUT", "/{bucket}", func(w http.ResponseWriter, _ *http.Request) {
		called = "bucket-put"

		w.WriteHeader(http.StatusOK)
	})

	// A bucket whose name *equals* an admin prefix (e.g. PUT /kiri) is
	// an unavoidable shadow until kiri grows a Host- or scheme-based
	// admin-route discriminator. The cases below cover only the
	// previously-broken substring case that the fix resolves.
	cases := []struct {
		name   string
		method string
		path   string
		want   string
	}{
		{"prefix exact", "GET", "/kiri/health", "kiri-health"},
		{"bucket name shares prefix substring", "PUT", "/kiri-audit-bad-bucket", "bucket-put"},
		{"bucket name with longer admin prefix substring", "PUT", "/lambda-deploy-bucket", "bucket-put"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			called = ""

			req := httptest.NewRequest(tc.method, tc.path, http.NoBody)
			rec := httptest.NewRecorder()
			r.ServeHTTP(rec, req)

			if called != tc.want {
				t.Fatalf("path=%s: got handler %q, want %q (status=%d)",
					tc.path, called, tc.want, rec.Code)
			}
		})
	}
}

// TestExtractRoutePrefix_BoundaryGuard makes sure pattern registration
// itself doesn't classify a wildcard like `/{bucket}` as a `/kiri`
// prefixed route.
func TestExtractRoutePrefix_BoundaryGuard(t *testing.T) {
	t.Parallel()

	cases := []struct {
		pattern string
		want    string
	}{
		{"/kiri/health", "/kiri"},
		{"/lambda/2015-03-31/functions", "/lambda"},
		{"/{bucket}", ""},
		{"/{bucket}/{key...}", ""},
		{"/kirisomething", ""}, // no slash boundary → not a prefix
	}

	for _, tc := range cases {
		t.Run(tc.pattern, func(t *testing.T) {
			got := extractRoutePrefix(tc.pattern)
			if got != tc.want {
				t.Errorf("extractRoutePrefix(%q) = %q, want %q", tc.pattern, got, tc.want)
			}
		})
	}
}

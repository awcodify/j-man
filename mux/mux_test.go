package mux

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/awcodify/j-man/utils"
	"github.com/stretchr/testify/assert"
)

var expectedResponse = "Hello World!"

func TestServeHTTP(t *testing.T) {
	app := http.NewServeMux()
	app.HandleFunc("/test", testHandler)
	subdomains := Subdomains{
		"test": app,
	}
	r := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "http://test.localhost/test", nil)
	utils.DieIf(err)
	subdomains.ServeHTTP(r, req)

	// Check the status code is what we expect.
	assert.Equal(t, r.Code, http.StatusOK)

	// Check the response body is what we expect.
	assert.Equal(t, r.Body.String(), expectedResponse)

	r2 := httptest.NewRecorder()

	req, err = http.NewRequest("GET", "http://notfound.localhost/test", nil)
	utils.DieIf(err)
	subdomains.ServeHTTP(r2, req)

	// Since `notfound` subdomain not registered, it should return 404
	assert.Equal(t, r2.Code, http.StatusNotFound)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, expectedResponse)
}

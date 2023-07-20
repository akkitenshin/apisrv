package apiserver

import (
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"
)

func testAPIServer_HandleHello(t *testing.T) {
	s := New(NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	s.handleHello().ServerHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "Hello")
}

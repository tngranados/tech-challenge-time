package controllers

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/gin-gonic/gin"
	"github.com/tngranados/tech-challenge-time/storage"
)

func performRequest(tb testing.TB, r http.Handler, method, path string, body []byte) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method, path, bytes.NewBuffer(body))
	require.NoError(tb, err)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func createTestServer(tb testing.TB) (*gin.Engine, func()) {
	// Create test db.
	db, cleanup := storage.NewTestDB(tb)

	router, err := SetupRouter(true, db)
	require.NoError(tb, err)

	return router, cleanup
}

func checkCode(expectedCode, actualCode int) error {
	if actualCode != expectedCode {
		return fmt.Errorf("expected return code %v, but got %v instead", expectedCode, actualCode)
	}
	return nil
}

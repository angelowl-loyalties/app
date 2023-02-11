package tests

import (
	"github.com/cs301-itsa/project-2022-23t2-g1-t7/profiler/routes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// refer here: https://circleci.com/blog/gin-gonic-testing/

func TestUserHealth(t *testing.T) {
	router := gin.Default()
	routes.InitialiseRoutes(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)

	expectedResponse := `{"status":"OK"}`
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expectedResponse, w.Body.String())
}

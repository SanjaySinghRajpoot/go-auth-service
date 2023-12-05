package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetAllUsers(t *testing.T) {
	r := SetUpRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

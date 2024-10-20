package main

import (
	"conta-pagamento/controllers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupTestRoutes() *gin.Engine {
	return gin.Default()
}

func TestStatusCodeGetAccounts(t *testing.T) {
	routerMap := SetupTestRoutes()

	routerMap.GET("/account/:id", controllers.GetById)

	req, _ := http.NewRequest("GET", "/3", nil)
	res := httptest.NewRecorder()

	routerMap.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code, "asd")
}

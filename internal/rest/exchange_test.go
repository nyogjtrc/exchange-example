package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestExchange(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	assert.NoError(t, LoadRateFile("../../rate.json"))

	QueryExchange(r)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/exchange?source=USD&target=JPY&amount=$1,525", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, `{"msg":"success","amount":"$170,496.53"}`, w.Body.String())
}

func TestLoadRate(t *testing.T) {
	if assert.NoError(t, LoadRateFile("../../rate.json")) {
		assert.Equal(t, 1.0, rateMap["TWD"]["TWD"])
		assert.Equal(t, 1.0, rateMap["JPY"]["JPY"])
		assert.Equal(t, 1.0, rateMap["USD"]["USD"])
	}
}

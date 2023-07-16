package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestExchange(t *testing.T) {
	r := gin.Default()

	QueryExchange(r)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "?source=USD&target=JPY&amount=$1,525", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, `{"msg":"success","amount":"$170.496.53"}`, w.Body.String())
}

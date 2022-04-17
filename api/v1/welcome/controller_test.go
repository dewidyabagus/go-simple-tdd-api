package welcome_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	echo "github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"go-simple-api/api/v1/welcome"
)

var (
	welcomeJSON = `{"message":"Hello World!"}`
)

func TestGetIndex(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	handler := welcome.NewController()
	if assert.NoError(t, handler.GetIndex(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), welcomeJSON)
	}

}

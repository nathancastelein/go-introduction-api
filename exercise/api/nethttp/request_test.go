package nethttp

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadInputs(t *testing.T) {
	t.Run("QueryParamString", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest(http.MethodGet, "/?input=bar", nil)

		// Act
		param := QueryParamString(req, "input")

		// Assert
		require.Equal(t, "bar", param)
	})

	t.Run("QueryParamInt", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest(http.MethodGet, "/?input=42", nil)

		// Act
		param, err := QueryParamInt(req, "input")

		// Assert
		require.Nil(t, err)
		require.Equal(t, 42, param)
	})

	t.Run("Header", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Add("Content-Type", "application/json")

		// Act
		param := Header(req, "Content-Type")

		// Assert
		require.Equal(t, "application/json", param)
	})

	t.Run("PathParameter", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.SetPathValue("name", "Ada")

		// Act
		param := Path(req, "name")

		// Assert
		require.Equal(t, "Ada", param)
	})

	t.Run("PathParameterInt", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.SetPathValue("id", "42")

		// Act
		param, err := PathInt(req, "id")

		// Assert
		require.NoError(t, err)
		require.Equal(t, 42, param)
	})

	t.Run("Body", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBufferString(`Hello World`))

		// Act
		param, err := Body(req)

		// Assert
		require.NoError(t, err)
		require.Equal(t, "Hello World", string(param))
	})
}

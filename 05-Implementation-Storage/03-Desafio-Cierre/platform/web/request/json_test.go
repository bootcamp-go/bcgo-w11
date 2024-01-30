package request_test

import (
	"app/platform/web/request"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// Tests for JSON function
func TestRequestJSON(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// arrange
		type schema struct {
			Name string `json:"name"`
		}

		// act
		inputSchema := schema{}
		inputRequest := http.Request{
			Header: http.Header{"Content-Type": []string{"application/json"}},
			Body: io.NopCloser(strings.NewReader(`{"name":"test"}`)),
		}
		err := request.JSON(&inputRequest, &inputSchema)

		// assert
		expectedSchema := schema{Name: "test"}
		require.NoError(t, err)
		require.Equal(t, expectedSchema, inputSchema)
	})

	t.Run("error - content-type", func(t *testing.T) {
		// arrange
		type schema struct {
			Name string `json:"name"`
		}

		// act
		inputSchema := schema{}
		inputRequest := http.Request{
			Header: http.Header{"Content-Type": []string{"application/xml"}},
			Body: io.NopCloser(strings.NewReader(`{"name":"test"}`)),
		}
		err := request.JSON(&inputRequest, &inputSchema)

		// assert
		expectedSchema := schema{}
		require.ErrorIs(t, err, request.ErrRequestContentTypeNotJSON)
		require.EqualError(t, err, "request content type is not application/json")
		require.Equal(t, expectedSchema, inputSchema)
	})

	t.Run("error - json", func(t *testing.T) {
		// arrange
		type schema struct {
			Name string `json:"name"`
		}

		// act
		inputSchema := schema{}
		inputRequest := http.Request{
			Header: http.Header{"Content-Type": []string{"application/json"}},
			Body: io.NopCloser(strings.NewReader(`{"name":"test"`)),
		}
		err := request.JSON(&inputRequest, &inputSchema)

		// assert
		expectedSchema := schema{}
		require.ErrorIs(t, err, request.ErrRequestJSONInvalid)
		require.EqualError(t, err, "request json invalid. unexpected EOF")
		require.Equal(t, expectedSchema, inputSchema)
	})
}
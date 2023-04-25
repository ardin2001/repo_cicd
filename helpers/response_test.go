package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponse(t *testing.T) {
	responseModel := ResponseModel{
		Data:    "OK",
		Message: "coba_test",
		Status:  true,
	}

	assert.Equal(t, "OK", responseModel.Data)
	assert.Equal(t, "coba_test", responseModel.Message)
	assert.Equal(t, true, responseModel.Status)
}

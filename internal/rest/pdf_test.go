package rest_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	faker "github.com/go-faker/faker/v4"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/patcharanant/go-pdf-api/domain"
	"github.com/patcharanant/go-pdf-api/internal/rest"
)

func TestProcess(t *testing.T) {
	var mockRequest domain.ProcessRequest
	err := faker.FakeData(&mockRequest)
	assert.NoError(t, err)
	j, err := json.Marshal(mockRequest)

	assert.NoError(t, err)
	mockUCase := new(mocks.PDFService)
	mockUCase.On("Process", mock.Anything).Return(mockRequest, nil)

	e := echo.New()
	req, err := http.NewRequestWithContext(context.TODO(),
		echo.POST, "/v1/process", strings.NewReader(string(j)))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := rest.PDFHandler{
		Service: mockUCase,
	}
	err = handler.Process(c)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestProcessError(t *testing.T) {
	var mockRequest domain.ProcessRequest
	err := faker.FakeData(&mockRequest)
	assert.NoError(t, err)
	j, err := json.Marshal(mockRequest)
	mockErr := errors.New("mock")
	assert.NoError(t, err)
	mockUCase := new(mocks.PDFService)
	mockUCase.On("Process", mock.Anything).Return(nil, mockErr)

	e := echo.New()
	req, err := http.NewRequestWithContext(context.TODO(),
		echo.POST, "/v1/process", strings.NewReader(string(j)))

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := rest.PDFHandler{
		Service: mockUCase,
	}
	err = handler.Process(c)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockUCase.AssertExpectations(t)
}

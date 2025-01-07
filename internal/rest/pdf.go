package rest

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/patcharanant/go-pdf-api/domain"
)

type PDFService interface {
	Start(ctx context.Context, req domain.StartRequest) (*domain.StartResponse, error)
	Upload(ctx context.Context, req domain.UploadRequest) (*domain.UploadResponse, error)
	Process(ctx context.Context, req domain.ProcessRequest) (*domain.ProcessResponse, error)
	Download(ctx context.Context, req domain.DownloadRequest) (*domain.DownloadResponse, error)
}

type PDFHandler struct {
	PDFService PDFService
}

// NewArticleHandler will initialize the articles/ resources endpoint
func NewPDFHandler(e *echo.Echo, svc PDFService) {
	handler := &PDFHandler{
		PDFService: svc,
	}
	e.GET("/start/:tool", handler.Upload)
	e.POST("/upload", handler.Upload)
	e.POST("/process", handler.Process)
	e.GET("/download/:task", handler.Download)
}

func (h *PDFHandler) Start(c echo.Context) error {
	var req domain.StartRequest
	ctx := c.Request().Context()

	if err := c.Bind(&req); err != nil {
		return c.JSON(domain.GetStatusCode(err), domain.ResponseError{Message: domain.ErrBadParamInput.Error()})
	}

	result, err := h.PDFService.Start(ctx, req)
	if err != nil {
		return c.JSON(domain.GetStatusCode(err), domain.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func (h *PDFHandler) Upload(c echo.Context) error {

	var req domain.UploadRequest
	ctx := c.Request().Context()

	if err := c.Bind(&req); err != nil {
		return c.JSON(domain.GetStatusCode(err), domain.ResponseError{Message: domain.ErrBadParamInput.Error()})
	}

	result, err := h.PDFService.Upload(ctx, req)
	if err != nil {
		return c.JSON(domain.GetStatusCode(err), domain.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func (h *PDFHandler) Process(c echo.Context) error {
	var req domain.ProcessRequest
	ctx := c.Request().Context()

	if err := c.Bind(&req); err != nil {
		return c.JSON(domain.GetStatusCode(err), domain.ResponseError{Message: domain.ErrBadParamInput.Error()})
	}

	result, err := h.PDFService.Process(ctx, req)
	if err != nil {
		return c.JSON(domain.GetStatusCode(err), domain.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func (h *PDFHandler) Download(c echo.Context) error {
	var req domain.DownloadRequest
	ctx := c.Request().Context()

	if err := c.Bind(&req); err != nil {
		return c.JSON(domain.GetStatusCode(err), domain.ResponseError{Message: domain.ErrBadParamInput.Error()})
	}

	result, err := h.PDFService.Download(ctx, req)
	if err != nil {
		return c.JSON(domain.GetStatusCode(err), domain.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

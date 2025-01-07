package main

import (
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"

	"github.com/joho/godotenv"
	"github.com/patcharanant/go-pdf-api/internal/repository/pdfcpu"
	"github.com/patcharanant/go-pdf-api/internal/rest"
	"github.com/patcharanant/go-pdf-api/internal/rest/middleware"
	"github.com/patcharanant/go-pdf-api/pdf"
)

const (
	defaultTimeout = 30
	defaultAddress = ":9090"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// prepare echo

	e := echo.New()
	e.Use(middleware.CORS)
	timeoutStr := os.Getenv("CONTEXT_TIMEOUT")
	timeout, err := strconv.Atoi(timeoutStr)
	if err != nil {
		log.Println("failed to parse timeout, using default timeout")
		timeout = defaultTimeout
	}
	timeoutContext := time.Duration(timeout) * time.Second
	e.Use(middleware.SetRequestContextWithTimeout(timeoutContext))
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(200, "PDF API IS RUNNING")
	})
	// Prepare Repository
	pdfCpuRepo := pdfcpu.NewPDFCPURepository()

	// Build service Layer
	pdfSvc := pdf.NewService(pdfCpuRepo)
	rest.NewPDFHandler(e, pdfSvc)

	// Start Server
	address := os.Getenv("SERVER_ADDRESS")
	if address == "" {
		address = defaultAddress
	}
	log.Fatal(e.Start(address)) //nolint
}

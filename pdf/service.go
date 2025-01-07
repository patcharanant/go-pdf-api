package pdf

import (
	"context"
	"errors"

	"github.com/patcharanant/go-pdf-api/domain"
	"github.com/patcharanant/go-pdf-api/utils"
)

type PDFCPURepository interface {
	Optimize(inFile string, outFile string) error
	Split(inFile string, outDir string, pageNrs []int) error
}

type Service struct {
	pdfcpuRepo PDFCPURepository
}

var (
	ErrToolNotFound error = errors.New("Tools Not Found")
	ErrNotImplement error = errors.New("Tool Not Implement Yet")
)

func NewService(pdfcpuRepo PDFCPURepository) *Service {
	return &Service{
		pdfcpuRepo: pdfcpuRepo,
	}
}

func (s *Service) Start(ctx context.Context, req domain.StartRequest) (*domain.StartResponse, error) {
	return &domain.StartResponse{}, nil
}

func (s *Service) Upload(ctx context.Context, req domain.UploadRequest) (*domain.UploadResponse, error) {
	return &domain.UploadResponse{}, nil
}
func (s *Service) Process(ctx context.Context, req domain.ProcessRequest) (*domain.ProcessResponse, error) {
	//check task id ==> skip
	switch req.Tool {
	case "compress":
		inFile := utils.ToInFilePath(req.Files.ServerFileName)
		outFile := utils.ToOutFilePath(req.Files.Filename)
		err := s.pdfcpuRepo.Optimize(inFile, outFile)
		if err != nil {
			return nil, err
		}
		return &domain.ProcessResponse{
			DownloadFileName: req.Files.Filename,
			FileSize:         utils.GetFileSize(inFile),
			OutputFileSize:   utils.GetFileSize(outFile),
			OutputExtensions: "[pdf]",
			Timer:            "0.00",
			Status:           "FileSuccess",
		}, nil
	case "edit":
		return nil, ErrNotImplement
	case "split":
		// inFile := utils.ToInFilePath(req.Files.ServerFileName)
		// outFile := utils.ToOutFilePath(req.Files.Filename) // outFIr?
		// err := s.pdfcpuRepo.Split(inFile, outFile)
		// if err != nil {
		// 	return nil, err
		// }
		// //
		// return &domain.ProcessResponse{
		// 	DownloadFileName: req.Files.Filename,
		// 	FileSize:         utils.GetFileSize(inFile),
		// 	OutputFileSize:   utils.GetFileSize(outFile),
		// 	OutputExtensions: "[pdf]",
		// 	Timer:            "0.00",
		// 	Status:           "FileSuccess",
		// }, nil
	default:
		return nil, ErrToolNotFound
	}
	return nil, ErrToolNotFound
}
func (s *Service) Download(ctx context.Context, req domain.DownloadRequest) (*domain.DownloadResponse, error) {
	return &domain.DownloadResponse{}, nil
}

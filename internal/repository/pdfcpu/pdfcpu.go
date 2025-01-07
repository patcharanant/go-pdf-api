package pdfcpu

import (
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

type PDFCPURepository struct {
	conf *model.Configuration
}

func NewPDFCPURepository() *PDFCPURepository {
	conf := model.NewDefaultConfiguration()
	return &PDFCPURepository{
		conf: conf,
	}
}

func (r *PDFCPURepository) Optimize(inFile string, outFile string) error {
	err := api.OptimizeFile(inFile, outFile, r.conf)
	if err != nil {
		return err
	}
	return nil
}

func (r *PDFCPURepository) Split(inFile string, outDir string, pageNrs []int) error {
	err := api.SplitByPageNrFile(inFile, outDir, pageNrs, r.conf)
	if err != nil {
		return err
	}
	return nil
}

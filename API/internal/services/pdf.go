package services

// PDFService: stub for automatic PDF generation and archival (Mission 1).
type PDFService struct{ OutputDir string }

func NewPDFService(outputDir string) *PDFService { return &PDFService{OutputDir: outputDir} }

func (s *PDFService) Generate(title string, content []byte) (string, error) {
	// TODO: generate real PDF and keep a copy accessible.
	return s.OutputDir + "/" + title + ".pdf", nil
}

package utils

//PDFRenderer is an interface for PDF content renderer
type PDFRenderer interface {
	RenderPDF() (*string, error)
}

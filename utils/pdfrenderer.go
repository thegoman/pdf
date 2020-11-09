package utils

//PDFRenderer is an interface for PDF content rendering
type PDFRenderer interface {
	RenderPDF() (*string, error)
}

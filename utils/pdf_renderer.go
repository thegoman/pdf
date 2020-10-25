package utils

type PDFRenderer interface {
	RenderPDF() (*string, error)
}

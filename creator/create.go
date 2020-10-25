package creator

import (
	"github.com/thegoman/pdf/utils"
)

//PDF is the main structure which stores all pdf data and interfaces. It creates the PDF File
type PDF struct {
	objects []utils.PDFRenderer
}
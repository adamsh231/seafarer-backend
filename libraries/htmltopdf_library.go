package libraries

import (
	"html/template"
	"os"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/adamsh231/seafarer-packages/helpers"
)

type HtmlToPdfLibrary struct {
}

func NewHtmlToPdfLibrary() HtmlToPdfLibrary {
	return HtmlToPdfLibrary{}
}

func (lib HtmlToPdfLibrary) GeneratePdf(nameFile, pathFile, OutputFile string, fmap template.FuncMap, dataTemplate interface{}) (pdfFile *os.File, err error) {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, err
	}

	// parsing email template
	templateHelper := helpers.NewTemplateHelper()

	tpl, err := templateHelper.ParseTemplateToBufferFuncMap(nameFile, pathFile, fmap, dataTemplate)

	if err != nil {
		return nil, err
	}

	pdfg.OutputFile = OutputFile

	pdfg.AddPage(wkhtmltopdf.NewPageReader(tpl))
	err = pdfg.Create()
	if err != nil {
		return nil, err
	}

	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdfg.Dpi.Set(300)
	pdfFile, err = os.Open(pdfg.OutputFile)
	if err != nil {
		return nil, err
	}
	return pdfFile, err
}

package utils

import (
	"bytes"
	"fmt"
	"html/template"
)

type TemplateUtil struct {
}

func NewTemplateUtil() TemplateUtil {
	return TemplateUtil{}
}

func (util TemplateUtil) ParseTemplateToBuffer(filePathToParse string, data interface{}) (tpl *bytes.Buffer, err error) {

	// new buffer
	tpl = new(bytes.Buffer)

	// parse file to template
	tmpl, err := template.ParseFiles(filePathToParse)
	if err != nil {
		fmt.Println(err.Error())
		return tpl, err
	}

	// execute template into buffer
	err = tmpl.Execute(tpl, data)
	if err != nil {
		fmt.Println(err.Error())
		return tpl, err
	}

	return tpl, err
}

func (util TemplateUtil) ParseTemplateToBufferFuncMap(fileNamePathToParse, filePathToParse string, fMap template.FuncMap, data interface{}) (tpl *bytes.Buffer, err error) {

	tpl = new(bytes.Buffer)

	// parse file to template
	tmpl, err := template.New(fileNamePathToParse).Funcs(fMap).ParseFiles(filePathToParse)
	if err != nil {
		fmt.Println(err.Error())
		return tpl, err
	}

	// execute template into buffer
	err = tmpl.Funcs(fMap).Execute(tpl, data)
	if err != nil {
		fmt.Println(err.Error())
		return tpl, err
	}

	return tpl, err
}

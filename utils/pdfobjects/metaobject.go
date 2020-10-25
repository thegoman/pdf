package pdfobjects

import (
	"fmt"
	"strings"
	"time"
)

//MetaObject represents a PDF meta object which stores document meta data
type MetaObject struct {
	title        string
	author       string
	subject      string
	keywords     []string
	creator      string
	producer     string
	creationDate time.Time
	modDate      time.Time
}

//NewMetaObject creates a new meta object
func NewMetaObject(title string) (*MetaObject, error) {
	mo := &MetaObject{
		title:    "",
		author:   "",
		subject:  "",
		keywords: make([]string, 0),
		creator:  "GoMan PDF",
		producer: "GoMan PDF",
	}

	err := mo.SetTitle(title)
	if err != nil {
		return nil, err
	}

	return mo, nil
}

//RenderPDF renders the meta object to a PDF object
func (mo *MetaObject) RenderPDF() (*string, error) {
	output := fmt.Sprintf("\\Title (%s)\n", mo.title)

	if mo.author != "" {
		output += fmt.Sprintf("\\Author (%s)\n", mo.author)
	}

	output += fmt.Sprintf("\\Subject (%s)\n", mo.subject)

	if len(mo.keywords) > 0 {
		output += fmt.Sprintf("\\Keywords (%s)\n", strings.Join(mo.keywords, " "))
	}

	output += fmt.Sprintf("\\Creator (%s)\n", mo.creator)
	output += fmt.Sprintf("\\Producer (%s)\n", mo.producer)

	tm := time.Now()
	mo.creationDate = tm
	mo.modDate = tm

	output += fmt.Sprintf("\\CreationDate (D:%s)\n", mo.creationDate.Format("20060102150405-07'00'"))
	output += fmt.Sprintf("\\ModDate (D:%s)\n", mo.modDate.Format("20060102150405-07'00'"))

	return &output, nil
}

//SetTitle sets the title of the document
func (mo *MetaObject) SetTitle(title string) error {
	if title == "" {
		return fmt.Errorf("invalid title: %q", title)
	}
	mo.title = title
	return nil
}

//GetTitle gets the title of the document
func (mo *MetaObject) GetTitle() string {
	return mo.title
}

//SetAuthor sets the author of the document
func (mo *MetaObject) SetAuthor(author string) {
	mo.author = author
}

//GetAuthor gets the author of the document
func (mo *MetaObject) GetAuthor() string {
	return mo.author
}

//GetSubject gets the subject of the document
func (mo *MetaObject) GetSubject() string {
	return mo.subject
}

//SetSubject sets the subject of the document
func (mo *MetaObject) SetSubject(subject string) {
	mo.subject = subject
}

//GetKeywords gets all keywords of the document as []string array
func (mo *MetaObject) GetKeywords() []string {
	return mo.keywords
}

//AddKeyword add a new keyword to the document
func (mo *MetaObject) AddKeyword(keyword string) {
	mo.keywords = append(mo.keywords, keyword)
}

//GetCreator get the document creator
func (mo *MetaObject) GetCreator() string {
	return mo.creator
}

//SetCreator set the document creator i.e. the program which creates the pdf
func (mo *MetaObject) SetCreator(creator string) error {
	if creator == "" {
		return fmt.Errorf("invalid creator: %q", creator)
	}
	mo.creator = creator
	return nil
}

//GetProducer gets the library which creates the pdf: this can't be changed
func (mo *MetaObject) GetProducer() string {
	return mo.producer
}

package pdf_objects

import (
	"fmt"
	"strings"
	"time"
)

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

func (mo *MetaObject) SetTitle(title string) error {
	if title == "" {
		return fmt.Errorf("invalid title: %q", title)
	}
	mo.title = title
	return nil
}

func (mo *MetaObject) GetTitle() string {
	return mo.title
}

func (mo *MetaObject) SetAuthor(author string) {
	mo.author = author
}

func (mo *MetaObject) GetAuthor() string {
	return mo.author
}

func (mo *MetaObject) GetSubject() string {
	return mo.subject
}

func (mo *MetaObject) SetSubject(subject string) {
	mo.subject = subject
}

func (mo *MetaObject) GetKeywords() []string {
	return mo.keywords
}

func (mo *MetaObject) AddKeyword(keyword string) {
	mo.keywords = append(mo.keywords, keyword)
}

func (mo *MetaObject) GetCreator() string {
	return mo.creator
}

func (mo *MetaObject) SetCreator(creator string) error {
	if creator == "" {
		return fmt.Errorf("invalid creator: %q", creator)
	}
	mo.creator = creator
	return nil
}

func (mo *MetaObject) GetProducer() string {
	return mo.producer
}

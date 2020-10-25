package pdfobjects

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

var (
	texts = []struct {
		input string
	}{
		{"A Text"},
		{"Another Text"},
		{"And Another One"},
	}
)

func TestMetaObject_GetTitle(t *testing.T) {
	mo, _ := NewMetaObject("Hello World")
	assert.EqualValues(t, "Hello World", mo.GetTitle())
}

func TestMetaObject_GetAuthor(t *testing.T) {
	mo, _ := NewMetaObject("Test")
	assert.EqualValues(t, "", mo.GetAuthor())
}

func TestMetaObject_GetSubject(t *testing.T) {
	mo, _ := NewMetaObject("Test")
	assert.EqualValues(t, "", mo.GetSubject())
}

func TestMetaObject_GetKeywords(t *testing.T) {
	mo, _ := NewMetaObject("Test")
	kw := mo.GetKeywords()
	assert.EqualValues(t, 0, len(kw))
}

func TestMetaObject_GetCreator(t *testing.T) {
	mo, _ := NewMetaObject("Test")
	assert.EqualValues(t, "GoMan PDF", mo.GetCreator())
}

func TestMetaObject_GetProducer(t *testing.T) {
	mo, _ := NewMetaObject("Test")
	assert.EqualValues(t, "GoMan PDF", mo.GetProducer())
}

func TestMetaObject_ErrorSetEmptyTitle(t *testing.T) {
	_, err := NewMetaObject("")
	assert.NotNil(t, err)
	assert.Errorf(t, err, "invalid title: \"\"")
}

func TestMetaObject_SetTitle(t *testing.T) {
	mo, _ := NewMetaObject("Test")
	for _, title := range texts {
		_ = mo.SetTitle(title.input)
		assert.EqualValues(t, title.input, mo.GetTitle())
	}
}

func TestMetaObject_SetSubject(t *testing.T) {
	mo, _ := NewMetaObject("Test")
	for _, author := range texts {
		mo.SetAuthor(author.input)
		assert.EqualValues(t, author.input, mo.GetAuthor())
	}
}

func TestMetaObject_SetAuthor(t *testing.T) {
	mo, _ := NewMetaObject("Test")
	for _, author := range texts {
		mo.SetAuthor(author.input)
		assert.EqualValues(t, author.input, mo.GetAuthor())
	}
}

func TestMetaObject_AddKeyword(t *testing.T) {
	mo, _ := NewMetaObject("Test")
	for _, kw := range texts {
		mo.AddKeyword(kw.input)
	}

	kwords := mo.GetKeywords()
	assert.EqualValues(t, len(texts), len(kwords))

	for i, kw := range texts {
		assert.EqualValues(t, kw.input, kwords[i])
	}

}

func TestMetaObject_ErrorSetEmptyCreator(t *testing.T) {
	mo, _ := NewMetaObject("")
	err := mo.SetCreator("")
	assert.NotNil(t, err)
	assert.Errorf(t, err, "invalid title: \"\"")
}

func TestMetaObject_SetCreator(t *testing.T) {
	mo, _ := NewMetaObject("Test")
	for _, creator := range texts {
		_ = mo.SetCreator(creator.input)
		assert.EqualValues(t, creator.input, mo.GetCreator())
	}
}

func TestMetaObject_RenderPDF(t *testing.T) {
	mo, _ := NewMetaObject("Test")

	obj, _ := mo.RenderPDF()

	t.Run("render title", func(t *testing.T) {
		for _, title := range texts {
			_ = mo.SetTitle(title.input)
			obj, _ = mo.RenderPDF()
			assert.Contains(t, *obj, fmt.Sprintf("\\Title (%s)\n", title.input))
		}
	})

	t.Run("render author", func(t *testing.T) {
		assert.NotContains(t, "\\Author", *obj)
		for _, author := range texts {
			mo.SetAuthor(author.input)
			obj, _ = mo.RenderPDF()
			assert.Contains(t, *obj, fmt.Sprintf("\\Author (%s)\n", author.input))
		}
	})

	t.Run("render subject", func(t *testing.T) {
		assert.NotContains(t, "\\Subject", *obj)
		for _, subject := range texts {
			mo.SetSubject(subject.input)
			obj, _ = mo.RenderPDF()
			assert.Contains(t, *obj, fmt.Sprintf("\\Subject (%s)\n", subject.input))
		}
	})

	t.Run("render keywords", func(t *testing.T) {
		assert.NotContains(t, "\\Keywords", *obj)
		for _, kw := range texts {
			mo.AddKeyword(kw.input)
			obj, _ = mo.RenderPDF()
		}

		assert.Contains(t, *obj, fmt.Sprintf(
			"\\Keywords (%s)\n",
			strings.Join(mo.GetKeywords(), " "),
		))
	})

	t.Run("render producer", func(t *testing.T) {
		obj, _ = mo.RenderPDF()
		assert.Contains(t, *obj, "\\Producer (GoMan PDF)\n")
	})

	t.Run("render creation date", func(t *testing.T) {
		obj, _ = mo.RenderPDF()
		tm := time.Now()

		assert.Contains(t, *obj, fmt.Sprintf("\\CreationDate (D:%s", tm.Format("20060102")))
	})

	t.Run("render modification date", func(t *testing.T) {
		obj, _ = mo.RenderPDF()
		tm := time.Now()

		assert.Contains(t, *obj, fmt.Sprintf("\\ModDate (D:%s", tm.Format("20060102")))
	})

	fmt.Println(*obj)
}

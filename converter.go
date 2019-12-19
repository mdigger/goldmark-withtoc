// Package converter adds the ability to get a table of contents to the
// goldmark parser.
package withtoc

import (
	"io"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/util"
)

// Converter is markdown converting function.
type Converter = func(source []byte, writer io.Writer) ([]Header, error)

func markdown(m goldmark.Markdown) Converter {
	m.Parser().AddOptions(
		parser.WithAttribute(),
		parser.WithAutoHeadingID(),
		parser.WithASTTransformers(
			util.Prioritized(defaultTransformer, 0),
		),
	)
	return func(source []byte, writer io.Writer) ([]Header, error) {
		var ctx = parser.NewContext(parser.WithIDs(newIDs()))
		if err := m.Convert(source, writer, parser.WithContext(ctx)); err != nil {
			return nil, err
		}
		if toc, ok := ctx.Get(tocResultKey).([]Header); ok {
			return toc, nil
		}
		return nil, nil
	}
}

// New return markdown converter with table of content support.
func New(options ...goldmark.Option) Converter {
	return markdown(goldmark.New(options...))
}

// Convert from markdown to html and return TOC.
func Convert(m goldmark.Markdown, source []byte, writer io.Writer) ([]Header, error) {
	return markdown(m)(source, writer)
}

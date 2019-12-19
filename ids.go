package converter

import (
	"fmt"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/util"
)

type ids struct {
	counter int
	values  map[string]interface{}
}

func newIDs() parser.IDs {
	return &ids{
		values: make(map[string]interface{}),
	}
}

func (s *ids) Generate(value []byte, kind ast.NodeKind) []byte {
	for {
		s.counter++
		result := fmt.Sprintf("toc:%02d", s.counter)
		if _, ok := s.values[result]; !ok {
			s.values[result] = true
			return []byte(result)
		}
	}
}

func (s *ids) Put(value []byte) {
	s.values[util.BytesToReadOnlyString(value)] = true
}
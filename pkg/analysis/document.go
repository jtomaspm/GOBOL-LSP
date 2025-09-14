package analysis

import (
	"github.com/jtomaspm/GOBOL/pkg/lexer"
)

type Document struct {
	URI   string
	Lines []DocumentLine
}

type DocumentLine struct {
	Content string
	Tokens  []lexer.Token
}

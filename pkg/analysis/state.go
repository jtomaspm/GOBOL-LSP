package analysis

import (
	"fmt"
	"strings"

	"github.com/jtomaspm/GOBOL-LSP/pkg/lsp"
	"github.com/jtomaspm/GOBOL/pkg/lexer"
)

type State struct {
	Documents map[string]Document
}

func NewState() *State {
	return &State{Documents: map[string]Document{}}
}

func (s *State) OpenDocument(uri, text string) {
	document := Document{
		URI:   uri,
		Lines: []DocumentLine{},
	}
	lines := strings.Split(text, "\n")
	for line, content := range lines {
		l := lexer.New(strings.NewReader(content))
		tokens := []lexer.Token{}
		for _, token := range l.GetTokens() {
			token.Line = line
			tokens = append(tokens, token)
		}
		document.Lines = append(document.Lines, DocumentLine{
			Content: content,
			Tokens:  tokens,
		})
	}
	s.Documents[uri] = document
}

func (s *State) UpdateDocument(uri, text string) {
	s.OpenDocument(uri, text)
}

func (s *State) Hover(uri string, position lsp.Position, response *lsp.HoverResponse) {
	document, ok := s.Documents[uri]
	if !ok {
		return
	}
	if len(document.Lines) < position.Line {
		return
	}
	line := document.Lines[position.Line]
	b := []byte(line.Content)
	if len(b) < position.Character-1 {
		return
	}
	if b[position.Character-1] == byte(' ') {
		return
	}
	var result *lexer.Token
	for _, token := range line.Tokens {
		if !(token.Column <= position.Character) {
			break
		}
		result = &token
	}
	if result == nil {
		return
	}
	response.Result = lsp.HoverResult{
		Contents: fmt.Sprintf("Type: %s    Value: %s    Character: %c", result.Type, result.Literal, b[position.Character-1]),
	}
}

func (s *State) Definition(uri string, position lsp.Position, response *lsp.DefinitionResponse) {
	response.Result = lsp.Location{
		URI: uri,
		Range: lsp.Range{
			Start: lsp.Position{
				Line:      position.Line - 1,
				Character: 8,
			},
			End: lsp.Position{
				Line:      position.Line - 1,
				Character: 16,
			},
		},
	}
}

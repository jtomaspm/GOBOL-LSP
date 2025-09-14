package analysis

import (
	"fmt"
	"strings"

	"github.com/jtomaspm/GOBOL-LSP/pkg/lsp"
)

type State struct {
	Documents map[string][]string
}

func NewState() *State {
	return &State{Documents: map[string][]string{}}
}

func (s *State) OpenDocument(uri, text string) {
	s.Documents[uri] = strings.Split(text, "\n")
}

func (s *State) UpdateDocument(uri, text string) {
	s.Documents[uri] = strings.Split(text, "\n")
}

func (s *State) Hover(uri string, position lsp.Position, response *lsp.HoverResponse) {
	response.Result.Contents = fmt.Sprintf("File: %s, Character: %d/%d", uri, position.Character+1, len(s.Documents[uri][position.Line]))
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

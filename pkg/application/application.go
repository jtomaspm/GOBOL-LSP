package application

import (
	"bufio"
	"encoding/json"
	"os"

	"github.com/jtomaspm/GOBOL-LSP/pkg/analysis"
	"github.com/jtomaspm/GOBOL-LSP/pkg/lsp"
	"github.com/jtomaspm/GOBOL-LSP/pkg/rpc"
)

type Application struct {
	logger   *Logger
	state    *analysis.State
	scanner  *bufio.Scanner
	settings Settings
}

func NewApplication() Application {
	settings := NewSettings()
	logger := NewLogger(settings)
	scanner := NewScanner(settings)
	state := analysis.NewState()
	return Application{
		logger:   logger,
		settings: settings,
		scanner:  scanner,
		state:    state,
	}
}

func (app *Application) Run() {
	for app.scanner.Scan() {
		msg := app.scanner.Bytes()
		method, content, err := rpc.DecodeMessage(msg)
		if err != nil {
			app.logger.Printf("Error: %s", err)
			continue
		}
		app.handleMessage(method, content)
	}
}

func (app *Application) handleMessage(method string, content []byte) {
	app.logger.Printf("Received message: [%s] %s", method, content)
	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(content, &request); err != nil {
			app.logger.Printf("Error: %s", err)
		}
		app.logger.Printf("Connected: %s [%s]", request.Params.ClientInfo.Name, request.Params.ClientInfo.Version)
		response := lsp.NewInitializeResponse(request.ID)
		encodedResponse := []byte(rpc.EncodeMessage(response))
		os.Stdout.Write(encodedResponse)
	case "textDocument/didOpen":
		var request lsp.DidOpenTextDocumentNotification
		if err := json.Unmarshal(content, &request); err != nil {
			app.logger.Printf("Error: %s", err)
		}
		app.logger.Printf("Opened: %s", request.Params.TextDocument.URI)
		app.state.OpenDocument(request.Params.TextDocument.URI, request.Params.TextDocument.Text)
	case "textDocument/didChange":
		var request lsp.DidChangeTextDocumentNotification
		if err := json.Unmarshal(content, &request); err != nil {
			app.logger.Printf("Error: %s", err)
		}
		app.logger.Printf("Changed: %s", request.Params.TextDocument.URI)
		for _, change := range request.Params.ContentChanges {
			app.state.UpdateDocument(request.Params.TextDocument.URI, change.Text)
		}
	}
}

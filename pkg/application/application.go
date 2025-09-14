package application

import (
	"bufio"
	"encoding/json"
	"io"

	"github.com/jtomaspm/GOBOL-LSP/pkg/analysis"
	"github.com/jtomaspm/GOBOL-LSP/pkg/lsp"
	"github.com/jtomaspm/GOBOL-LSP/pkg/rpc"
)

type Application struct {
	logger             *Logger
	state              *analysis.State
	scanner            *bufio.Scanner
	settings           Settings
	writer             io.Writer
	notificationCenter *NotificationCenter
}

func NewApplication() Application {
	settings := NewSettings()
	logger := NewLogger(settings)
	scanner := NewScanner(settings)
	writer := NewWriter(settings)
	state := analysis.NewState()
	nc := NewNotificationCenter(writer, logger)
	return Application{
		logger:             logger,
		settings:           settings,
		scanner:            scanner,
		state:              state,
		writer:             writer,
		notificationCenter: nc,
	}
}

func (app *Application) Run() {
	app.notificationCenter.Run()
	app.logger.Println("GOBOL-LSP running...")
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
		app.writeResponse(response)
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
	case "textDocument/hover":
		var request lsp.HoverRequest
		if err := json.Unmarshal(content, &request); err != nil {
			app.logger.Printf("Error: %s", err)
		}
		app.logger.Printf("Hover: %s", request.Params.TextDocument.URI)
		response := lsp.HoverResponse{
			Response: lsp.Response{
				ID:  &request.ID,
				RPC: "2.0",
			},
			Result: lsp.HoverResult{
				Contents: "",
			},
		}
		app.state.Hover(request.Params.TextDocument.URI, request.Params.Position, &response)
		app.writeResponse(response)
	case "textDocument/definition":
		var request lsp.DefinitionRequest
		if err := json.Unmarshal(content, &request); err != nil {
			app.logger.Printf("Error: %s", err)
		}
		app.logger.Printf("Definition: %s", request.Params.TextDocument.URI)
		response := lsp.DefinitionResponse{
			Response: lsp.Response{
				ID:  &request.ID,
				RPC: "2.0",
			},
			Result: lsp.Location{},
		}
		app.state.Definition(request.Params.TextDocument.URI, request.Params.Position, &response)
		app.writeResponse(response)
	}
}

func (app *Application) writeResponse(msg any) {
	encodedResponse := []byte(rpc.EncodeMessage(msg))
	if _, err := app.writer.Write(encodedResponse); err != nil {
		app.logger.Printf("failed to write response: %v", err)
	}
}

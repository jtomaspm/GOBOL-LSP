package rpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

const VERSION = "2.0"

type RPCMessage struct {
	Method string `json:"method"`
}

func EncodeMessage(msg any) string {
	content, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content)
}

// Returns the method name and content length
func DecodeMessage(data []byte) (string, []byte, error) {
	header, content, found := bytes.Cut(data, []byte("\r\n\r\n"))
	if !found {
		return "", nil, fmt.Errorf("invalid message: missing header-body separator: %s", data)
	}

	contentLengthBytes := header[len("Content-Length: "):]
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		return "", nil, fmt.Errorf("invalid Content-Length: %v", err)
	}

	content = content[:contentLength]

	var baseMessage RPCMessage
	if err = json.Unmarshal(content, &baseMessage); err != nil || baseMessage.Method == "" {
		return "", nil, fmt.Errorf("invalid JSON content: %v", err)
	}

	return baseMessage.Method, content, nil
}

func Split(data []byte, _ bool) (advance int, token []byte, err error) {
	header, content, found := bytes.Cut(data, []byte("\r\n\r\n"))
	if !found {
		return 0, nil, nil // Need more data
	}

	contentLengthBytes := header[len("Content-Length: "):]
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		return 0, nil, fmt.Errorf("invalid Content-Length: %v", err)
	}

	if len(content) < contentLength {
		return 0, nil, nil // Need more data
	}

	totalLength := len(header) + 4 + contentLength
	return totalLength, data[:totalLength], nil
}

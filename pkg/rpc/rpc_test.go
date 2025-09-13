package rpc_test

import (
	"testing"

	"github.com/jtomaspm/GOBOL-LSP/pkg/rpc"
)

func TestEncodeMessage(t *testing.T) {
	//Arrange
	msg := struct {
		A string `json:"a"`
		B int    `json:"b"`
	}{
		A: "hello",
		B: 42,
	}

	//Act
	encoded := rpc.EncodeMessage(msg)

	//Assert
	expected := "Content-Length: 20\r\n\r\n{\"a\":\"hello\",\"b\":42}"
	if encoded != expected {
		t.Fatalf("Expected %q, got %q", expected, encoded)
	}
}

func TestDecodeMessage(t *testing.T) {
	//Arrange
	data := []byte("Content-Length: 25\r\n\r\n{\"method\":\"hello\",\"b\":42}")

	//Act
	method, content, err := rpc.DecodeMessage(data)

	//Assert
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	expectedMethod := "hello"
	expectedContent := []byte("{\"method\":\"hello\",\"b\":42}")
	if string(content) != string(expectedContent) {
		t.Fatalf("Expected content %q, got %q", expectedContent, content)
	}
	if method != expectedMethod {
		t.Fatalf("Expected method %q, got %q", expectedMethod, method)
	}
}

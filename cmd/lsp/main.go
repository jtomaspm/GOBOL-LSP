package main

import (
	"bufio"
	"log"
	"os"

	"github.com/jtomaspm/GOBOL-LSP/pkg/rpc"
)

func main() {
	logger := getLogger("/home/pop/Code/GOBOL-LSP/gobol-lsp.log")
	logger.Println("GOBOL LSP started")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, content, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("Error: %s", err)
			continue
		}
		handleMessage(logger, method, content)
	}
}

func handleMessage(logger *log.Logger, method string, content []byte) {
	logger.Printf("Received message: [%s] %s", method, content)
}

func getLogger(filename string) *log.Logger {
	f, err := os.OpenFile(filename, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	return log.New(f, "[GOBOL-LSP]", log.Ldate|log.Ltime|log.Lshortfile)
}

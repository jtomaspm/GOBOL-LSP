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
		msg := scanner.Text()
		handleMessage(logger, msg)
	}
}

func handleMessage(logger *log.Logger, msg any) {
	logger.Println("Received message:\n", msg)
}

func getLogger(filename string) *log.Logger {
	f, err := os.OpenFile(filename, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	return log.New(f, "[GOBOL-LSP]", log.Ldate|log.Ltime|log.Lshortfile)
}

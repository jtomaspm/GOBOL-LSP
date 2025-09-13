package application

import (
	"bufio"
	"os"

	"github.com/jtomaspm/GOBOL-LSP/pkg/rpc"
)

func NewScanner(settings Settings) *bufio.Scanner {
	switch settings.Interface {
	case "stdio":
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Split(rpc.Split)
		return scanner
	default:
		panic(settings.Interface + " is not supported...")
	}
}

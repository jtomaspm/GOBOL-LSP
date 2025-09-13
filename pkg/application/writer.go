package application

import (
	"io"
	"os"
)

func NewWriter(settings Settings) io.Writer {
	switch settings.Interface {
	case "stdio":
		return os.Stdout
	default:
		panic(settings.Interface + " is not supported...")
	}
}

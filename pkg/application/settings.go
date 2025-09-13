package application

import (
	"flag"
	"fmt"
	"os"
)

type Settings struct {
	LogPath   string `json:"logPath"`
	Interface string `json:"interface"`
}

func NewSettings() Settings {
	// Define flags
	helpFlag := flag.Bool("help", false, "Show help")
	hFlag := flag.Bool("h", false, "Show help (short)")
	logPathFlag := flag.String("log_path", "", "Path to log file")
	ifaceFlag := flag.String("interface", "stdio", "Interface type")

	// Parse command-line flags
	flag.Parse()

	// Show help if requested
	if *helpFlag || *hFlag {
		fmt.Println("Usage: GOBOL-LSP -interface <stdio|tcp> -log_path <path>")
		os.Exit(0)
	}

	// Return settings
	return Settings{
		LogPath:   *logPathFlag,
		Interface: *ifaceFlag,
	}
}

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
	if helpFlag := flag.Bool("help", false, "Show help"); helpFlag != nil && *helpFlag {
		fmt.Println("GOBOL-LSP -interface <stdio|tcp> -log_path <path>")
		os.Exit(0)
	}
	if helpFlag := flag.Bool("h", false, "Show help"); helpFlag != nil && *helpFlag {
		fmt.Println("GOBOL-LSP -interface <stdio|tcp> -log_path <path>")
		os.Exit(0)
	}
	logPath := ""
	if logPathFlag := flag.String("log_path", "", "Path to log file"); logPathFlag != nil {
		logPath = *logPathFlag
	}
	iface := ""
	if ifaceFlag := flag.String("interface", "stdio", "Interface type"); ifaceFlag != nil {
		iface = *ifaceFlag
	}
	return Settings{
		LogPath:   logPath,
		Interface: iface,
	}
}

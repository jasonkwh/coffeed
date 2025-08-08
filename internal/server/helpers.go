package server

import (
	"os"
	"path/filepath"
)

func getSocketPath() string {
	if runtimeDir := os.Getenv("XDG_RUNTIME_DIR"); runtimeDir != "" {
		return filepath.Join(runtimeDir, "coffeed.sock")
	}

	// Fallback for systems without XDG_RUNTIME_DIR
	homeDir, _ := os.UserHomeDir()
	return filepath.Join(homeDir, ".local", "run", "coffeed.sock")
}

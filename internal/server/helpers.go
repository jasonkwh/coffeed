package server

import (
	"os"
	"path/filepath"
)

func getSocketPath() (string, error) {
	if runtimeDir := os.Getenv("XDG_RUNTIME_DIR"); runtimeDir != "" {
		return filepath.Join(runtimeDir, "coffeed.sock"), nil
	}

	// Fallback for systems without XDG_RUNTIME_DIR
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, ".local", "run", "coffeed.sock"), nil
}

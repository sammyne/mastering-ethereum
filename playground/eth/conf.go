package eth

import (
	"os"
	"path/filepath"
)

var defaultDataDir string

func init() {
	home, ok := os.LookupEnv("HOME")
	if !ok {
		panic("missing $HOME env")
	}

	defaultDataDir = filepath.Join(home, ".mastering-ethereum")
}

// DefaultDataDir returns the default data directory
func DefaultDataDir() string {
	return defaultDataDir
}

// DefaultKeyDir returns the default keystore directory
func DefaultKeyDir() string {
	return filepath.Join(defaultDataDir, "keystore")
}

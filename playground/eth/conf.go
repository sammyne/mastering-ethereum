package eth

import (
	"path/filepath"

	"os"
)

var defaultDataDir string

func init() {
	home, err := os.UserHomeDir()
	if nil != err {
		panic(err)
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

// DefaultPassphrase defines the default passphrase to use
func DefaultPassphrase() string {
	return "hello"
}

// ProviderURL is the url of the kind Infura
func ProviderURL() string {
	return "https://ropsten.infura.io/v3/f3df74d615a74774821985274dedcc9e"
}

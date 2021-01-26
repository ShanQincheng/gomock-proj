package internal

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// GoimportsMockDir invoke command "goimports -w MockDir"
func GoimportsMockDir() error {
	goimports := filepath.Join(os.Getenv("GOPATH"), `bin`, `goimports`)

	flag0 := `-w`
	cmd := exec.Command(goimports, flag0, MockDir)
	output, err := cmd.CombinedOutput() // goimports -w test/mocks
	if err != nil {
		return fmt.Errorf("cmd.CombinedOutput: %s, command: %s, output: %s", err, cmd.String(), string(output))
	}

	return nil
}

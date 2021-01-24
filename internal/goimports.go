package internal

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func GoimportsMockDir() error {
	goimports := filepath.Join(os.Getenv("GOPATH"), `bin`, `goimports`)

	flag0 := `-w`
	cmd := exec.Command(goimports, flag0, MockDir)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("cmd.CombinedOutput: %s, command: %s, output: %s", err, cmd.String(), string(output))
	}

	return nil
}

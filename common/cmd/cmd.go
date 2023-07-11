package cmd

import (
	"context"
	"os/exec"
)

// Cmd 执行命令
func Cmd(ctx context.Context, resultOutput chan<- string, resultErr chan<- error, instruction string) {
	cmd := exec.CommandContext(ctx, "/bin/sh", "-c", instruction)
	output, err := cmd.CombinedOutput()
	resultOutput <- string(output)
	resultErr <- err
}

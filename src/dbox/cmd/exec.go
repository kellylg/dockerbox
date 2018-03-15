package cmd

import (
	"github.com/goinbox/golog"
	"github.com/goinbox/shell"

	"strings"
)

type ExecCommand struct {
}

func (e *ExecCommand) Run(args []string, logger golog.ILogger) {
	dconfItem, err := getDconfItemFromArgs(args)
	if err != nil {
		logger.Error([]byte("get dconfItem error: " + err.Error()))
		return
	}

	cmd := "docker exec -it " + dconfItem.ContainerName + " "
	cmd += dconfItem.Exec.ShellCmd + " '"
	cmd += dconfItem.Exec.PreCmd + ";"
	cmd += strings.Join(args, " ") + "'"

	logger.Debug([]byte("cmd: " + cmd))

	shell.RunCmdBindTerminal(cmd)
}

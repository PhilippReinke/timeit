package app

import (
	"context"
	"fmt"
)

const (
	NAME = "timeit"
)

type App struct {
	cmds map[string]Command
}

func New(cmds ...Command) *App {
	cmdMap := make(map[string]Command, len(cmds))
	for _, c := range cmds {
		cmdMap[c.Name()] = c
	}

	return &App{cmds: cmdMap}
}

func (a *App) Usage() {
	fmt.Printf("CLI based stopwatch and timer.\n\n")

	fmt.Printf("Usage: %v [OPTIONS] <COMMAND>\n\n", NAME)

	fmt.Println("Commands:")
	for _, c := range a.cmds {
		fmt.Printf("  %s: %s\n", c.Name(), c.Description())
	}
}

func (a *App) Run(ctx context.Context, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("no command provided")
	}

	name := args[0]
	cmd, ok := a.cmds[name]
	if !ok {
		return fmt.Errorf("unknown command: %s", name)
	}

	return cmd.Run(ctx, args[1:])
}

package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"

	"github.com/PhilippReinke/timeit/app"
	"github.com/PhilippReinke/timeit/cmds/stopwatch"
	timerCmd "github.com/PhilippReinke/timeit/cmds/timer"
)

func main() {
	var help bool
	flag.BoolVar(&help, "help", false, "Display help this app.")
	flag.Parse()

	stopwatchCmd := stopwatch.New()
	timerCmd := timerCmd.New()
	app := app.New(stopwatchCmd, timerCmd)

	if help {
		app.Usage()
		return
	}

	// graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		<-signalChan
		cancel()
	}()

	if err := app.Run(ctx, os.Args[1:]); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

package stopwatch

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/PhilippReinke/timeit/app"
)

type Impl struct {
	name string
	fs   *flag.FlagSet
	help bool
}

var _ app.Command = &Impl{}

func New() *Impl {
	cmd := &Impl{
		name: "stopwatch",
		fs:   flag.NewFlagSet("stopwatch", flag.ExitOnError),
	}

	cmd.fs.Usage = func() {
		fmt.Printf("%v\n\n", cmd.Description())
		fmt.Printf("Usage: %v %v\n", app.NAME, cmd.Name())
		// fmt.Println("Options:")
		// cmd.fs.PrintDefaults()
	}

	return cmd
}

func (i *Impl) Name() string {
	return i.name
}

func (_ *Impl) Description() string {
	return "Start a stopwatch."
}

func (i *Impl) Run(ctx context.Context, args []string) error {
	if err := i.fs.Parse(args); err != nil {
		return err
	}

	defer func() {
		fmt.Print("\n")
	}()

	start := time.Now()
	lastPrint := ""
	for {
		select {
		case <-time.Tick(time.Millisecond * 10):
			elapsed := time.Since(start)
			h := elapsed / time.Hour
			m := (elapsed % time.Hour) / time.Minute
			s := (elapsed % time.Minute) / time.Second

			out := fmt.Sprintf("%02d:%02d:%02d", h, m, s)
			if out == lastPrint {
				continue
			}
			fmt.Print("\r", out)
			lastPrint = out
		case <-ctx.Done():
			return nil
		}
	}
}

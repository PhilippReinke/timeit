package timer

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"time"

	"github.com/gen2brain/beeep"

	"github.com/PhilippReinke/timeit/app"
)

type Impl struct {
	name   string
	fs     *flag.FlagSet
	hour   int
	min    int
	sec    int
	notify bool
	help   bool
}

var _ app.Command = &Impl{}

func New() *Impl {
	cmd := &Impl{
		name: "timer",
		fs:   flag.NewFlagSet("timer", flag.ExitOnError),
	}

	cmd.fs.IntVar(
		&cmd.hour,
		"hour",
		0,
		"Timer duration in hours.",
	)
	cmd.fs.IntVar(
		&cmd.min,
		"min",
		0,
		"Timer duration in minutes.",
	)
	cmd.fs.IntVar(
		&cmd.sec,
		"sec",
		0,
		"Timer duration in seconds.",
	)
	cmd.fs.BoolVar(
		&cmd.notify,
		"notify",
		true,
		"OS notification when timer done.",
	)

	cmd.fs.Usage = func() {
		fmt.Printf("%v\n\n", cmd.Description())
		fmt.Printf("Usage: %v %v [OPTIONS]\n\n", app.NAME, cmd.Name())
		fmt.Println("Options:")
		cmd.fs.PrintDefaults()
	}

	return cmd
}

func (i *Impl) Name() string {
	return i.name
}

func (_ *Impl) Description() string {
	return "Start a timer for given duration."
}

func (i *Impl) Run(ctx context.Context, args []string) error {
	if err := i.fs.Parse(args); err != nil {
		return err
	}

	// vailidate input
	if i.hour < 0 || i.min < 0 || i.sec < 0 {
		return errors.New("time must be non-negative")
	}
	if i.hour+i.min+i.sec == 0 {
		return errors.New("zero timer is not allowed")
	}

	// timer logic
	defer func() {
		fmt.Print("\n")
	}()

	duration := time.Hour*time.Duration(i.hour) +
		time.Minute*time.Duration(i.min) +
		time.Second*time.Duration(i.sec)
	finish := time.Now().Add(duration)
	lastPrint := ""
	for {
		select {
		case <-time.Tick(time.Millisecond * 10):
			difference := time.Until(finish)

			if difference < 0 {
				if i.notify {
					beeep.Notify("Timer finished", duration.String()+" have elapsed.", "")
				}

				return nil
			}

			h := difference / time.Hour
			m := (difference % time.Hour) / time.Minute
			s := (difference % time.Minute) / time.Second

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

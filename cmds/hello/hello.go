package hello

import (
	"context"
	"flag"
	"fmt"

	"github.com/PhilippReinke/timeit/app"
)

type Impl struct {
	name  string
	fs    *flag.FlagSet
	help  bool
	greet string
}

var _ app.Command = &Impl{}

func New() *Impl {
	cmd := &Impl{
		name: "hello",
		fs:   flag.NewFlagSet("hello", flag.ExitOnError),
	}

	cmd.fs.StringVar(
		&cmd.greet,
		"greet",
		"world",
		"Name to greet.",
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
	return "Print customisable greeting formula."
}

func (i *Impl) Run(_ context.Context, args []string) error {
	if err := i.fs.Parse(args); err != nil {
		return err
	}

	fmt.Printf("Hello %s :)\n", i.greet)

	return nil
}

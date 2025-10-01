package app

import (
	"context"
)

type Command interface {
	Name() string
	Description() string
	Run(ctx context.Context, args []string) error
}

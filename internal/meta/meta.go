package meta

import (
	"context"
)

type Module interface {
	Scan(ctx context.Context, dir string) error
}

type Meta struct {
	Main     *Main
	CMD      []*CMD
	Apps     []*App
	Libs     []*Lib
	Storages []*Storage

	HasProjectStruct bool
}

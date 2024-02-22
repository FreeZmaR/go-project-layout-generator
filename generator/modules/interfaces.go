package modules

import (
	"context"
)

type Module interface {
	Name() string
	Description() string
	SelectedVersion() string
	Versions() []string

	SelectVersion(version string) error

	// Generate generates the module in the given directory.
	Generate(ctx context.Context, dir string, subModules []Module) error
}

type HTTPRouter interface {
	Name() string
	Version() string
	Package() string
	InstanceTemplate() string
}

package generator

import "context"

type Generator struct {
	dirPath   string
	mainState *mainState
	structure *Structure
}

type mainState struct {
	projectName string
	goVersion   string
}

func New() *Generator {
	return &Generator{mainState: &mainState{}}
}

func (g *Generator) Run(ctx context.Context) error {
	return nil
}

func (g *Generator) SetProjectName(name string) *Generator {
	g.mainState.projectName = name

	return g
}

func (g *Generator) SetGoVersion(version string) *Generator {
	g.mainState.goVersion = version

	return g
}

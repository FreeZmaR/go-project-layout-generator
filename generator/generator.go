package generator

import (
	"context"
)

type Generator struct {
	outputDir      string
	projectSetting *ProjectSetting
	structure      *Structure
}

func New() *Generator {
	return &Generator{projectSetting: newDefaultProjectSetting()}
}

func (g *Generator) Run(ctx context.Context) error {
	return nil
}

func (g *Generator) ParseStructure(ctx context.Context, filePath string) error {
	return nil
}

func (g *Generator) ParseDefaultStructure(ctx context.Context) error {
	return nil
}

func (g *Generator) ProjectSetting() *ProjectSetting {
	return g.projectSetting
}

func (g *Generator) SetOutputDir(dir string) *Generator {
	g.outputDir = dir

	return g
}

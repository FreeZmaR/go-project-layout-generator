package generator

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"os"
)

type Generator struct {
	outputDir         string
	projectSetting    *ProjectSetting
	structure         *Structure
	isStructureParsed bool
}

func New() *Generator {
	return &Generator{
		projectSetting:    newDefaultProjectSetting(),
		structure:         &Structure{},
		isStructureParsed: false,
	}
}

func (g *Generator) Run(ctx context.Context) error {
	if err := g.genOutputDir(ctx); err != nil {
		return err
	}

	return g.genProject(ctx)
}

func (g *Generator) ParseStructure(ctx context.Context, filePath string) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return g.parseStructure(file)
}

func (g *Generator) ParseDefaultStructure(ctx context.Context) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}

	g.projectSetting.SetDefault()

	structFile, err := defaultStructureFS.Open(defaultStructureFileName)
	if err != nil {
		return err
	}

	return g.parseStructure(structFile)
}

func (g *Generator) ProjectSetting() *ProjectSetting {
	return g.projectSetting
}

func (g *Generator) SetOutputDir(dir string) *Generator {
	g.outputDir = dir

	return g
}
func (g *Generator) parseStructure(reader io.Reader) error {
	if err := json.NewDecoder(reader).Decode(g.structure); err != nil {
		return err
	}

	if g.projectSetting.modName != "" {
		g.structure.ModName = g.projectSetting.modName
	}

	if g.projectSetting.projectName != "" {
		g.structure.ProjectName = g.projectSetting.projectName
	}

	if g.projectSetting.goVersion != "" {
		g.structure.GoVersion = g.projectSetting.goVersion
	}

	g.isStructureParsed = true

	return nil
}
func (g *Generator) genOutputDir(ctx context.Context) error {
	return g.genDir(ctx, "")
}

func (g *Generator) genProject(ctx context.Context) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}

	if nil == g.structure || !g.isStructureParsed {
		return errors.New("structure is not defined")
	}

	for _, dir := range g.structure.Dirs {
		if err := g.genStructureDir(ctx, "", dir); err != nil {
			return err
		}
	}

	for _, file := range g.structure.Files {
		if err := g.genFile(ctx, "", file); err != nil {
			return err
		}
	}

	return nil
}

func (g *Generator) genStructureDir(ctx context.Context, parentPath string, dir Dir) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}

	if g.projectSetting.withCodeExample && !dir.IsExample {
		return nil
	}

	if !g.projectSetting.withCodeExample && !dir.IsDefault {
		return nil
	}

	path := dir.Name
	if parentPath != "" {
		path = parentPath + "/" + dir.Name
	}

	if err := g.genDir(ctx, path); err != nil {
		return err
	}

	for _, file := range dir.Files {
		if err := g.genFile(ctx, path, file); err != nil {
			return err
		}
	}

	for _, nestedDir := range dir.Dirs {
		if err := g.genStructureDir(ctx, path, nestedDir); err != nil {
			return err
		}
	}

	return nil
}

func (g *Generator) genDir(ctx context.Context, dirName string) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}

	path := g.outputDir
	if dirName != "" {
		path += "/" + dirName
	}

	path = g.replacePlaceholders(path)

	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return os.MkdirAll(path, 0755)
		}

		return err
	}

	if !info.IsDir() {
		return errors.New(path + " is not a directory")
	}

	return nil
}

func (g *Generator) genFile(ctx context.Context, path string, file File) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}

	if g.projectSetting.withCodeExample && !file.IsExample {
		return nil
	}

	if !g.projectSetting.withCodeExample && !file.IsDefault {
		return nil
	}

	path = g.outputDir + "/" + path

	if path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}

	path = g.replacePlaceholders(path + "/" + file.Name)

	_, err := os.Stat(path)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	}

	f, errF := os.Create(path)
	if errF != nil {
		return errF
	}
	defer f.Close()

	if file.Content != "" {
		_, errW := io.WriteString(f, g.replacePlaceholders(file.Content))

		return errW
	}

	return nil
}

func (g *Generator) replacePlaceholders(content string) string {
	return replaceAll(
		content,
		g.structure.ModName,
		g.structure.ProjectName,
		g.structure.GoVersion,
	)
}

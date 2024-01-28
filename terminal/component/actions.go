package component

import (
	"context"
	"fmt"
	"github.com/FreeZmaR/go-project-layout-generator/generator"
)

func GenerateDefaultProjectAction(cancel Component, gen *generator.Generator, eventStack *EventStack) {
	inputOutputDir := NewInput("Output directory (default: . )", "Enter output directory")
	inputModeName := NewInput(
		fmt.Sprintf("Project Mod name (default: %s)", generator.DefaultModName),
		"Enter project name for mod.go",
	)
	inputProjectName := NewInput(
		fmt.Sprintf("Project name (default: %s)", generator.DefaultProjectName),
		"Enter project name",
	)
	inputGoVersion := NewInput(
		fmt.Sprintf("Go version (default: %s)", generator.DefaultGoVersion),
		"Enter go version",
	)

	inputPool := newInputInfoPool(inputGoVersion, inputProjectName, inputModeName)

	eventStack.Push(inputOutputDir)
	eventStack.Push(
		NewSelective(
			inputModeName,
			cancel,
			func() bool {
				if inputOutputDir.IsCancelled() {
					return false
				}

				if inputOutputDir.Value() != "" {
					gen.SetOutputDir(inputOutputDir.Value())
				}

				dirName := gen.OutputDir()
				if dirName == "" {
					dirName = "."
				}

				inputPool.PutInfoValue("output dir", dirName)

				return true
			},
		),
	)
	eventStack.Push(
		NewSelective(
			inputProjectName,
			cancel,
			func() bool {
				if inputModeName.IsCancelled() {
					return false
				}

				if inputModeName.Value() != "" {
					gen.ProjectSetting().SetModName(inputModeName.Value())
				}

				inputPool.PutInfoValue("mode name", gen.ProjectSetting().ModName())

				return true
			},
		),
	)
	eventStack.Push(
		NewSelective(
			inputGoVersion,
			cancel,
			func() bool {
				if inputProjectName.IsCancelled() {
					return false
				}

				if inputProjectName.Value() != "" {
					gen.ProjectSetting().SetGoVersion(inputProjectName.Value())
				}

				inputPool.PutInfoValue("project name", gen.ProjectSetting().ProjectName())

				return true
			},
		),
	)

	eventStack.Push(
		NewSelective(
			NewLoader(
				"Generating project",
				func(ctx context.Context) error {
					gen.ProjectSetting().SetWithCodeExample(false)

					if err := gen.ParseDefaultStructure(ctx); err != nil {
						return err
					}

					return gen.Run(ctx)
				},
				func(data any) error {
					fmt.Println("Project generated")

					return nil
				},
				func(data any) error {
					err, ok := data.(error)
					if ok {
						fmt.Println("Error generating project: ", err.Error())

						return err
					}

					fmt.Println("Error while generating project")

					return nil
				},
			),
			cancel,
			func() bool {
				if inputGoVersion.IsCancelled() {
					return false
				}

				if inputGoVersion.Value() != "" {
					gen.ProjectSetting().SetGoVersion(inputGoVersion.Value())
				}

				inputPool.PutInfoValue("go version", gen.ProjectSetting().GoVersion())

				return true
			},
		),
	)
}

//func GenerateDefaultProjectWithExampleAction(gen *generator.Generator, eventStack *EventStack) error {
//}

package component

import (
	"context"
	"fmt"
	"github.com/FreeZmaR/go-project-layout-generator/generator"
)

func GenerateDefaultProjectAction(cancel Component, gen *generator.Generator, eventStack *EventStack) {

	inputOutputDir := NewInput("Output directory", "Enter output directory")
	inputModeName := NewInput("Project Mod name", "Enter project name for mod.go")
	inputProjectName := NewInput("Project name", "Enter project name")

	eventStack.Push(inputOutputDir)
	eventStack.Push(
		NewSelective(
			inputModeName,
			cancel,
			func() bool {
				if inputOutputDir.IsCancelled() {
					return false
				}

				gen.SetOutputDir(inputOutputDir.Value())

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

				gen.ProjectSetting().SetModName(inputModeName.Value())

				return true
			},
		),
	)
	eventStack.Push(
		NewSelective(
			inputModeName,
			cancel,
			func() bool {
				if inputProjectName.IsCancelled() {
					return false
				}

				gen.ProjectSetting().SetProjectName(inputProjectName.Value())

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
				if inputModeName.IsCancelled() {
					return false
				}

				return true
			},
		),
	)
}

//func GenerateDefaultProjectWithExampleAction(gen *generator.Generator, eventStack *EventStack) error {
//}

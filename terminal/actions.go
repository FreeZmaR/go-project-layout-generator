package terminal

import (
	"context"
	"fmt"
	"github.com/FreeZmaR/go-project-layout-generator/generator"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func CreateDefaultProjectMenuAction(gen *generator.Generator) MenuItemAction {
	return func(style lipgloss.Style) (tea.Model, tea.Cmd) {
		gen.
			SetOutputDir("./template").
			ProjectSetting().
			SetWithCodeExample(false)

		return RunLoaderAction(
			style,
			"Generating project...",
			func(ctx context.Context, cancel context.CancelFunc) {
				defer cancel()

				if err := gen.ParseDefaultStructure(ctx); err != nil {
					fmt.Println("Error: ", err.Error())

					return
				}

				if err := gen.Run(ctx); err != nil {
					fmt.Println("Error: ", err.Error())
				}
			},
		)
	}
}

func CreateDefaultProjectWithExampleMenuAction(gen *generator.Generator) MenuItemAction {
	return func(style lipgloss.Style) (tea.Model, tea.Cmd) {
		gen.ProjectSetting().SetWithCodeExample(true)

		return RunLoaderAction(
			style,
			"Generating project with example...",
			func(ctx context.Context, cancel context.CancelFunc) {
				_ = gen.Run(ctx)
				cancel()
			},
		)
	}
}

func RunLoaderAction(
	style lipgloss.Style,
	text string,
	fn func(ctx context.Context, cancel context.CancelFunc),
) (tea.Model, tea.Cmd) {
	ctx, cancel := context.WithCancel(context.Background())

	go fn(ctx, cancel)

	l := NewLoader(ctx, cancel, style, text)

	return l, l.Init()
}

package terminal

import (
	"github.com/FreeZmaR/go-project-layout-generator/generator"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func CreateDefaultProjectMenuAction(gen *generator.Generator) MenuItemAction {
	return func(style lipgloss.Style) (tea.Model, tea.Cmd) {
		//TODO: logic
		return nil, nil
	}
}

func CreateDefaultProjectWithExampleMenuAction(gen *generator.Generator) MenuItemAction {
	return func(style lipgloss.Style) (tea.Model, tea.Cmd) {
		//TODO: logic
		return nil, nil
	}
}

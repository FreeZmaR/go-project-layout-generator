package terminal

import (
	"github.com/FreeZmaR/go-project-layout-generator/generator"
	"github.com/FreeZmaR/go-project-layout-generator/terminal/component"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

type Terminal struct {
	controller Controller
	generator  *generator.Generator
	eventStack *component.EventStack
}

func New() *Terminal {
	t := &Terminal{
		generator:  generator.New(),
		eventStack: component.NewEventStack(),
	}

	t.controller = NewController(t.eventStack, buildMenu(t.generator, t.eventStack))

	return t
}

func (t *Terminal) Run() error {
	_, err := tea.NewProgram(t.controller, tea.WithOutput(os.Stderr)).Run()

	return err
}

package terminal

import (
	"github.com/FreeZmaR/go-project-layout-generator/generator"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

type Terminal struct {
	controller Controller
	generator  *generator.Generator
}

func New() *Terminal {
	t := &Terminal{generator: generator.New()}
	t.controller = NewController(buildMenu(t.generator))

	return t
}

func (t *Terminal) Run() error {
	_, err := tea.NewProgram(t.controller, tea.WithOutput(os.Stderr)).Run()

	return err
}

package terminal

import (
	"github.com/FreeZmaR/go-project-layout-generator/generator"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
)

type menuBuilder struct {
	menu      *list.Model
	generator *generator.Generator
}

func buildMenu(gen *generator.Generator) list.Model {
	builder := menuBuilder{
		generator: gen,
	}

	builder.menu = builder.makeMenuList()
	builder.makeMainMenu()

	return *builder.menu
}

func (b menuBuilder) makeMenuList() *list.Model {
	l := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)

	l.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			key.NewBinding(
				key.WithKeys("<-"),
				key.WithHelp("<-", "back"),
			),
		}
	}

	l.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{
			key.NewBinding(
				key.WithKeys("enter", "<-"),
				key.WithHelp("enter", "select"),
				key.WithHelp("<-", "back"),
			),
		}
	}

	l.KeyMap.Filter.Unbind()
	l.KeyMap.GoToEnd.Unbind()
	l.KeyMap.GoToStart.Unbind()

	return &l
}

func (b menuBuilder) makeMainMenu() {
	b.menu.Title = "Generator menu"

	b.makeCreateProjectMenu(1)
	b.makeCreateModuleMenu(2)
}

func (b menuBuilder) makeCreateProjectMenu(position int) {
	menu := b.makeMenuList()

	b.makeDefaultProjectMenu(menu, b.menu, 1)

	b.menu.InsertItem(
		position,
		NewMenuItem(
			"Create project",
			"Create a new project",
			menu,
			nil,
			nil,
		),
	)
}

func (b menuBuilder) makeCreateModuleMenu(position int) {
	b.menu.InsertItem(
		position,
		NewMenuItem(
			"Create module",
			"Create a new module in the project",
			nil,
			nil,
			nil,
		),
	)
}

func (b menuBuilder) makeDefaultProjectMenu(l, parent *list.Model, position int) {
	l.InsertItem(position,
		NewMenuItem(
			"Default project",
			"Create a new project with default structure",
			nil,
			parent,
			nil,
		),
	)
}

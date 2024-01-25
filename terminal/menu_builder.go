package terminal

import (
	"github.com/FreeZmaR/go-project-layout-generator/generator"
	"github.com/FreeZmaR/go-project-layout-generator/terminal/component"
)

type menuBuilder struct {
	menu       *component.List
	generator  *generator.Generator
	eventStack *component.EventStack
}

func buildMenu(gen *generator.Generator, eventStack *component.EventStack) component.Component {
	builder := menuBuilder{
		generator:  gen,
		eventStack: eventStack,
	}

	builder.makeMainMenu()

	return builder.menu
}

func (b *menuBuilder) makeMainMenu() {
	b.menu = component.NewList(
		"Generator menu",
		b.makeCreateProjectMenu(),
		b.makeCreateModuleMenu(),
	)
}

func (b *menuBuilder) makeCreateProjectMenu() component.ListItem {
	menu := component.NewList(
		"Create project",
		b.makeDefaultProjectMenu(),
		b.makeDefaultProjectWithExampleMenu(),
	)

	return component.NewListItem(
		"Create project",
		"Create a new project",
		func(_ any) error {
			b.menu.SetIsNotDone()
			b.eventStack.Push(menu)

			return nil
		},
		func(_ any) error {
			b.menu.SetIsNotDone()
			b.eventStack.Push(b.menu)

			return nil
		},
	)
}

func (b *menuBuilder) makeCreateModuleMenu() component.ListItem {
	return component.NewListItem(
		"Create module",
		"Create a new module",
		func(_ any) error {
			return nil
		},
		func(_ any) error {
			return nil
		},
	)
}

func (b *menuBuilder) makeDefaultProjectMenu() component.ListItem {
	return component.NewListItem(
		"Create project",
		"Create a new project with default settings",
		func(data any) error {
			var list *component.List

			list = data.(*component.List)

			component.GenerateDefaultProjectAction(list, b.generator, b.eventStack)

			return nil
		},
		func(_ any) error {
			b.menu.SetIsNotDone()
			b.eventStack.Push(b.menu)

			return nil
		},
	)
}

func (b *menuBuilder) makeDefaultProjectWithExampleMenu() component.ListItem {
	return component.NewListItem(
		"Create default project with example",
		"Crate a new default project with code example(http-server)",
		func(_ any) error {
			return nil
		},
		func(_ any) error {
			b.menu.SetIsNotDone()
			b.eventStack.Push(b.menu)

			return nil
		},
	)
}

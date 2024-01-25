package terminal

import (
	"github.com/FreeZmaR/go-project-layout-generator/terminal/component"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Controller struct {
	style      lipgloss.Style
	eventStack *component.EventStack
	component  component.Component
	height     int
	width      int
}

var _ tea.Model = (*Controller)(nil)

func NewController(eventStack *component.EventStack, c component.Component) Controller {
	return Controller{
		style:      lipgloss.NewStyle().Margin(1, 2),
		eventStack: eventStack,
		component:  c,
	}
}

func (c Controller) Init() tea.Cmd {
	return c.component.Init()
}

func (c Controller) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch t := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := c.style.GetFrameSize()
		c.height = t.Height - v
		c.width = t.Width - h

		c.component.SetSize(c.width, c.height)
	}

	var cmd tea.Cmd
	c.component, cmd = c.component.Update(msg)

	c.component = c.handleEvent()

	return c, cmd
}

func (c Controller) View() string {
	return c.style.Render(c.component.View())
}

func (c Controller) handleEvent() component.Component {
	if !c.eventStack.Has() {
		return c.component
	}

	if !c.component.IsDone() {
		return c.component
	}

	comp := c.eventStack.Get()
	if comp == nil {
		return c.component
	}

	comp.SetSize(c.width, c.height)

	return comp
}

package terminal

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Controller struct {
	style           lipgloss.Style
	menu            list.Model
	showMenu        bool
	actionModel     tea.Model
	showActionModel bool
	height          int
	width           int
}

var _ tea.Model = (*Controller)(nil)

func NewController(l list.Model) Controller {
	return Controller{
		style:           lipgloss.NewStyle().Margin(1, 2),
		menu:            l,
		showMenu:        true,
		showActionModel: false,
	}
}

func (c Controller) Init() tea.Cmd {
	return nil
}

func (c Controller) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch t := msg.(type) {
	case tea.KeyMsg:
		return c.handleKeyAction(t)
	case tea.WindowSizeMsg:
		h, v := c.style.GetFrameSize()
		c.height = t.Height - v
		c.width = t.Width - h

		c.menu.SetSize(c.width, c.height)
	}

	if c.showMenu {
		return c.updateMenu(msg, c)
	}

	if c.showActionModel {
		return c.updateActionModel(msg, c)
	}

	return c, nil
}

func (c Controller) View() string {
	if c.showMenu {
		return c.style.Render(c.menu.View())
	}

	if c.showActionModel {
		return c.style.Render(c.actionModel.View())
	}

	return "something went wrong"
}

func (c Controller) updateMenu(msg tea.Msg, controller Controller) (Controller, tea.Cmd) {
	var cmd tea.Cmd
	controller.menu, cmd = controller.menu.Update(msg)

	return controller, cmd
}

func (c Controller) updateActionModel(msg tea.Msg, controller Controller) (Controller, tea.Cmd) {
	var cmd tea.Cmd
	controller.actionModel, cmd = controller.actionModel.Update(msg)

	return controller, cmd
}

func (c Controller) handleKeyAction(key tea.KeyMsg) (Controller, tea.Cmd) {
	switch key.String() {
	case "ctrl+c":
		return c, tea.Quit
	case "enter":
		return c.handleEnterAction(key)
	case "backspace", tea.KeyCtrlQuestionMark.String():
		return c.handleBackAction(key)
	}

	return c.updateMenu(key, c)
}

func (c Controller) handleEnterAction(msg tea.Msg) (Controller, tea.Cmd) {
	item, ok := c.menu.SelectedItem().(MenuItem)
	if !ok {
		return c.updateMenu(msg, c)
	}

	if nil == item.menu {
		return c.updateMenu(msg, c)
	}

	var cmd tea.Cmd

	c.menu, cmd = (*item.menu).Update(msg)
	c.menu.SetSize(c.width, c.height)

	return c, cmd
}

func (c Controller) handleBackAction(msg tea.Msg) (Controller, tea.Cmd) {
	if c.showActionModel {
		c.showActionModel = false
		c.showMenu = true

		var cmd tea.Cmd
		c.menu, cmd = c.menu.Update(msg)

		return c, cmd
	}

	item, ok := (c.menu.Items()[0]).(MenuItem)
	if !ok || nil == item.back {
		return c.updateMenu(msg, c)
	}

	var cmd tea.Cmd
	c.menu, cmd = (*item.back).Update(msg)
	c.menu.SetSize(c.width, c.height)

	return c, cmd
}

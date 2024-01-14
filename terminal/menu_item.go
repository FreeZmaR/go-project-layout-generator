package terminal

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MenuItem struct {
	title       string
	description string
	menu        *list.Model
	back        *list.Model
	action      MenuItemAction
}

type MenuItemAction func(style lipgloss.Style) (tea.Model, tea.Cmd)

var _ list.DefaultItem = (*MenuItem)(nil)

func NewMenuItem(title, description string, menu, back *list.Model, action MenuItemAction) MenuItem {
	return MenuItem{
		title:       title,
		description: description,
		menu:        menu,
		back:        back,
		action:      action,
	}
}

func (m MenuItem) Title() string {
	return m.title
}

func (m MenuItem) Description() string {
	return m.description
}

func (m MenuItem) FilterValue() string {
	return m.Title()
}

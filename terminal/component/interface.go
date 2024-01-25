package component

import tea "github.com/charmbracelet/bubbletea"

type Component interface {
	Init() tea.Cmd
	Update(msg tea.Msg) (Component, tea.Cmd)
	SetSize(width, height int)
	View() string
	IsDone() bool
}

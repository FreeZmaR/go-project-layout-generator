package component

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Selective struct {
	success Component
	failure Component
	checkFN func() bool
}

var _ Component = (*Selective)(nil)

func NewSelective(success, failure Component, checkFN func() bool) Selective {
	return Selective{
		success: success,
		failure: failure,
		checkFN: checkFN,
	}
}

func (s Selective) Init() tea.Cmd {
	return nil
}

func (s Selective) Update(_ tea.Msg) (Component, tea.Cmd) {
	if s.checkFN() {
		return s.success, s.success.Init()
	}

	return s.failure, s.failure.Init()
}

func (s Selective) SetSize(_, _ int) {}

func (s Selective) View() string {
	if s.checkFN() {
		return s.success.View()
	}

	return s.failure.View()
}

func (s Selective) IsDone() bool {
	return true
}
